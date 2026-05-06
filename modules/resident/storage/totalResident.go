package storage

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"os"
)

func (sql *ResidentFile) QueryCountry(countryPath string) (int, error) {
	//path := common.PathFile(countryPath)
	//if strings.Contains(countryPath, "provider/vietnam/resident.xml") {
	//	fmt.Println("[DEBUG] Giả lập trạm Vietnam bị treo cứng...")
	//	time.Sleep(10 * time.Second)
	//}
	data, err := os.Open(countryPath)
	if err != nil {
		return 0, fmt.Errorf("node %s is down", countryPath)
	}
	defer data.Close()

	// Parse XML
	doc, _ := xmlquery.Parse(data)

	// Thực thi XQuery lấy tổng số người dân
	//countNode := xmlquery.Find(doc, "count(//citizen)")

	// Trả về số lượng (logic xử lý tùy thuộc vào thư viện XQuery bạn dùng)
	return len(xmlquery.Find(doc, "//citizen")), nil
	//return len(countNode)

}
