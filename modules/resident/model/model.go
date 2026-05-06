package model

import "encoding/xml"

type Result struct {
	Total   int
	Country string
}

// CensusData đại diện cho thẻ gốc <census_data>
type CensusData struct {
	XMLName xml.Name `xml:"census_data"`
	Country string   `xml:"country,attr"` // Thuộc tính country của thẻ census_data
	Regions []Region `xml:"region"`       // Danh sách các vùng
}

// Region đại diện cho thẻ <region>
type Region struct {
	ID       string    `xml:"id,attr"`   // Thuộc tính id
	Name     string    `xml:"name,attr"` // Thuộc tính name
	Citizens []Citizen `xml:"citizen"`   // Danh sách người dân trong vùng
}

// Citizen đại diện cho thẻ <citizen> - thông tin chi tiết từng người
type Citizen struct {
	ID         string `xml:"id,attr"`    // Thuộc tính id
	Name       string `xml:"name"`       // Thẻ <name>
	Age        int    `xml:"age"`        // Thẻ <age>
	Gender     string `xml:"gender"`     // Thẻ <gender>
	Occupation string `xml:"occupation"` // Thẻ <occupation>
}
