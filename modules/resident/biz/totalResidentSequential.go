package biz

import (
	"context"
	"fmt"
	"main.go/common"
	"time"
)

func (biz *ResidentCommon) NewTotalResidentSequential() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		select {
		case <-ctx.Done():
			return totalCount, fmt.Errorf("he thong bi treo hoac qua thoi gian xu ly: %v", ctx.Err())
		default:
			count, err := biz.Resident.QueryCountry(common.PathFile(c.primary))
			if err != nil {
				fmt.Printf("Node %s chinh bi loi, dang thu Replica...\n", c.name)
				count, err = biz.Resident.QueryCountry(common.PathFileReplica(c.primary))
				if err != nil {
					return totalCount, fmt.Errorf("ca file chinh va replica cua %s deu loi: %v", c.name, err)
				}
			}
			totalCount += count
		}
	}

	return totalCount, nil
}
