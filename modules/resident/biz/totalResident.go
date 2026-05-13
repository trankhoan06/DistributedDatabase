package biz

import (
	"fmt"
	"main.go/common"
	"main.go/modules/resident/model"
	"time"
)

func (biz *ResidentCommon) NewTotalResident() (int, error) {
	results := make(chan model.NodeResult, 3)
	countries := []string{biz.cfg.VietNamXml, biz.cfg.ThaiLanXml, biz.cfg.CambodiaXml}

	for _, country := range countries {
		go func(c string) {
			nodeChan := make(chan model.NodeResult, 1)

			go func() {
				path := common.PathFile(c)
				res, err := biz.Resident.QueryCountry(path)
				nodeChan <- model.NodeResult{Count: res, Err: err}
			}()

			select {
			case primary := <-nodeChan:
				if primary.Err == nil {
					results <- primary
				} else {
					fmt.Printf("[Error] Trạm %s lỗi: %v. Thử Replica...\n", c, primary.Err)
					results <- biz.tryReplica(c)
				}

			case <-time.After(time.Second * 3):
				fmt.Printf("[Timeout] Trạm %s treo sau 3s. Thử Replica...\n", c)
				results <- biz.tryReplica(c)
			}
		}(country)
	}

	total := 0
	for i := 0; i < len(countries); i++ {
		res := <-results
		if res.Err != nil {
			return 0, res.Err
		}
		total += res.Count
	}
	return total, nil
}

func (biz *ResidentCommon) tryReplica(c string) model.NodeResult {
	pathRep := common.PathFileReplica(c)
	resRep, errRep := biz.Resident.QueryCountry(pathRep)
	if errRep != nil {
		return model.NodeResult{Count: 0, Err: fmt.Errorf("cả Primary và Replica của %s đều thất bại", c)}
	}
	return model.NodeResult{Count: resRep, Err: nil}
}
