package biz

import (
	"fmt"
	"main.go/common"
	"time"
)

func (biz *ResidentCommon) NewTotalResident() (int, error) {
	results := make(chan int, 3)
	countries := []string{biz.cfg.VietNamXml, biz.cfg.ThaiLanXml, biz.cfg.CambodiaXml}

	for _, country := range countries {
		go func(c string) {
			// Tạo channel riêng cho từng node để kiểm soát việc "đứng"
			nodeChan := make(chan int, 1)

			// Chạy việc truy vấn trong một goroutine con khác
			go func() {
				path := common.PathFile(c)
				res, err := biz.Resident.QueryCountry(path)
				if err == nil {
					nodeChan <- res
				}
			}()

			// KIỂM SOÁT VIỆC "ĐỨNG" CỦA TỪNG NODE
			select {
			case res := <-nodeChan:
				// Node chính chạy nhanh, về kịp
				results <- res
			case <-time.After(time.Second * 3):
				// SAU 3 GIÂY MÀ NODE CHÍNH CHƯA XONG -> COI NHƯ BỊ ĐỨNG
				fmt.Printf("[Timeout] Trạm %s bị treo, chuyển ngay sang Replica...\n", c)

				pathRep := common.PathFileReplica(c)
				resRep, errRep := biz.Resident.QueryCountry(pathRep)

				if errRep == nil {
					results <- resRep
				} else {
					results <- 0 // Cả 2 đều hỏng
				}
			}
		}(country)
	}

	// Tầng tổng hợp (Aggregation) giữ nguyên
	total := 0
	for i := 0; i < len(countries); i++ {
		total += <-results
	}
	return total, nil
}
