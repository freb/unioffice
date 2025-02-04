// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.
package spreadsheet

import (
	"fmt"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/sml"
)

// SharedStrings is a shared strings table, where string data can be placed
// outside of the sheet contents and referenced from a sheet.
type SharedStrings struct {
	x         *sml.Sst
	cachedIDs map[string]int
}

// NewSharedStrings constructs a new Shared Strings table.
func NewSharedStrings() SharedStrings {
	return SharedStrings{x: sml.NewSst(),
		cachedIDs: make(map[string]int)}
}

// X returns the inner wrapped XML type.
func (s SharedStrings) X() *sml.Sst {
	return s.x
}

// AddString adds a string to the shared string cache.
func (s SharedStrings) AddString(v string) int {
	if id, ok := s.cachedIDs[v]; ok {
		return id
	}
	rst := sml.NewCT_Rst()
	rst.T = unioffice.String(v)
	s.x.Si = append(s.x.Si, rst)
	id := len(s.x.Si) - 1
	s.cachedIDs[v] = id
	s.x.CountAttr = unioffice.Uint32(uint32(len(s.x.Si)))
	s.x.UniqueCountAttr = s.x.CountAttr
	return id
}

// GetString retrieves a string from the shared strings table by index.
func (s SharedStrings) GetString(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("invalid string index %d, must be > 0", id)
	}
	if id >= len(s.x.Si) { // Use >= instead of > to prevent out-of-bounds panic
		return "", fmt.Errorf("invalid string index %d, table only has %d values", id, len(s.x.Si))
	}
	si := s.x.Si[id]
	// If si.T is present, return the plain text directly
	if si.T != nil {
		return *si.T, nil
	}
	// If si.R exists, concatenate all rich text runs
	if len(si.R) > 0 {
		var sb strings.Builder
		for _, r := range si.R {
			sb.WriteString(r.T)
		}
		return sb.String(), nil
	}

	// No valid text found
	return "", nil
}
