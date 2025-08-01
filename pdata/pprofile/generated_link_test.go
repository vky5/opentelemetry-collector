// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestLink_MoveTo(t *testing.T) {
	ms := generateTestLink()
	dest := NewLink()
	ms.MoveTo(dest)
	assert.Equal(t, NewLink(), ms)
	assert.Equal(t, generateTestLink(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestLink(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newLink(&otlpprofiles.Link{}, &sharedState)) })
	assert.Panics(t, func() { newLink(&otlpprofiles.Link{}, &sharedState).MoveTo(dest) })
}

func TestLink_CopyTo(t *testing.T) {
	ms := NewLink()
	orig := NewLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestLink()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newLink(&otlpprofiles.Link{}, &sharedState)) })
}

func TestLink_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestLink()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	iter := json.BorrowIterator(stream.Buffer())
	defer json.ReturnIterator(iter)
	dest := NewLink()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestLink_TraceID(t *testing.T) {
	ms := NewLink()
	assert.Equal(t, pcommon.TraceID(data.TraceID([16]byte{})), ms.TraceID())
	testValTraceID := pcommon.TraceID(data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestLink_SpanID(t *testing.T) {
	ms := NewLink()
	assert.Equal(t, pcommon.SpanID(data.SpanID([8]byte{})), ms.SpanID())
	testValSpanID := pcommon.SpanID(data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func generateTestLink() Link {
	tv := NewLink()
	fillTestLink(tv)
	return tv
}

func fillTestLink(tv Link) {
	tv.orig.TraceId = data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.SpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
}
