// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package primitives

import (
	"github.com/berachain/beacon-kit/mod/primitives/math"
	"github.com/berachain/beacon-kit/mod/primitives/version"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ ExecutionPayload = (*ExecutableDataDeneb)(nil)

// ExecutableDataDeneb is the execution payload for Deneb.
//

//go:generate go run github.com/ferranbt/fastssz/sszgen -path payload.go -objs ExecutableDataDeneb -include ./primitives.go,./execution.go,./math,./withdrawal.go,./bytes.go,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil,$GOPATH/pkg/mod/github.com/holiman/uint256@v1.2.4 -output payload.ssz.go
//go:generate go run github.com/fjl/gencodec -type ExecutableDataDeneb -field-override executableDataDenebMarshaling -out payload.json.go
//nolint:lll
type ExecutableDataDeneb struct {
	ParentHash    ExecutionHash    `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
	FeeRecipient  ExecutionAddress `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
	StateRoot     Bytes32          `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
	ReceiptsRoot  Bytes32          `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
	LogsBloom     []byte           `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
	Random        Bytes32          `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
	Number        math.U64         `json:"blockNumber"                  gencodec:"required"`
	GasLimit      math.U64         `json:"gasLimit"                     gencodec:"required"`
	GasUsed       math.U64         `json:"gasUsed"                      gencodec:"required"`
	Timestamp     math.U64         `json:"timestamp"                    gencodec:"required"`
	ExtraData     []byte           `json:"extraData"                    gencodec:"required" ssz-max:"32"`
	BaseFeePerGas math.Wei         `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
	BlockHash     ExecutionHash    `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
	Transactions  [][]byte         `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
	Withdrawals   []*Withdrawal    `json:"withdrawals"                                      ssz-max:"16"`
	BlobGasUsed   math.U64         `json:"blobGasUsed"`
	ExcessBlobGas math.U64         `json:"excessBlobGas"`
}

// JSON type overrides for ExecutableDataDeneb.
type executableDataDenebMarshaling struct {
	ExtraData    hexutil.Bytes
	LogsBloom    hexutil.Bytes
	Transactions []hexutil.Bytes
}

// Version returns the version of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) Version() uint32 {
	return version.Deneb
}

// IsNil checks if the ExecutableDataDeneb is nil.
func (d *ExecutableDataDeneb) IsNil() bool {
	return d == nil
}

// IsBlinded checks if the ExecutableDataDeneb is blinded.
func (d *ExecutableDataDeneb) IsBlinded() bool {
	return false
}

// GetParentHash returns the parent hash of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetParentHash() ExecutionHash {
	return d.ParentHash
}

// GetFeeRecipient returns the fee recipient address of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetFeeRecipient() ExecutionAddress {
	return d.FeeRecipient
}

// GetStateRoot returns the state root of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetStateRoot() Bytes32 {
	return d.StateRoot
}

// GetReceiptsRoot returns the receipts root of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetReceiptsRoot() Bytes32 {
	return d.ReceiptsRoot
}

// GetLogsBloom returns the logs bloom of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetLogsBloom() []byte {
	return d.LogsBloom
}

// GetPrevRandao returns the previous Randao value of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetPrevRandao() Bytes32 {
	return d.Random
}

// GetNumber returns the block number of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetNumber() math.U64 {
	return d.Number
}

// GetGasLimit returns the gas limit of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetGasLimit() math.U64 {
	return d.GasLimit
}

// GetGasUsed returns the gas used of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetGasUsed() math.U64 {
	return d.GasUsed
}

// GetTimestamp returns the timestamp of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetTimestamp() math.U64 {
	return d.Timestamp
}

// GetExtraData returns the extra data of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetExtraData() []byte {
	return d.ExtraData
}

// GetBaseFeePerGas returns the base fee per gas of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetBaseFeePerGas() math.Wei {
	return d.BaseFeePerGas
}

// GetBlockHash returns the block hash of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetBlockHash() ExecutionHash {
	return d.BlockHash
}

// GetTransactions returns the transactions of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetTransactions() [][]byte {
	return d.Transactions
}

// GetWithdrawals returns the withdrawals of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetWithdrawals() []*Withdrawal {
	return d.Withdrawals
}

// GetBlobGasUsed returns the blob gas used of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetBlobGasUsed() math.U64 {
	return d.BlobGasUsed
}

// GetExcessBlobGas returns the excess blob gas of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetExcessBlobGas() math.U64 {
	return d.ExcessBlobGas
}