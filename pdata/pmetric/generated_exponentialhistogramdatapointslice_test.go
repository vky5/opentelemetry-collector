// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestExponentialHistogramDataPointSlice(t *testing.T) {
	es := NewExponentialHistogramDataPointSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newExponentialHistogramDataPointSlice(&[]*otlpmetrics.ExponentialHistogramDataPoint{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewExponentialHistogramDataPoint()
	testVal := generateTestExponentialHistogramDataPoint()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestExponentialHistogramDataPoint(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestExponentialHistogramDataPointSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newExponentialHistogramDataPointSlice(&[]*otlpmetrics.ExponentialHistogramDataPoint{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewExponentialHistogramDataPointSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestExponentialHistogramDataPointSlice_CopyTo(t *testing.T) {
	dest := NewExponentialHistogramDataPointSlice()
	// Test CopyTo to empty
	NewExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, NewExponentialHistogramDataPointSlice(), dest)

	// Test CopyTo larger slice
	generateTestExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)

	// Test CopyTo same size slice
	generateTestExponentialHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
}

func TestExponentialHistogramDataPointSlice_EnsureCapacity(t *testing.T) {
	es := generateTestExponentialHistogramDataPointSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestExponentialHistogramDataPointSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), es)
}

func TestExponentialHistogramDataPointSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestExponentialHistogramDataPointSlice()
	dest := NewExponentialHistogramDataPointSlice()
	src := generateTestExponentialHistogramDataPointSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestExponentialHistogramDataPointSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}

	dest.MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestExponentialHistogramDataPointSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewExponentialHistogramDataPointSlice()
	emptySlice.RemoveIf(func(el ExponentialHistogramDataPoint) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestExponentialHistogramDataPointSlice()
	pos := 0
	filtered.RemoveIf(func(el ExponentialHistogramDataPoint) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestExponentialHistogramDataPointSliceAll(t *testing.T) {
	ms := generateTestExponentialHistogramDataPointSlice()
	assert.NotEmpty(t, ms.Len())

	var c int
	for i, v := range ms.All() {
		assert.Equal(t, ms.At(i), v, "element should match")
		c++
	}
	assert.Equal(t, ms.Len(), c, "All elements should have been visited")
}

func TestExponentialHistogramDataPointSlice_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestExponentialHistogramDataPointSlice()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	iter := json.BorrowIterator(stream.Buffer())
	defer json.ReturnIterator(iter)
	dest := NewExponentialHistogramDataPointSlice()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestExponentialHistogramDataPointSlice_Sort(t *testing.T) {
	es := generateTestExponentialHistogramDataPointSlice()
	es.Sort(func(a, b ExponentialHistogramDataPoint) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Less(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b ExponentialHistogramDataPoint) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Greater(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestExponentialHistogramDataPointSlice() ExponentialHistogramDataPointSlice {
	es := NewExponentialHistogramDataPointSlice()
	fillTestExponentialHistogramDataPointSlice(es)
	return es
}

func fillTestExponentialHistogramDataPointSlice(es ExponentialHistogramDataPointSlice) {
	*es.orig = make([]*otlpmetrics.ExponentialHistogramDataPoint, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpmetrics.ExponentialHistogramDataPoint{}
		fillTestExponentialHistogramDataPoint(newExponentialHistogramDataPoint((*es.orig)[i], es.state))
	}
}
