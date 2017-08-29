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

type CT_Members struct {
	// Item Count
	CountAttr *uint32
	// Hierarchy Level
	LevelAttr *uint32
	// Member
	Member []*CT_Member
}

func NewCT_Members() *CT_Members {
	ret := &CT_Members{}
	return ret
}
func (m *CT_Members) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.LevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "level"},
			Value: fmt.Sprintf("%v", *m.LevelAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	semember := xml.StartElement{Name: xml.Name{Local: "x:member"}}
	e.EncodeElement(m.Member, semember)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Members) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "level" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.LevelAttr = &pt
		}
	}
lCT_Members:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "member":
				tmp := NewCT_Member()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Member = append(m.Member, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Members
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Members) Validate() error {
	return m.ValidateWithPath("CT_Members")
}
func (m *CT_Members) ValidateWithPath(path string) error {
	for i, v := range m.Member {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Member[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}