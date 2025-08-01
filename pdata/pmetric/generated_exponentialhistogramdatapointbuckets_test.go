// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestExponentialHistogramDataPointBuckets_MoveTo(t *testing.T) {
	ms := generateTestExponentialHistogramDataPointBuckets()
	dest := NewExponentialHistogramDataPointBuckets()
	ms.MoveTo(dest)
	assert.Equal(t, NewExponentialHistogramDataPointBuckets(), ms)
	assert.Equal(t, generateTestExponentialHistogramDataPointBuckets(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestExponentialHistogramDataPointBuckets(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		ms.MoveTo(newExponentialHistogramDataPointBuckets(&otlpmetrics.ExponentialHistogramDataPoint_Buckets{}, &sharedState))
	})
	assert.Panics(t, func() {
		newExponentialHistogramDataPointBuckets(&otlpmetrics.ExponentialHistogramDataPoint_Buckets{}, &sharedState).MoveTo(dest)
	})
}

func TestExponentialHistogramDataPointBuckets_CopyTo(t *testing.T) {
	ms := NewExponentialHistogramDataPointBuckets()
	orig := NewExponentialHistogramDataPointBuckets()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestExponentialHistogramDataPointBuckets()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		ms.CopyTo(newExponentialHistogramDataPointBuckets(&otlpmetrics.ExponentialHistogramDataPoint_Buckets{}, &sharedState))
	})
}

func TestExponentialHistogramDataPointBuckets_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestExponentialHistogramDataPointBuckets()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	iter := json.BorrowIterator(stream.Buffer())
	defer json.ReturnIterator(iter)
	dest := NewExponentialHistogramDataPointBuckets()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestExponentialHistogramDataPointBuckets_Offset(t *testing.T) {
	ms := NewExponentialHistogramDataPointBuckets()
	assert.Equal(t, int32(0), ms.Offset())
	ms.SetOffset(int32(909))
	assert.Equal(t, int32(909), ms.Offset())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() {
		newExponentialHistogramDataPointBuckets(&otlpmetrics.ExponentialHistogramDataPoint_Buckets{}, &sharedState).SetOffset(int32(909))
	})
}

func TestExponentialHistogramDataPointBuckets_BucketCounts(t *testing.T) {
	ms := NewExponentialHistogramDataPointBuckets()
	assert.Equal(t, []uint64(nil), ms.BucketCounts().AsRaw())
	ms.BucketCounts().FromRaw([]uint64{1, 2, 3})
	assert.Equal(t, []uint64{1, 2, 3}, ms.BucketCounts().AsRaw())
}

func generateTestExponentialHistogramDataPointBuckets() ExponentialHistogramDataPointBuckets {
	tv := NewExponentialHistogramDataPointBuckets()
	fillTestExponentialHistogramDataPointBuckets(tv)
	return tv
}

func fillTestExponentialHistogramDataPointBuckets(tv ExponentialHistogramDataPointBuckets) {
	tv.orig.Offset = int32(909)
	tv.orig.BucketCounts = []uint64{1, 2, 3}
}
