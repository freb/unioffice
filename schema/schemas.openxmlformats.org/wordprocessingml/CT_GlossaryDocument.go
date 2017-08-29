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

type CT_GlossaryDocument struct {
	// Document Background
	Background *CT_Background
	DocParts   *CT_DocParts
}

func NewCT_GlossaryDocument() *CT_GlossaryDocument {
	ret := &CT_GlossaryDocument{}
	return ret
}
func (m *CT_GlossaryDocument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Background != nil {
		sebackground := xml.StartElement{Name: xml.Name{Local: "w:background"}}
		e.EncodeElement(m.Background, sebackground)
	}
	if m.DocParts != nil {
		sedocParts := xml.StartElement{Name: xml.Name{Local: "w:docParts"}}
		e.EncodeElement(m.DocParts, sedocParts)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GlossaryDocument) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GlossaryDocument:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "background":
				m.Background = NewCT_Background()
				if err := d.DecodeElement(m.Background, &el); err != nil {
					return err
				}
			case "docParts":
				m.DocParts = NewCT_DocParts()
				if err := d.DecodeElement(m.DocParts, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GlossaryDocument
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GlossaryDocument) Validate() error {
	return m.ValidateWithPath("CT_GlossaryDocument")
}
func (m *CT_GlossaryDocument) ValidateWithPath(path string) error {
	if m.Background != nil {
		if err := m.Background.ValidateWithPath(path + "/Background"); err != nil {
			return err
		}
	}
	if m.DocParts != nil {
		if err := m.DocParts.ValidateWithPath(path + "/DocParts"); err != nil {
			return err
		}
	}
	return nil
}