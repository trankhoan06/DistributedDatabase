package model

import "encoding/xml"

type Result struct {
	Total   int
	Country string
}
type NodeResult struct {
	Count int
	Err   error
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

// FragmentAnalysis lưu trữ thông tin phân tích tại mỗi trạm
type FragmentAnalysis struct {
	Country      string `json:"country"`
	SubQuery     string `json:"sub_query"`
	Result       int    `json:"result"`
	Overhead     int    `json:"overhead_bytes"`      // Kích thước của kết quả gửi về
	OriginalSize int64  `json:"original_size_bytes"` // Kích thước của toàn bộ file XML
}

// AnalysisResult lưu trữ kết quả phân tích toàn cục
type AnalysisResult struct {
	GlobalQuery      string             `json:"global_query"`
	GlobalResult     int                `json:"global_result"`
	Decomposition    []FragmentAnalysis `json:"decomposition"`
	TotalOverhead    int                `json:"total_communication_overhead_bytes"`
	TotalOriginal    int64              `json:"total_overhead_without_decomposition_bytes"`
	SavedBandwidth   int64              `json:"saved_bandwidth_bytes"`
	XQueryComplexity string             `json:"xquery_complexity"`
}
