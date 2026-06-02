package biz

import (
	"fmt"
	"main.go/common"
	"main.go/modules/resident/model"
	"os"
)

func (biz *ResidentCommon) AnalyzeCensusQuery() (*model.AnalysisResult, error) {
	configs := []struct {
		primary string
		name    string
	}{
		{biz.cfg.VietNamXml, "Vietnam"},
		{biz.cfg.ThaiLanXml, "Thailand"},
		{biz.cfg.CambodiaXml, "Cambodia"},
	}

	analysisResult := &model.AnalysisResult{
		GlobalQuery:      "count(//citizen)",
		Decomposition:    []model.FragmentAnalysis{},
		XQueryComplexity: "O(N) at local nodes for traversal, O(1) bandwidth cost per node. Global aggregation is O(K) where K is number of nodes.",
	}

	for _, c := range configs {
		path := common.PathFile(c.primary)
		
		count, err := biz.Resident.QueryCountry(path)
		if err != nil {
			path = common.PathFileReplica(c.primary)
			count, err = biz.Resident.QueryCountry(path)
			if err != nil {
				return nil, fmt.Errorf("lỗi đọc dữ liệu trạm %s", c.name)
			}
		}

		resultString := fmt.Sprintf("%d", count)
		overhead := len([]byte(resultString))

		var originalSize int64 = 0
		fileInfo, err := os.Stat(path)
		if err == nil {
			originalSize = fileInfo.Size()
		}

		frag := model.FragmentAnalysis{
			Country:      c.name,
			SubQuery:     "count(//citizen)",
			Result:       count,
			Overhead:     overhead,
			OriginalSize: originalSize,
		}

		analysisResult.Decomposition = append(analysisResult.Decomposition, frag)
		analysisResult.GlobalResult += count
		analysisResult.TotalOverhead += overhead
		analysisResult.TotalOriginal += originalSize
	}

	analysisResult.SavedBandwidth = analysisResult.TotalOriginal - int64(analysisResult.TotalOverhead)

	return analysisResult, nil
}
