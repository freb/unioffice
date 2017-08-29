// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WordprocessingContentPartNonVisual struct {
	CNvPr            *drawingml.CT_NonVisualDrawingProps
	CNvContentPartPr *drawingml.CT_NonVisualContentPartProperties
}

func NewCT_WordprocessingContentPartNonVisual() *CT_WordprocessingContentPartNonVisual {
	ret := &CT_WordprocessingContentPartNonVisual{}
	return ret
}
func (m *CT_WordprocessingContentPartNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.CNvPr != nil {
		secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
		e.EncodeElement(m.CNvPr, secNvPr)
	}
	if m.CNvContentPartPr != nil {
		secNvContentPartPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvContentPartPr"}}
		e.EncodeElement(m.CNvContentPartPr, secNvContentPartPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WordprocessingContentPartNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingContentPartNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvContentPartPr":
				m.CNvContentPartPr = drawingml.NewCT_NonVisualContentPartProperties()
				if err := d.DecodeElement(m.CNvContentPartPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingContentPartNonVisual
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_WordprocessingContentPartNonVisual) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingContentPartNonVisual")
}
func (m *CT_WordprocessingContentPartNonVisual) ValidateWithPath(path string) error {
	if m.CNvPr != nil {
		if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
			return err
		}
	}
	if m.CNvContentPartPr != nil {
		if err := m.CNvContentPartPr.ValidateWithPath(path + "/CNvContentPartPr"); err != nil {
			return err
		}
	}
	return nil
}