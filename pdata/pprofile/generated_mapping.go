// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// Mapping describes the mapping of a binary in memory, including its address range, file offset, and metadata like build ID
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewMapping function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Mapping struct {
	orig  *otlpprofiles.Mapping
	state *internal.State
}

func newMapping(orig *otlpprofiles.Mapping, state *internal.State) Mapping {
	return Mapping{orig: orig, state: state}
}

// NewMapping creates a new empty Mapping.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewMapping() Mapping {
	state := internal.StateMutable
	return newMapping(&otlpprofiles.Mapping{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Mapping) MoveTo(dest Mapping) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if ms.orig == dest.orig {
		return
	}
	*dest.orig = *ms.orig
	*ms.orig = otlpprofiles.Mapping{}
}

// MemoryStart returns the memorystart associated with this Mapping.
func (ms Mapping) MemoryStart() uint64 {
	return ms.orig.MemoryStart
}

// SetMemoryStart replaces the memorystart associated with this Mapping.
func (ms Mapping) SetMemoryStart(v uint64) {
	ms.state.AssertMutable()
	ms.orig.MemoryStart = v
}

// MemoryLimit returns the memorylimit associated with this Mapping.
func (ms Mapping) MemoryLimit() uint64 {
	return ms.orig.MemoryLimit
}

// SetMemoryLimit replaces the memorylimit associated with this Mapping.
func (ms Mapping) SetMemoryLimit(v uint64) {
	ms.state.AssertMutable()
	ms.orig.MemoryLimit = v
}

// FileOffset returns the fileoffset associated with this Mapping.
func (ms Mapping) FileOffset() uint64 {
	return ms.orig.FileOffset
}

// SetFileOffset replaces the fileoffset associated with this Mapping.
func (ms Mapping) SetFileOffset(v uint64) {
	ms.state.AssertMutable()
	ms.orig.FileOffset = v
}

// FilenameStrindex returns the filenamestrindex associated with this Mapping.
func (ms Mapping) FilenameStrindex() int32 {
	return ms.orig.FilenameStrindex
}

// SetFilenameStrindex replaces the filenamestrindex associated with this Mapping.
func (ms Mapping) SetFilenameStrindex(v int32) {
	ms.state.AssertMutable()
	ms.orig.FilenameStrindex = v
}

// AttributeIndices returns the AttributeIndices associated with this Mapping.
func (ms Mapping) AttributeIndices() pcommon.Int32Slice {
	return pcommon.Int32Slice(internal.NewInt32Slice(&ms.orig.AttributeIndices, ms.state))
}

// HasFunctions returns the hasfunctions associated with this Mapping.
func (ms Mapping) HasFunctions() bool {
	return ms.orig.HasFunctions
}

// SetHasFunctions replaces the hasfunctions associated with this Mapping.
func (ms Mapping) SetHasFunctions(v bool) {
	ms.state.AssertMutable()
	ms.orig.HasFunctions = v
}

// HasFilenames returns the hasfilenames associated with this Mapping.
func (ms Mapping) HasFilenames() bool {
	return ms.orig.HasFilenames
}

// SetHasFilenames replaces the hasfilenames associated with this Mapping.
func (ms Mapping) SetHasFilenames(v bool) {
	ms.state.AssertMutable()
	ms.orig.HasFilenames = v
}

// HasLineNumbers returns the haslinenumbers associated with this Mapping.
func (ms Mapping) HasLineNumbers() bool {
	return ms.orig.HasLineNumbers
}

// SetHasLineNumbers replaces the haslinenumbers associated with this Mapping.
func (ms Mapping) SetHasLineNumbers(v bool) {
	ms.state.AssertMutable()
	ms.orig.HasLineNumbers = v
}

// HasInlineFrames returns the hasinlineframes associated with this Mapping.
func (ms Mapping) HasInlineFrames() bool {
	return ms.orig.HasInlineFrames
}

// SetHasInlineFrames replaces the hasinlineframes associated with this Mapping.
func (ms Mapping) SetHasInlineFrames(v bool) {
	ms.state.AssertMutable()
	ms.orig.HasInlineFrames = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Mapping) CopyTo(dest Mapping) {
	dest.state.AssertMutable()
	copyOrigMapping(dest.orig, ms.orig)
}

// marshalJSONStream marshals all properties from the current struct to the destination stream.
func (ms Mapping) marshalJSONStream(dest *json.Stream) {
	dest.WriteObjectStart()
	if ms.orig.MemoryStart != uint64(0) {
		dest.WriteObjectField("memoryStart")
		dest.WriteUint64(ms.orig.MemoryStart)
	}
	if ms.orig.MemoryLimit != uint64(0) {
		dest.WriteObjectField("memoryLimit")
		dest.WriteUint64(ms.orig.MemoryLimit)
	}
	if ms.orig.FileOffset != uint64(0) {
		dest.WriteObjectField("fileOffset")
		dest.WriteUint64(ms.orig.FileOffset)
	}
	if ms.orig.FilenameStrindex != int32(0) {
		dest.WriteObjectField("filenameStrindex")
		dest.WriteInt32(ms.orig.FilenameStrindex)
	}
	if len(ms.orig.AttributeIndices) > 0 {
		dest.WriteObjectField("attributeIndices")
		internal.MarshalJSONStreamInt32Slice(internal.NewInt32Slice(&ms.orig.AttributeIndices, ms.state), dest)
	}
	if ms.orig.HasFunctions != false {
		dest.WriteObjectField("hasFunctions")
		dest.WriteBool(ms.orig.HasFunctions)
	}
	if ms.orig.HasFilenames != false {
		dest.WriteObjectField("hasFilenames")
		dest.WriteBool(ms.orig.HasFilenames)
	}
	if ms.orig.HasLineNumbers != false {
		dest.WriteObjectField("hasLineNumbers")
		dest.WriteBool(ms.orig.HasLineNumbers)
	}
	if ms.orig.HasInlineFrames != false {
		dest.WriteObjectField("hasInlineFrames")
		dest.WriteBool(ms.orig.HasInlineFrames)
	}
	dest.WriteObjectEnd()
}

func copyOrigMapping(dest, src *otlpprofiles.Mapping) {
	dest.MemoryStart = src.MemoryStart
	dest.MemoryLimit = src.MemoryLimit
	dest.FileOffset = src.FileOffset
	dest.FilenameStrindex = src.FilenameStrindex
	dest.AttributeIndices = internal.CopyOrigInt32Slice(dest.AttributeIndices, src.AttributeIndices)
	dest.HasFunctions = src.HasFunctions
	dest.HasFilenames = src.HasFilenames
	dest.HasLineNumbers = src.HasLineNumbers
	dest.HasInlineFrames = src.HasInlineFrames
}
