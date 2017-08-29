// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_NumVal struct {
	IdxAttr        uint32
	FormatCodeAttr *string
	V              string
}

func NewCT_NumVal() *CT_NumVal {
	ret := &CT_NumVal{}
	return ret
}
func (m *CT_NumVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	if m.FormatCodeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCode"},
			Value: fmt.Sprintf("%v", *m.FormatCodeAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sev := xml.StartElement{Name: xml.Name{Local: "v"}}
	gooxml.AddPreserveSpaceAttr(&sev, m.V)
	e.EncodeElement(m.V, sev)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NumVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
		}
		if attr.Name.Local == "formatCode" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatCodeAttr = &parsed
		}
	}
lCT_NumVal:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "v":
				if err := d.DecodeElement(m.V, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumVal
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NumVal) Validate() error {
	return m.ValidateWithPath("CT_NumVal")
}
func (m *CT_NumVal) ValidateWithPath(path string) error {
	return nil
}