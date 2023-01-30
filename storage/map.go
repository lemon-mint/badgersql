package storage

import (
	"context"

	"github.com/lemon-mint/tidb/v7/kv"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
)

type Map struct {
	index []kv.Key
	value [][]byte
}

func newMockMap() *Map {
	return &Map{
		index: make([]kv.Key, 0),
		value: make([][]byte, 0),
	}
}

func (s *Map) SetDiskFullOpt(level kvrpcpb.DiskFullOpt) {
	//TODO nothing.
}

func (s *Map) Iter(kv.Key, kv.Key) (kv.Iterator, error) {
	return nil, nil
}
func (s *Map) IterReverse(kv.Key) (kv.Iterator, error) {
	return nil, nil
}

func (s *Map) Get(_ context.Context, k kv.Key) ([]byte, error) {
	for i, key := range s.index {
		if key.Cmp(k) == 0 {
			return s.value[i], nil
		}
	}
	return nil, kv.ErrNotExist
}

func (s *Map) Set(k kv.Key, v []byte) error {
	for i, key := range s.index {
		if key.Cmp(k) == 0 {
			s.value[i] = v
			return nil
		}
	}
	s.index = append(s.index, k)
	s.value = append(s.value, v)
	return nil
}

func (s *Map) Delete(k kv.Key) error {
	for i, key := range s.index {
		if key.Cmp(k) == 0 {
			s.index = append(s.index[:i], s.index[i+1:]...)
			s.value = append(s.value[:i], s.value[i+1:]...)
			return nil
		}
	}
	return nil
}
