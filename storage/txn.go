package storage

import (
	"context"

	"github.com/lemon-mint/tidb/v7/kv"
	"github.com/lemon-mint/tidb/v7/parser/model"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/tikv/client-go/v2/tikv"
)

// Txn is a txn that returns a retryAble error when called Commit.
type Txn struct {
	opts  map[int]interface{}
	valid bool
}

func (t *Txn) SetAssertion(_ []byte, _ ...kv.FlagsOp) error {
	return nil
}

// Commit always returns a retryable error.
func (t *Txn) Commit(ctx context.Context) error {
	return kv.ErrTxnRetryable
}

func (t *Txn) Rollback() error {
	t.valid = false
	return nil
}

func (t *Txn) String() string {
	return ""
}

func (t *Txn) LockKeys(_ context.Context, _ *kv.LockCtx, _ ...kv.Key) error {
	return nil
}

func (t *Txn) LockKeysFunc(_ context.Context, _ *kv.LockCtx, fn func(), _ ...kv.Key) error {
	if fn != nil {
		fn()
	}
	return nil
}
func (t *Txn) SetOption(opt int, val interface{}) {
	t.opts[opt] = val
}

func (t *Txn) GetOption(opt int) interface{} {
	return t.opts[opt]
}

func (t *Txn) IsReadOnly() bool {
	return true
}

func (t *Txn) StartTS() uint64 {
	return uint64(0)
}

func (t *Txn) Get(ctx context.Context, k kv.Key) ([]byte, error) {
	return nil, nil
}

func (t *Txn) BatchGet(ctx context.Context, keys []kv.Key) (map[string][]byte, error) {
	return nil, nil
}

func (t *Txn) Iter(k kv.Key, upperBound kv.Key) (kv.Iterator, error) {
	return nil, nil
}

func (t *Txn) IterReverse(k kv.Key) (kv.Iterator, error) {
	return nil, nil
}

func (t *Txn) Set(k kv.Key, v []byte) error {
	return nil
}

func (t *Txn) Delete(k kv.Key) error {
	return nil
}

func (t *Txn) Valid() bool {
	return t.valid
}

func (t *Txn) Len() int {
	return 0
}

func (t *Txn) Size() int {
	return 0
}

func (t *Txn) GetMemBuffer() kv.MemBuffer {
	return nil
}

func (t *Txn) GetSnapshot() kv.Snapshot {
	return nil
}

func (t *Txn) NewStagingBuffer() kv.MemBuffer {
	return nil
}

func (t *Txn) Flush() (int, error) {
	return 0, nil
}

func (t *Txn) Discard() {
}

func (t *Txn) Reset() {
	t.valid = false
}

func (t *Txn) SetVars(vars interface{}) {
}

func (t *Txn) GetVars() interface{} {
	return nil
}

func (t *Txn) CacheTableInfo(id int64, info *model.TableInfo) {
}

func (t *Txn) GetTableInfo(id int64) *model.TableInfo {
	return nil
}

func (t *Txn) SetDiskFullOpt(level kvrpcpb.DiskFullOpt) {
	// TODO nothing
}

func (t *Txn) GetMemDBCheckpoint() *tikv.MemDBCheckpoint {
	return nil
}

func (t *Txn) RollbackMemDBToCheckpoint(_ *tikv.MemDBCheckpoint) {
	// TODO nothing
}

func (t *Txn) ClearDiskFullOpt() {
	// TODO nothing
}

func (t *Txn) UpdateMemBufferFlags(_ []byte, _ ...kv.FlagsOp) {

}

func (t *Txn) SetMemoryFootprintChangeHook(func(uint64)) {

}

func (t *Txn) Mem() uint64 {
	return 0
}

func (t *Txn) StartAggressiveLocking() error                   { return nil }
func (t *Txn) RetryAggressiveLocking(_ context.Context) error  { return nil }
func (t *Txn) CancelAggressiveLocking(_ context.Context) error { return nil }
func (t *Txn) DoneAggressiveLocking(_ context.Context) error   { return nil }
func (t *Txn) IsInAggressiveLockingMode() bool                 { return false }

// newMockTxn new a mockTxn.
func newMockTxn() kv.Transaction {
	return &Txn{
		opts:  make(map[int]interface{}),
		valid: true,
	}
}
