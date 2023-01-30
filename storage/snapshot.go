package storage

import (
	"context"

	"github.com/lemon-mint/tidb/v7/kv"
	deadlockpb "github.com/pingcap/kvproto/pkg/deadlock"
)

type Snapshot struct {
	store kv.Retriever
}

func (s *Snapshot) Get(ctx context.Context, k kv.Key) ([]byte, error) {
	return s.store.Get(ctx, k)
}

func (s *Snapshot) SetPriority(priority int) {
}

func (s *Snapshot) BatchGet(ctx context.Context, keys []kv.Key) (map[string][]byte, error) {
	m := make(map[string][]byte, len(keys))
	for _, k := range keys {
		v, err := s.store.Get(ctx, k)
		if kv.IsErrNotFound(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		m[string(k)] = v
	}
	return m, nil
}

func (s *Snapshot) Iter(k kv.Key, upperBound kv.Key) (kv.Iterator, error) {
	return s.store.Iter(k, upperBound)
}

func (s *Snapshot) IterReverse(k kv.Key) (kv.Iterator, error) {
	return s.store.IterReverse(k)
}

func (s *Snapshot) SetOption(opt int, val interface{}) {}

func (s *Snapshot) GetLockWaits() []deadlockpb.WaitForEntry {
	return nil
}
