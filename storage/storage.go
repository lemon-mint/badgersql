package storage

import (
	"context"

	"github.com/lemon-mint/tidb/v7/kv"
	deadlockpb "github.com/pingcap/kvproto/pkg/deadlock"
	"github.com/tikv/client-go/v2/oracle"
	"github.com/tikv/client-go/v2/tikv"
)

// Storage is used to start a must commit-failed txn.
type Storage struct{}

func (s *Storage) GetCodec() tikv.Codec {
	return nil
}

func (s *Storage) Begin(opts ...tikv.TxnOption) (kv.Transaction, error) {
	return newMockTxn(), nil
}

func (*Txn) IsPessimistic() bool {
	return false
}

func (s *Storage) GetSnapshot(ver kv.Version) kv.Snapshot {
	return &Snapshot{
		store: newMockMap(),
	}
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) UUID() string {
	return ""
}

// CurrentVersion returns current max committed version.
func (s *Storage) CurrentVersion(txnScope string) (kv.Version, error) {
	return kv.NewVersion(1), nil
}

func (s *Storage) GetClient() kv.Client {
	return nil
}

func (s *Storage) GetMPPClient() kv.MPPClient {
	return nil
}

func (s *Storage) GetOracle() oracle.Oracle {
	return nil
}

func (s *Storage) SupportDeleteRange() (supported bool) {
	return false
}

func (s *Storage) Name() string {
	return "KVMockStorage"
}

func (s *Storage) Describe() string {
	return "KVMockStorage is a mock Store implementation, only for unittests in KV package"
}

func (s *Storage) ShowStatus(ctx context.Context, key string) (interface{}, error) {
	return nil, nil
}

func (s *Storage) GetMemCache() kv.MemManager {
	return nil
}

func (s *Storage) GetLockWaits() ([]*deadlockpb.WaitForEntry, error) {
	return nil, nil
}

func (s *Storage) GetMinSafeTS(txnScope string) uint64 {
	return 0
}

// newMockStorage creates a new mockStorage.
func newMockStorage() kv.Storage {
	return &Storage{}
}
