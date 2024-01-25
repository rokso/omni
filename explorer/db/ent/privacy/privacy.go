// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The BlockQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BlockQueryRuleFunc func(context.Context, *ent.BlockQuery) error

// EvalQuery return f(ctx, q).
func (f BlockQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlockQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BlockQuery", q)
}

// The BlockMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BlockMutationRuleFunc func(context.Context, *ent.BlockMutation) error

// EvalMutation calls f(ctx, m).
func (f BlockMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BlockMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BlockMutation", m)
}

// The ChainQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ChainQueryRuleFunc func(context.Context, *ent.ChainQuery) error

// EvalQuery return f(ctx, q).
func (f ChainQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChainQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ChainQuery", q)
}

// The ChainMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ChainMutationRuleFunc func(context.Context, *ent.ChainMutation) error

// EvalMutation calls f(ctx, m).
func (f ChainMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ChainMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ChainMutation", m)
}

// The MsgQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MsgQueryRuleFunc func(context.Context, *ent.MsgQuery) error

// EvalQuery return f(ctx, q).
func (f MsgQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MsgQuery", q)
}

// The MsgMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MsgMutationRuleFunc func(context.Context, *ent.MsgMutation) error

// EvalMutation calls f(ctx, m).
func (f MsgMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MsgMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MsgMutation", m)
}

// The ReceiptQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ReceiptQueryRuleFunc func(context.Context, *ent.ReceiptQuery) error

// EvalQuery return f(ctx, q).
func (f ReceiptQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ReceiptQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ReceiptQuery", q)
}

// The ReceiptMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ReceiptMutationRuleFunc func(context.Context, *ent.ReceiptMutation) error

// EvalMutation calls f(ctx, m).
func (f ReceiptMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ReceiptMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ReceiptMutation", m)
}

// The XProviderCursorQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type XProviderCursorQueryRuleFunc func(context.Context, *ent.XProviderCursorQuery) error

// EvalQuery return f(ctx, q).
func (f XProviderCursorQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.XProviderCursorQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.XProviderCursorQuery", q)
}

// The XProviderCursorMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type XProviderCursorMutationRuleFunc func(context.Context, *ent.XProviderCursorMutation) error

// EvalMutation calls f(ctx, m).
func (f XProviderCursorMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.XProviderCursorMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.XProviderCursorMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.BlockQuery:
		return q.Filter(), nil
	case *ent.ChainQuery:
		return q.Filter(), nil
	case *ent.MsgQuery:
		return q.Filter(), nil
	case *ent.ReceiptQuery:
		return q.Filter(), nil
	case *ent.XProviderCursorQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.BlockMutation:
		return m.Filter(), nil
	case *ent.ChainMutation:
		return m.Filter(), nil
	case *ent.MsgMutation:
		return m.Filter(), nil
	case *ent.ReceiptMutation:
		return m.Filter(), nil
	case *ent.XProviderCursorMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}