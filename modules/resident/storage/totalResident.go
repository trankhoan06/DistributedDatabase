package storage

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"os"
)

func QueryCountry(countryPath string) int {
	f, err := os.Open(countryPath)
	if err != nil {
		fmt.Println("Lỗi mở file:", err)
		return 0
	}
	defer f.Close()

	// Parse XML
	doc, _ := xmlquery.Parse(f)

	// Thực thi XQuery lấy tổng số người dân
	//countNode := xmlquery.Find(doc, "count(//citizen)")

	// Trả về số lượng (logic xử lý tùy thuộc vào thư viện XQuery bạn dùng)
	return len(xmlquery.Find(doc, "//citizen"))
	//return len(countNode)

}
