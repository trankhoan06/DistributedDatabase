package biz

//
//import (
//	"fmt"
//	"github.com/antchfx/xmlquery"
//	"os"
//	"strings"
//)
//
//func TransactionImport(country string, newRecords []*xmlquery.Node) error {
//	dirPath := fmt.Sprintf("provider/%s", strings.ToLower(country))
//	os.MkdirAll(dirPath, os.ModePerm)
//	finalPath := dirPath + "/resident.xml"
//	tempPath := finalPath + ".tmp"
//
//	// 1. Khởi tạo/Đọc dữ liệu hiện có của trạm
//	var doc *xmlquery.Node
//	if _, err := os.Stat(finalPath); err == nil {
//		f, _ := os.Open(finalPath)
//		doc, _ = xmlquery.Parse(f)
//		f.Close()
//	} else {
//		// Tạo mới nếu trạm chưa có file
//		doc, _ = xmlquery.Parse(strings.NewReader(fmt.Sprintf("<census_data country=\"%s\"></census_data>", country)))
//	}
//
//	// 2. Duyệt qua từng bản ghi mới để phân loại vào Region
//	root := xmlquery.FindOne(doc, "//census_data")
//	for _, rec := range newRecords {
//		regID := rec.SelectAttr("region_id")
//		regName := rec.SelectAttr("region_name")
//
//		// Kiểm tra xem Region đã tồn tại trong trạm chưa
//		regionNode := xmlquery.FindOne(doc, fmt.Sprintf("//region[@id='%s']", regID))
//		if regionNode == nil {
//			// Nếu chưa có Region -> Tạo mới
//			regionNode = &xmlquery.Node{
//				Type: xmlquery.ElementNode,
//				Data: "region",
//				Attr: []xmlquery.Attr{
//					{Name: xmlquery.Name{Local: "id"}, Value: regID},
//					{Name: xmlquery.Name{Local: "name"}, Value: regName},
//				},
//			}
//			xmlquery.AddChild(root, regionNode)
//		}
//
//		// Tạo Node Citizen từ bản ghi record
//		citizen := &xmlquery.Node{Type: xmlquery.ElementNode, Data: "citizen", Attr: []xmlquery.Attr{{Name: xmlquery.Name{Local: "id"}, Value: xmlquery.FindOne(rec, "id").InnerText()}}}
//		xmlquery.AddChild(citizen, &xmlquery.Node{Type: xmlquery.ElementNode, Data: "name", FirstChild: &xmlquery.Node{Type: xmlquery.TextNode, Data: xmlquery.FindOne(rec, "name").InnerText()}})
//		xmlquery.AddChild(citizen, &xmlquery.Node{Type: xmlquery.ElementNode, Data: "age", FirstChild: &xmlquery.Node{Type: xmlquery.TextNode, Data: xmlquery.FindOne(rec, "age").InnerText()}})
//		xmlquery.AddChild(citizen, &xmlquery.Node{Type: xmlquery.ElementNode, Data: "gender", FirstChild: &xmlquery.Node{Type: xmlquery.TextNode, Data: xmlquery.FindOne(rec, "gender").InnerText()}})
//		xmlquery.AddChild(citizen, &xmlquery.Node{Type: xmlquery.ElementNode, Data: "occupation", FirstChild: &xmlquery.Node{Type: xmlquery.TextNode, Data: xmlquery.FindOne(rec, "occupation").InnerText()}})
//
//		// Thêm Citizen vào Region tương ứng
//		xmlquery.AddChild(regionNode, citizen)
//	}
//
//	// 3. TRANSACTION Ghi file: Ghi ra file tạm trước
//	err := os.WriteFile(tempPath, []byte(doc.OutputXML(true)), 0644)
//	if err != nil {
//		return err
//	}
//
//	// COMMIT: Đổi tên file tạm thành file chính (Thao tác Atomic)
//	err = os.Rename(tempPath, finalPath)
//	if err != nil {
//		os.Remove(tempPath)
//		return err
//	}
//
//	fmt.Printf("✔ Hoàn tất nạp dữ liệu cho trạm: %s\n", country)
//	return nil
//}
