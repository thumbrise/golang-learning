package search

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/thumbrise/golang-learning/internal/database/storage/core"
	"github.com/thumbrise/golang-learning/internal/database/storage/stats"
)

var ErrNotSupported = errors.New("SeqSearcher: method not supported")

type SeqSearcher[TRecord core.Record] struct {
	heap *core.Heap[TRecord]
}

func NewSeqSearcher[TRecord core.Record](heap *core.Heap[TRecord]) *SeqSearcher[TRecord] {
	return &SeqSearcher[TRecord]{
		heap: heap,
	}
}

func (d *SeqSearcher[TRecord]) Search(fieldName string, value string) []string {
	var result []string

	d.heap.Iterate(func(rec TRecord) {
		if rec.Get(fieldName) == value {
			result = append(result, rec.PK())
		}
	})

	return result
}

func (d *SeqSearcher[TRecord]) Insert(ctid string, fieldName string, value string) {
	err := fmt.Errorf("%w: Insert", ErrNotSupported)
	slog.Warn(err.Error())
}

func (d *SeqSearcher[TRecord]) Delete(fieldName string, value string) {
	err := fmt.Errorf("%w: Delete", ErrNotSupported)
	slog.Warn(err.Error())
}

func (d *SeqSearcher[TRecord]) DeleteCTID(ctid string, fieldName string, value string) {
	err := fmt.Errorf("%w: DeleteCTID", ErrNotSupported)
	slog.Warn(err.Error())
}

func (d *SeqSearcher[TRecord]) Update(ctid string, fieldName string, oldValue string, newValue string) {
	err := fmt.Errorf("%w: Update", ErrNotSupported)
	slog.Warn(err.Error())
}

func (d *SeqSearcher[TRecord]) SizeBytes() int {
	err := fmt.Errorf("%w: SizeBytes", ErrNotSupported)
	slog.Warn(err.Error())

	return 0
}

func (d *SeqSearcher[TRecord]) Depth() int {
	err := fmt.Errorf("%w: Depth", ErrNotSupported)
	slog.Warn(err.Error())

	return 0
}

func (d *SeqSearcher[TRecord]) Stats() *stats.IndexStats {
	rows := uint32(d.heap.Len())

	return &stats.IndexStats{
		Cost: 1 * rows,
		Rows: rows,
	}
}

func (d *SeqSearcher[TRecord]) Type() string {
	return d.String()
}

func (d *SeqSearcher[TRecord]) String() string {
	return "SeqScan"
}
