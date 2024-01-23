package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Sender = (*SimpleSender)(nil)

// SimpleSender is a sender service that does best effort to send transactions to the destination chain.
type SimpleSender struct {
	clients  map[uint64]*ethclient.Client          // rpc clients per chain id
	sessions map[uint64]bindings.OmniPortalSession // sessions per chain id
}

// NewSimpleSender creates a new sender service that does best effort to send transactions to the destination chain.
func NewSimpleSender(chains []netconf.Chain, rpcClientPerChain map[uint64]*ethclient.Client,
	privateKey ecdsa.PrivateKey,
) (SimpleSender, error) {
	sessions := make(map[uint64]bindings.OmniPortalSession)
	for _, chain := range chains {
		session, err := newSession(chain, rpcClientPerChain[chain.ID], privateKey)
		if err != nil {
			return SimpleSender{}, err
		}

		sessions[chain.ID] = session
	}

	return SimpleSender{
		sessions: sessions,
		clients:  rpcClientPerChain,
	}, nil
}

// newSession creates a new session for the given chain. used to interact with portal contract.
func newSession(chain netconf.Chain, rpcClient *ethclient.Client,
	privateKey ecdsa.PrivateKey) (bindings.OmniPortalSession, error) {
	contract, err := bindings.NewOmniPortal(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return bindings.OmniPortalSession{}, errors.Wrap(err, "new contract")
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(&privateKey, big.NewInt(int64(chain.ID)))
	if err != nil {
		return bindings.OmniPortalSession{}, errors.Wrap(err, "new transactor")
	}

	session := bindings.OmniPortalSession{
		Contract:     contract,
		TransactOpts: *transactor,
		CallOpts: bind.CallOpts{
			From: crypto.PubkeyToAddress(privateKey.PublicKey),
		},
	}

	return session, nil
}

// SendTransaction sends the given submission to the destination chain. Best effort sending.
func (s SimpleSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	xChainSubmission := TranslateSubmission(submission)

	session, ok := s.sessions[submission.DestChainID]
	if !ok {
		return errors.New("session not found", "dest_chain_id", submission.DestChainID)
	}

	tx, err := session.Xsubmit(xChainSubmission)
	if err != nil {
		// todo(Lazar): handle error
		return errors.Wrap(err, "submit tx")
	}

	log.Info(ctx, "Submitted_tx",
		"tx_hash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_price", tx.GasPrice(),
	)

	// todo(Lazar): handle success case, metrics and cache

	return nil
}