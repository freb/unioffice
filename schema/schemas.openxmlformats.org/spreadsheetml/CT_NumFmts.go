// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_NumFmts struct {
	// Number Format Count
	CountAttr *uint32
	// Number Formats
	NumFmt []*CT_NumFmt
}

func NewCT_NumFmts() *CT_NumFmts {
	ret := &CT_NumFmts{}
	return ret
}
func (m *CT_NumFmts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "x:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumFmts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_NumFmts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "numFmt":
				tmp := NewCT_NumFmt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.NumFmt = append(m.NumFmt, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumFmts
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumFmts) Validate() error {
	return m.ValidateWithPath("CT_NumFmts")
}
func (m *CT_NumFmts) ValidateWithPath(path string) error {
	for i, v := range m.NumFmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/NumFmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}