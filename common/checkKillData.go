package common

import (
	"fmt"
	"os"
)

func CheckKillData(primaryPath string, replicaPath string) (string, error) {
	_, err := os.ReadFile(primaryPath)
	if err == nil {
		fmt.Printf("--- Đang sử dụng dữ liệu chính: %s ---\n", primaryPath)
		return primaryPath, nil
	}

	fmt.Printf("!!! Cảnh báo: Không tìm thấy %s. Đang chuyển sang replica ... !!!\n", primaryPath)
	_, err1 := os.ReadFile(replicaPath)
	if err1 == nil {
		fmt.Printf("--- Đang sử dụng dữ liệu dự phòng: %s ---\n", replicaPath)
		return replicaPath, nil
	}

	return "", fmt.Errorf("lỗi: cả file chính và dự phòng đều không tồn tại")
}
