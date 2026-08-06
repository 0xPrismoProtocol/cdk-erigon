package jsonrpc

import (
	"context"
	"testing"
)

// eth_feeHistory and eth_maxPriorityFeePerGas exist to price EIP-1559
// transactions, which this fork rejects (fork-12 has no typed-tx encoding in
// batchL2Data). Both must refuse immediately — before touching the database —
// so that fee-probing wallets fall back to eth_gasPrice and legacy txs.
// Calling them on a zero-value APIImpl proves the guard runs first: any db
// access would nil-panic.

func TestFeeHistoryDisabledOnLegacyOnlyChain(t *testing.T) {
	api := &APIImpl{}
	res, err := api.FeeHistory(context.Background(), 4, 0, []float64{50})
	if err == nil {
		t.Fatalf("FeeHistory must return an error on a legacy-only chain, got result: %v", res)
	}
}

func TestMaxPriorityFeePerGasDisabledOnLegacyOnlyChain(t *testing.T) {
	api := &APIImpl{}
	res, err := api.MaxPriorityFeePerGas(context.Background())
	if err == nil {
		t.Fatalf("MaxPriorityFeePerGas must return an error on a legacy-only chain, got result: %v", res)
	}
}
