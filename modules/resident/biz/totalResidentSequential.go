package biz

import (
	"fmt"
	"main.go/common"
	"time"
)

func (biz *ResidentCommon) NewTotalResidentSequential() (int, error) {
	var totalCount int
	configs := []struct {
		primary string
		name    string
	}{
		{biz.cfg.VietNamXml, "Vietnam"},
		{biz.cfg.ThaiLanXml, "Thailand"},
		{biz.cfg.CambodiaXml, "Cambodia"},
	}

	for _, c := range configs {
		// Thử trạm chính với giới hạn 3 giây
		count, err := biz.queryWithTimeout(common.PathFile(c.primary), 3*time.Second)

		if err != nil {
			// Nếu LỖI hoặc TREO -> Thử Replica ngay
			fmt.Printf("Trạm %s gặp sự cố (%v), đang thử Replica...\n", c.name, err)

			// Đọc Replica (có thể không cần timeout cho replica hoặc để ngắn hơn)
			count, err = biz.Resident.QueryCountry(common.PathFileReplica(c.primary))
			if err != nil {
				return totalCount, fmt.Errorf("cả chính và phụ của %s đều hỏng", c.name)
			}
		}
		totalCount += count
	}
	return totalCount, nil
}

// Đây là "người giám sát"
func (biz *ResidentCommon) queryWithTimeout(path string, timeout time.Duration) (int, error) {
	ch := make(chan int, 1)
	errCh := make(chan error, 1)

	go func() {
		res, err := biz.Resident.QueryCountry(path)
		if err != nil {
			errCh <- err
			return
		}
		ch <- res
	}()

	select {
	case res := <-ch:
		return res, nil
	case err := <-errCh:
		return 0, err
	case <-time.After(timeout):
		return 0, fmt.Errorf("treo") // Trả lỗi để hàm chính biết mà nhảy sang Replica
	}
}
