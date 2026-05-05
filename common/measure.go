package common

import "fmt"

func ExecuteAndMeasure(path string, query string) {
	// 1. Giả lập thực thi XQuery và lấy kết quả
	result := "300" // Giả sử kết quả trả về từ node

	// 2. Đo kích thước bytes
	overhead := len([]byte(result))

	fmt.Printf("Query: %s\n", query)
	fmt.Printf("Overhead: %d bytes\n", overhead)
}
