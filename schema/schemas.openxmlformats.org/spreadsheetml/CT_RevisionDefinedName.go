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

	"baliance.com/gooxml"
)

type CT_RevisionDefinedName struct {
	// Local Name Sheet Id
	LocalSheetIdAttr *uint32
	// Custom View
	CustomViewAttr *bool
	// Name
	NameAttr string
	// Function
	FunctionAttr *bool
	// Old Function
	OldFunctionAttr *bool
	// Function Group Id
	FunctionGroupIdAttr *uint8
	// Old Function Group Id
	OldFunctionGroupIdAttr *uint8
	// Shortcut Key
	ShortcutKeyAttr *uint8
	// Old Short Cut Key
	OldShortcutKeyAttr *uint8
	// Named Range Hidden
	HiddenAttr *bool
	// Old Hidden
	OldHiddenAttr *bool
	// New Custom Menu
	CustomMenuAttr *string
	// Old Custom Menu Text
	OldCustomMenuAttr *string
	// Description
	DescriptionAttr *string
	// Old Description
	OldDescriptionAttr *string
	// New Help Topic
	HelpAttr *string
	// Old Help Topic
	OldHelpAttr *string
	// Status Bar
	StatusBarAttr *string
	// Old Status Bar
	OldStatusBarAttr *string
	// Name Comment
	CommentAttr *string
	// Old Name Comment
	OldCommentAttr *string
	// Formula
	Formula *string
	// Old Formula
	OldFormula *string
	ExtLst     *CT_ExtensionList
	RIdAttr    *uint32
	UaAttr     *bool
	RaAttr     *bool
}

func NewCT_RevisionDefinedName() *CT_RevisionDefinedName {
	ret := &CT_RevisionDefinedName{}
	return ret
}
func (m *CT_RevisionDefinedName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.LocalSheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "localSheetId"},
			Value: fmt.Sprintf("%v", *m.LocalSheetIdAttr)})
	}
	if m.CustomViewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customView"},
			Value: fmt.Sprintf("%v", *m.CustomViewAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.FunctionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "function"},
			Value: fmt.Sprintf("%v", *m.FunctionAttr)})
	}
	if m.OldFunctionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldFunction"},
			Value: fmt.Sprintf("%v", *m.OldFunctionAttr)})
	}
	if m.FunctionGroupIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "functionGroupId"},
			Value: fmt.Sprintf("%v", *m.FunctionGroupIdAttr)})
	}
	if m.OldFunctionGroupIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldFunctionGroupId"},
			Value: fmt.Sprintf("%v", *m.OldFunctionGroupIdAttr)})
	}
	if m.ShortcutKeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shortcutKey"},
			Value: fmt.Sprintf("%v", *m.ShortcutKeyAttr)})
	}
	if m.OldShortcutKeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldShortcutKey"},
			Value: fmt.Sprintf("%v", *m.OldShortcutKeyAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	if m.OldHiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldHidden"},
			Value: fmt.Sprintf("%v", *m.OldHiddenAttr)})
	}
	if m.CustomMenuAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customMenu"},
			Value: fmt.Sprintf("%v", *m.CustomMenuAttr)})
	}
	if m.OldCustomMenuAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldCustomMenu"},
			Value: fmt.Sprintf("%v", *m.OldCustomMenuAttr)})
	}
	if m.DescriptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "description"},
			Value: fmt.Sprintf("%v", *m.DescriptionAttr)})
	}
	if m.OldDescriptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldDescription"},
			Value: fmt.Sprintf("%v", *m.OldDescriptionAttr)})
	}
	if m.HelpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "help"},
			Value: fmt.Sprintf("%v", *m.HelpAttr)})
	}
	if m.OldHelpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldHelp"},
			Value: fmt.Sprintf("%v", *m.OldHelpAttr)})
	}
	if m.StatusBarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "statusBar"},
			Value: fmt.Sprintf("%v", *m.StatusBarAttr)})
	}
	if m.OldStatusBarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldStatusBar"},
			Value: fmt.Sprintf("%v", *m.OldStatusBarAttr)})
	}
	if m.CommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "comment"},
			Value: fmt.Sprintf("%v", *m.CommentAttr)})
	}
	if m.OldCommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oldComment"},
			Value: fmt.Sprintf("%v", *m.OldCommentAttr)})
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%v", *m.UaAttr)})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%v", *m.RaAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Formula != nil {
		seformula := xml.StartElement{Name: xml.Name{Local: "x:formula"}}
		gooxml.AddPreserveSpaceAttr(&seformula, *m.Formula)
		e.EncodeElement(m.Formula, seformula)
	}
	if m.OldFormula != nil {
		seoldFormula := xml.StartElement{Name: xml.Name{Local: "x:oldFormula"}}
		gooxml.AddPreserveSpaceAttr(&seoldFormula, *m.OldFormula)
		e.EncodeElement(m.OldFormula, seoldFormula)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RevisionDefinedName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "localSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.LocalSheetIdAttr = &pt
		}
		if attr.Name.Local == "customView" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomViewAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "function" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FunctionAttr = &parsed
		}
		if attr.Name.Local == "oldFunction" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OldFunctionAttr = &parsed
		}
		if attr.Name.Local == "functionGroupId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.FunctionGroupIdAttr = &pt
		}
		if attr.Name.Local == "oldFunctionGroupId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.OldFunctionGroupIdAttr = &pt
		}
		if attr.Name.Local == "shortcutKey" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.ShortcutKeyAttr = &pt
		}
		if attr.Name.Local == "oldShortcutKey" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.OldShortcutKeyAttr = &pt
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "oldHidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OldHiddenAttr = &parsed
		}
		if attr.Name.Local == "customMenu" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CustomMenuAttr = &parsed
		}
		if attr.Name.Local == "oldCustomMenu" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldCustomMenuAttr = &parsed
		}
		if attr.Name.Local == "description" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DescriptionAttr = &parsed
		}
		if attr.Name.Local == "oldDescription" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldDescriptionAttr = &parsed
		}
		if attr.Name.Local == "help" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HelpAttr = &parsed
		}
		if attr.Name.Local == "oldHelp" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldHelpAttr = &parsed
		}
		if attr.Name.Local == "statusBar" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StatusBarAttr = &parsed
		}
		if attr.Name.Local == "oldStatusBar" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldStatusBarAttr = &parsed
		}
		if attr.Name.Local == "comment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommentAttr = &parsed
		}
		if attr.Name.Local == "oldComment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OldCommentAttr = &parsed
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.RIdAttr = &pt
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
		}
	}
lCT_RevisionDefinedName:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "formula":
				m.Formula = new(string)
				if err := d.DecodeElement(m.Formula, &el); err != nil {
					return err
				}
			case "oldFormula":
				m.OldFormula = new(string)
				if err := d.DecodeElement(m.OldFormula, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionDefinedName
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_RevisionDefinedName) Validate() error {
	return m.ValidateWithPath("CT_RevisionDefinedName")
}
func (m *CT_RevisionDefinedName) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}