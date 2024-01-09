// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title Events
 * @dev Contains all events that are tested against. Required until solc allows
 *      for referencing events defined in other contracts.
 */
contract Events {
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );
}