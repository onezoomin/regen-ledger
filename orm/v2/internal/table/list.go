package table

import (
	"bytes"
	"fmt"

	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"

	"google.golang.org/protobuf/reflect/protoreflect"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"
	"github.com/regen-network/regen-ledger/orm/v2/orm"

	"google.golang.org/protobuf/proto"
)

func (s *TableModel) List(kvStore kv.ReadKVStore, opts *orm.ListOptions) orm.Iterator {
	if opts == nil {
		opts = &orm.ListOptions{}
	}

	var cdc ormkey.CodecI
	var idx *Index
	if opts.Index != "" {
		var ok bool
		idx, ok = s.IndexesByFields[opts.Index]
		if !ok {
			return orm.ErrIterator{Err: fmt.Errorf("can't find indexer %s", opts.Index)}
		}
		cdc = idx.Codec.Codec
	} else {
		cdc = s.PkCodec
	}

	var start, end []byte
	var err error
	if opts.Cursor != nil && !opts.Reverse {
		start = opts.Cursor
	} else {
		var prefix []protoreflect.Value
		if len(opts.Prefix) != 0 {
			if len(opts.Start) != 0 {
				return orm.ErrIterator{Err: ormerrors.InvalidListOptions}
			}
			prefix = opts.Prefix
		} else if len(opts.Start) != 0 {
			if len(opts.Prefix) != 0 {
				return orm.ErrIterator{Err: ormerrors.InvalidListOptions}
			}
			prefix = opts.Start
		}

		start, err = cdc.Encode(prefix)
		if err != nil {
			return orm.ErrIterator{Err: err}
		}
	}

	if opts.Cursor != nil && opts.Reverse {
		start = storetypes.PrefixEndBytes(opts.Cursor)
	} else {
		var prefix []protoreflect.Value
		if len(opts.Prefix) != 0 {
			if len(opts.End) != 0 {
				return orm.ErrIterator{Err: ormerrors.InvalidListOptions}
			}
			prefix = opts.Prefix
		} else if len(opts.End) != 0 {
			if len(opts.Prefix) != 0 {
				return orm.ErrIterator{Err: ormerrors.InvalidListOptions}
			}
			prefix = opts.End
		}

		end, err = cdc.Encode(prefix)
		if err != nil {
			return orm.ErrIterator{Err: err}
		}
		end = storetypes.PrefixEndBytes(end)
	}

	var iterator kv.KVStoreIterator
	if !opts.Reverse {
		iterator = kvStore.Iterator(start, end)
	} else {
		iterator = kvStore.ReverseIterator(start, end)
	}

	if idx != nil {
		return &idxIterator{
			kv:       kvStore,
			store:    s,
			iterator: iterator,
			start:    true,
			cdc:      idx.Codec,
		}
	} else {
		return &pkIterator{
			kv:       kvStore,
			store:    s,
			iterator: iterator,
			start:    true,
		}
	}
}

type pkIterator struct {
	orm.UnimplementedIterator

	kv       kv.ReadKVStore
	store    *TableModel
	iterator kv.KVStoreIterator
	start    bool
}

func (t *pkIterator) Close() {
	_ = t.iterator.Close()
}

func (t *pkIterator) isIterator() {}

func (t *pkIterator) Next(message proto.Message) (bool, error) {
	if t.start {
		t.start = false
	} else {
		t.iterator.Next()
	}

	if !t.iterator.Valid() {
		return false, nil
	}

	k := t.iterator.Key()
	pkValues, err := t.store.PkCodec.Decode(bytes.NewReader(k))
	if err != nil {
		return false, err
	}

	bz := t.iterator.Value()
	err = t.store.PkCodec.Unmarshal(pkValues, bz, message)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t pkIterator) Cursor() orm.Cursor {
	return t.iterator.Key()
}

type idxIterator struct {
	orm.UnimplementedIterator

	kv       kv.ReadKVStore
	store    *TableModel
	iterator kv.KVStoreIterator
	start    bool
	cdc      *ormkey.IndexKeyCodec
}

func (t *idxIterator) Close() {
	_ = t.iterator.Close()
}

func (t *idxIterator) isIterator() {}

func (t *idxIterator) Next(message proto.Message) (bool, error) {
	if t.start {
		t.start = false
	} else {
		t.iterator.Next()
	}

	if !t.iterator.Valid() {
		return false, nil
	}

	k := t.iterator.Key()
	pkValues, err := t.cdc.ReadPrimaryKey(bytes.NewReader(k))
	if err != nil {
		return false, err
	}

	pk, err := t.store.PkCodec.Encode(pkValues)
	if err != nil {
		return false, err
	}

	bz := t.kv.Get(pk)
	err = t.store.PkCodec.Unmarshal(pkValues, bz, message)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t idxIterator) Cursor() orm.Cursor {
	return t.iterator.Key()
}
