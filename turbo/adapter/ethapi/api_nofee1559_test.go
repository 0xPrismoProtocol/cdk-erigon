package ethapi

import (
	"math/big"
	"testing"

	"github.com/ledgerwatch/erigon/core/types"
)

// Fork-12 batchL2Data has no typed-transaction encoding, so this chain accepts
// only legacy (type-0) transactions. Wallets (ethers, viem, MetaMask) decide to
// build an EIP-1559 transaction solely from the presence of baseFeePerGas in
// RPC block/header responses — so the RPC layer must never advertise it, even
// though London is active in consensus (the chainspec cannot be changed without
// changing the genesis hash).
func TestRPCMarshalHeaderOmitsBaseFeePerGas(t *testing.T) {
	head := &types.Header{
		Number:     big.NewInt(158253),
		Difficulty: big.NewInt(0),
		BaseFee:    big.NewInt(7), // London active in consensus: header carries a base fee
	}

	fields := RPCMarshalHeader(head)

	if _, ok := fields["baseFeePerGas"]; ok {
		t.Fatalf("RPCMarshalHeader must not emit baseFeePerGas on a legacy-only chain, got %v", fields["baseFeePerGas"])
	}
}
