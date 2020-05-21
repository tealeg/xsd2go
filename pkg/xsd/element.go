package xsd

import (
	"encoding/xml"
)

// Element defines single XML element
type Element struct {
	XMLName     xml.Name     `xml:"http://www.w3.org/2001/XMLSchema element"`
	Name        string       `xml:"name,attr"`
	ComplexType *ComplexType `xml:"complexType"`
}

func (e *Element) Attributes() []Attribute {
	if e.ComplexType != nil {
		return e.ComplexType.Attributes
	}
	return []Attribute{}
}

func (e *Element) compile(s *Schema) {
	if e.ComplexType != nil {
		// Handle improbable name clash. Consider XSD defining two attributes on the element:
		// "id" and "Id", this would create name clash given the camelization we do.
		goNames := map[string]uint{}
		for idx, _ := range e.ComplexType.Attributes {
			attribute := &e.ComplexType.Attributes[idx]

			count := goNames[attribute.GoName()]
			count += 1
			goNames[attribute.GoName()] = count
			attribute.DuplicateCount = count
			// Second GoName may be different depending on the DuplicateCount
			goNames[attribute.GoName()] = count
		}
	}
}
