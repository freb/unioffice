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
	"strconv"
	"time"
)

type CT_RangePr struct {
	// Source Data Set Beginning Range
	AutoStartAttr *bool
	// Source Data Ending Range
	AutoEndAttr *bool
	// Group By
	GroupByAttr ST_GroupBy
	// Numeric Grouping Start Value
	StartNumAttr *float64
	// Numeric Grouping End Value
	EndNumAttr *float64
	// Date Grouping Start Value
	StartDateAttr *time.Time
	// Date Grouping End Value
	EndDateAttr *time.Time
	// Grouping Interval
	GroupIntervalAttr *float64
}

func NewCT_RangePr() *CT_RangePr {
	ret := &CT_RangePr{}
	return ret
}
func (m *CT_RangePr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.AutoStartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoStart"},
			Value: fmt.Sprintf("%v", *m.AutoStartAttr)})
	}
	if m.AutoEndAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoEnd"},
			Value: fmt.Sprintf("%v", *m.AutoEndAttr)})
	}
	if m.GroupByAttr != ST_GroupByUnset {
		attr, err := m.GroupByAttr.MarshalXMLAttr(xml.Name{Local: "groupBy"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StartNumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "startNum"},
			Value: fmt.Sprintf("%v", *m.StartNumAttr)})
	}
	if m.EndNumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endNum"},
			Value: fmt.Sprintf("%v", *m.EndNumAttr)})
	}
	if m.StartDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "startDate"},
			Value: fmt.Sprintf("%v", *m.StartDateAttr)})
	}
	if m.EndDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "endDate"},
			Value: fmt.Sprintf("%v", *m.EndDateAttr)})
	}
	if m.GroupIntervalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "groupInterval"},
			Value: fmt.Sprintf("%v", *m.GroupIntervalAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RangePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "autoStart" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoStartAttr = &parsed
		}
		if attr.Name.Local == "autoEnd" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoEndAttr = &parsed
		}
		if attr.Name.Local == "groupBy" {
			m.GroupByAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "startNum" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.StartNumAttr = &parsed
		}
		if attr.Name.Local == "endNum" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.EndNumAttr = &parsed
		}
		if attr.Name.Local == "startDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.StartDateAttr = &parsed
		}
		if attr.Name.Local == "endDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.EndDateAttr = &parsed
		}
		if attr.Name.Local == "groupInterval" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.GroupIntervalAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RangePr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RangePr) Validate() error {
	return m.ValidateWithPath("CT_RangePr")
}
func (m *CT_RangePr) ValidateWithPath(path string) error {
	if err := m.GroupByAttr.ValidateWithPath(path + "/GroupByAttr"); err != nil {
		return err
	}
	return nil
}