package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config"
	ginResident "main.go/modules/resident/transport/gin"
)

func main() {
	cfg := config.GetConfig()
	r := gin.Default()

	resident := r.Group("/resident")
	{
		resident.GET("/total", ginResident.TotalResidents(cfg))
		resident.GET("/total_sequential ", ginResident.TotalResidentsSequential(cfg))
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//worker.InitEmailWorker(ctx, sender, 100)
	//go func() {
	//	sig := make(chan os.Signal, 1)
	//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//	<-sig
	//	cancel()
	//}()
	r.Run(":3000")

	//	countries := []struct {
	//		Name       string
	//		Folder     string
	//		Prefix     string
	//		RegionList []string
	//	}{
	//		{"Vietnam", "vietnam", "VN", []string{"Hanoi", "Saigon", "DaNang"}},
	//		{"Thailand", "thailan", "TH", []string{"Bangkok", "Phuket", "ChiangMai"}},
	//		{"Cambodia", "cambodia", "KH", []string{"PhnomPenh", "SiemReap", "Sihanoukville"}},
	//	}
	//
	//	totalCitizens := 10000
	//	perCountry := totalCitizens / len(countries)
	//
	//	globalCounter := 1
	//
	//	for _, c := range countries {
	//		fmt.Printf(" đang xử lý trạm: %s...\n", c.Name)
	//
	//		// 1. Tạo thư mục provider/country
	//		folderPath := fmt.Sprintf("provider/%s", c.Folder)
	//		os.MkdirAll(folderPath, 0755)
	//
	//		// 2. Hàm sinh dữ liệu và nhóm theo Region ngay lập tức
	//		regionsData := make(map[string]*model.Region)
	//		for _, rName := range c.RegionList {
	//			regID := fmt.Sprintf("%s-%s", c.Prefix, strings.ToUpper(rName[:2]))
	//			regionsData[regID] = &model.Region{ID: regID, Name: rName}
	//		}
	//
	//		// Sinh 333+ người cho mỗi nước
	//		for i := 0; i < perCountry; i++ {
	//			rName := c.RegionList[i%len(c.RegionList)]
	//			regID := fmt.Sprintf("%s-%s", c.Prefix, strings.ToUpper(rName[:2]))
	//
	//			citizen := model.Citizen{
	//				ID:         fmt.Sprintf("%s-%d", c.Prefix, globalCounter),
	//				Name:       fmt.Sprintf("Person %d", globalCounter),
	//				Age:        20 + (globalCounter % 40),
	//				Gender:     []string{"Male", "Female", "Other"}[globalCounter%3],
	//				Occupation: fmt.Sprintf("Job-%d", globalCounter),
	//			}
	//			regionsData[regID].Citizens = append(regionsData[regID].Citizens, citizen)
	//			globalCounter++
	//		}
	//
	//		// 3. Chuyển cấu trúc dữ liệu thành XML string
	//		xmlContent := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<census_data country=\"%s\">\n", c.Name)
	//		for _, reg := range regionsData {
	//			xmlContent += fmt.Sprintf("    <region id=\"%s\" name=\"%s\">\n", reg.ID, reg.Name)
	//			for _, cit := range reg.Citizens {
	//				xmlContent += fmt.Sprintf(`        <citizen id="%s">
	//            <name>%s</name>
	//            <age>%d</age>
	//            <gender>%s</gender>
	//            <occupation>%s</occupation>
	//        </citizen>
	//`, cit.ID, cit.Name, cit.Age, cit.Gender, cit.Occupation)
	//			}
	//			xmlContent += "    </region>\n"
	//		}
	//		xmlContent += "</census_data>"
	//
	//		// 4. Ghi đồng thời vào Resident và Resident_Replica
	//		targetFiles := []string{"resident.xml", "resident_replica.xml"}
	//		for _, fName := range targetFiles {
	//			fullPath := fmt.Sprintf("%s/%s", folderPath, fName)
	//			err := os.WriteFile(fullPath, []byte(xmlContent), 0644)
	//			if err != nil {
	//				fmt.Printf("❌ Lỗi ghi %s: %v\n", fullPath, err)
	//			}
	//		}
	//	}
	//
	//	fmt.Println("\n✅ HOÀN THÀNH: Đã tạo 1000 người chia đều vào các trạm Primary & Replica!")
}
