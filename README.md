Distributed Population Census System (Go & XML)
🚀 Đề tài #76: Query Decomposition & Global Aggregation
Hệ thống mô phỏng một cơ sở dữ liệu dân cư phân tán được chia mảnh theo quốc gia (Vietnam, Thailand, Cambodia). Hệ thống sử dụng ngôn ngữ Go để tối ưu hóa việc truy vấn song song và xử lý lỗi cục bộ (Fault Tolerance).
🛠 Core Features
Horizontal Fragmentation: Dữ liệu XML được phân mảnh ngang dựa trên thuộc tính quốc gia.
Parallel Query Execution: Sử dụng Goroutines để thực thi truy vấn tại tất cả các trạm cùng lúc, giảm thời gian phản hồi từ O(N) xuống O(1) (theo số lượng trạm).
Smart Failover Mechanism:Tự động phát hiện trạm bị sập (Crash) hoặc treo (Freeze) bằng context.WithTimeout.Tự động chuyển hướng truy vấn sang Replica Node nếu trạm chính không phản hồi sau 3 giây.Query Decomposition: Phân rã truy vấn tổng thành các câu lệnh XQuery (count, sum) thực thi trực tiếp tại trạm để giảm thiểu Communication Overhead.📂 Project StructureBash.
├── common/             # Các tiện ích cấu hình và xử lý đường dẫn file
├── modules/
│   └── resident/
│       ├── biz/        # Tầng nghiệp vụ: Điều phối, Parallel & Failover logic
│       ├── model/      # Định nghĩa cấu trúc dữ liệu (structs)
│       └── storage/    # Tầng dữ liệu: Thực thi XQuery trên file XML vật lý
├── provider/           # Giả lập các Nodes lưu trữ dữ liệu
│   ├── vietnam/        # resident.xml & resident_replica.xml
│   ├── thailan/        # resident.xml & resident_replica.xml
│   └── cambodia/       # resident.xml & resident_replica.xml
├── config.json         # Cấu hình hệ thống và đường dẫn trạm
└── main.go             # Điểm khởi chạy ứng dụng
⚙️ Installation & Setup1. Yêu cầu hệ thốngGo version 1.18 trở lên.Thư viện hỗ trợ XQuery:Bashgo get github.com/antchfx/xmlquery
2. Chạy ứng dụngBash# Clone dự án
   git clone https://github.com/username/project-id-76.git

# Di chuyển vào thư mục
cd project-id-76

# Khởi chạy hệ thống
go run main.go
🧪 Simulation Scenarios (Demo Guide)
Kịch bản 1: Truy vấn song song bình thường
Hệ thống gọi đồng thời 3 trạm. Kết quả sẽ được tổng hợp ngay khi trạm chậm nhất hoàn thành. Thời gian phản hồi sẽ xấp xỉ thời gian của trạm có dữ liệu lớn nhất.Kịch bản 2: Xử lý lỗi "Treo trạm" (Node Freeze)Mở file storage/resident_file.go.Bật đoạn mã giả lập time.Sleep(10 * time.Second) cho trạm Vietnam.Kết quả: Hệ thống sẽ đợi 3 giây, sau đó in ra log: [Timeout] Trạm Vietnam bị treo, đang gọi dự phòng... và trả về kết quả từ file Replica thành công.📊 Metrics AnalysisCommunication Overhead: Cực thấp. Thay vì gửi toàn bộ file XML qua mạng, mỗi trạm chỉ gửi về 1 số nguyên (kết quả của hàm count).Complexity: Độ phức tạp tập trung ở tầng Biz (quản lý Goroutine & Channel), giúp tầng Storage và Client trở nên đơn giản và nhẹ nhàng.Availability: Đạt mức sẵn sàng cao nhờ cơ chế dự phòng 1-1 cho mỗi mảnh dữ liệu.   