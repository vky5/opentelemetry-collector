// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"iter"
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

// AttributeUnitSlice logically represents a slice of AttributeUnit.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewAttributeUnitSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type AttributeUnitSlice struct {
	orig  *[]*otlpprofiles.AttributeUnit
	state *internal.State
}

func newAttributeUnitSlice(orig *[]*otlpprofiles.AttributeUnit, state *internal.State) AttributeUnitSlice {
	return AttributeUnitSlice{orig: orig, state: state}
}

// NewAttributeUnitSlice creates a AttributeUnitSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewAttributeUnitSlice() AttributeUnitSlice {
	orig := []*otlpprofiles.AttributeUnit(nil)
	state := internal.StateMutable
	return newAttributeUnitSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewAttributeUnitSlice()".
func (es AttributeUnitSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es AttributeUnitSlice) At(i int) AttributeUnit {
	return newAttributeUnit((*es.orig)[i], es.state)
}

// All returns an iterator over index-value pairs in the slice.
//
//	for i, v := range es.All() {
//	    ... // Do something with index-value pair
//	}
func (es AttributeUnitSlice) All() iter.Seq2[int, AttributeUnit] {
	return func(yield func(int, AttributeUnit) bool) {
		for i := 0; i < es.Len(); i++ {
			if !yield(i, es.At(i)) {
				return
			}
		}
	}
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new AttributeUnitSlice can be initialized:
//
//	es := NewAttributeUnitSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es AttributeUnitSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlpprofiles.AttributeUnit, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty AttributeUnit.
// It returns the newly added AttributeUnit.
func (es AttributeUnitSlice) AppendEmpty() AttributeUnit {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlpprofiles.AttributeUnit{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es AttributeUnitSlice) MoveAndAppendTo(dest AttributeUnitSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if es.orig == dest.orig {
		return
	}
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es AttributeUnitSlice) RemoveIf(f func(AttributeUnit) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es AttributeUnitSlice) CopyTo(dest AttributeUnitSlice) {
	dest.state.AssertMutable()
	*dest.orig = copyOrigAttributeUnitSlice(*dest.orig, *es.orig)
}

// Sort sorts the AttributeUnit elements within AttributeUnitSlice given the
// provided less function so that two instances of AttributeUnitSlice
// can be compared.
func (es AttributeUnitSlice) Sort(less func(a, b AttributeUnit) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}

// marshalJSONStream marshals all properties from the current struct to the destination stream.
func (ms AttributeUnitSlice) marshalJSONStream(dest *json.Stream) {
	dest.WriteArrayStart()
	if len(*ms.orig) > 0 {
		ms.At(0).marshalJSONStream(dest)
	}
	for i := 1; i < len(*ms.orig); i++ {
		dest.WriteMore()
		ms.At(i).marshalJSONStream(dest)
	}
	dest.WriteArrayEnd()
}

// unmarshalJSONIter unmarshals all properties from the current struct from the source iterator.
func (ms AttributeUnitSlice) unmarshalJSONIter(iter *json.Iterator) {
	iter.ReadArrayCB(func(iter *json.Iterator) bool {
		*ms.orig = append(*ms.orig, &otlpprofiles.AttributeUnit{})
		ms.At(ms.Len() - 1).unmarshalJSONIter(iter)
		return true
	})
}

func copyOrigAttributeUnitSlice(dest, src []*otlpprofiles.AttributeUnit) []*otlpprofiles.AttributeUnit {
	if cap(dest) < len(src) {
		dest = make([]*otlpprofiles.AttributeUnit, len(src))
		data := make([]otlpprofiles.AttributeUnit, len(src))
		for i := range src {
			dest[i] = &data[i]
		}
	}
	dest = dest[:len(src)]
	for i := range src {
		copyOrigAttributeUnit(dest[i], src[i])
	}
	return dest
}
