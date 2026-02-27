package linearprob_test

import (
	"testing"

	"github.com/thumbrise/demo/internal/structures/hashtable/store"
	"github.com/thumbrise/demo/internal/structures/hashtable/store/types/openaddr/linearprob"
)

func TestStore_Set(t *testing.T) {
	stor := linearprob.NewStore[int](10)

	res := stor.Set(&store.HashedItem[int]{
		Key:   "key",
		Hash:  0,
		Value: 1,
	})
	if !res {
		t.Error("expected true, got false")
	}
}

func TestStore_Get(t *testing.T) {
	stor := linearprob.NewStore[int](10)

	res := stor.Set(&store.HashedItem[int]{
		Key:   "key",
		Value: 1,
	})
	if !res {
		t.Error("expected true, got false")
	}

	getRes := stor.Get(&store.HashedItem[int]{
		Key: "key",
	})
	if getRes.GetValue() != 1 {
		t.Error("expected 1, got", getRes.GetValue())
	}
}

func TestStore_Delete(t *testing.T) {
	stor := linearprob.NewStore[int](10)

	res := stor.Set(&store.HashedItem[int]{
		Key:   "key",
		Value: 1,
	})
	if !res {
		t.Error("expected true, got false")
	}

	deleteRes := stor.Delete(&store.HashedItem[int]{
		Key: "key",
	})
	if !deleteRes {
		t.Error("expected true, got false")
	}
}
