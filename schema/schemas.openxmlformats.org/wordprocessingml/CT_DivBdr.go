// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_DivBdr struct {
	// Top Border for HTML div
	Top *CT_Border
	// Left Border for HTML div
	Left *CT_Border
	// Bottom Border for HTML div
	Bottom *CT_Border
	// Right Border for HTML div
	Right *CT_Border
}

func NewCT_DivBdr() *CT_DivBdr {
	ret := &CT_DivBdr{}
	return ret
}
func (m *CT_DivBdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DivBdr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DivBdr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "top":
				m.Top = NewCT_Border()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case "left":
				m.Left = NewCT_Border()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case "bottom":
				m.Bottom = NewCT_Border()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case "right":
				m.Right = NewCT_Border()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DivBdr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DivBdr) Validate() error {
	return m.ValidateWithPath("CT_DivBdr")
}
func (m *CT_DivBdr) ValidateWithPath(path string) error {
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	return nil
}