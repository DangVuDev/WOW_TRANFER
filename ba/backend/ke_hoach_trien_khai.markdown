# Kế Hoạch Triển Khai WoWToken Web API

Tài liệu này trình bày kế hoạch chi tiết để triển khai WoWToken Web API, một API RESTful cho ứng dụng ngân hàng phi tập trung với mô hình ví custodial, theo tài liệu đã cung cấp. Kế hoạch bao gồm các giai đoạn phát triển, kiểm thử, triển khai và bảo trì, tập trung vào tính mô-đun, bảo mật và khả năng mở rộng. Dự án sử dụng Go với các thư viện như `gorilla/mux`, `go-ethereum`, `mongo-driver`, `go-redis` và `jwt-go`.

---

## 1. Tổng Quan Dự Án
- **Mục tiêu**: Xây dựng API RESTful cho hợp đồng thông minh `WoWToken` và quản lý người dùng, cho phép người dùng không chuyên về kỹ thuật tương tác với hệ thống ngân hàng phi tập trung thông qua mô hình ví custodial.
- **Phạm vi**:
  - Triển khai các endpoint liên quan đến token (ví dụ: `/balances/{address}`, `/transfer`).
  - Triển khai các endpoint sự kiện đa chữ ký (ví dụ: `/events/mint`, `/events/sign`).
  - Triển khai các endpoint quản lý người dùng (ví dụ: `/users/register`, `/users/login`).
  - Đảm bảo tương tác blockchain an toàn, xác thực JWT và mã hóa khóa riêng.
  - Sử dụng MongoDB cho dữ liệu người dùng, Redis cho bộ nhớ đệm và Infura/Alchemy để kết nối Ethereum.
- **Đội ngũ**: Giả định đội gồm 2-3 lập trình viên (backend, blockchain, DevOps) và một quản lý dự án.
- **Thời gian**: 8 tuần cho phát triển, kiểm thử và triển khai ban đầu (bắt đầu từ 1/7/2025).

---

## 2. Các Giai Đoạn Triển Khai và Thời Gian

### Giai đoạn 1: Thiết lập và Lập kế hoạch (Tuần 1, 1/7 - 7/7/2025)
**Mục tiêu**: Thiết lập nền tảng dự án, công cụ và môi trường.
- **Nhiệm vụ**:
  1. **Thiết lập kho mã nguồn**:
     - Tạo kho Git (ví dụ: GitHub).
     - Khởi tạo Go module (`go mod init wowtoken-api`).
     - Thêm các phụ thuộc (`gorilla/mux`, `go-ethereum`, `mongo-driver`, `go-redis`, `jwt-go`, `viper`).
  2. **Xác định cấu trúc thư mục**:
     - Tạo cấu trúc thư mục (`cmd`, `internal`, `pkg`, `scripts`, `tests`) như đã đề xuất.
     - Tạo các tệp ban đầu (ví dụ: `main.go`, `config.go`, các tệp controller/service trống).
  3. **Cấu hình môi trường**:
     - Thiết lập `config.yaml` với các giá trị placeholder cho `PORT`, `INFURA_URL`, `MONGODB_URI`, `REDIS_ADDR`, `JWT_SECRET`, `ENCRYPTION_KEY`.
     - Cấu hình biến môi trường cho phát triển cục bộ.
  4. **Thiết lập hạ tầng**:
     - Chuẩn bị các phiên bản MongoDB và Redis (cục bộ hoặc đám mây, ví dụ: MongoDB Atlas, Redis Labs).
     - Lấy khóa API Infura/Alchemy để truy cập Ethereum mainnet/testnet.
     - Thiết lập Docker cho phát triển và triển khai cục bộ.
  5. **Xác định tương tác hợp đồng thông minh**:
     - Lấy ABI và địa chỉ của hợp đồng `WoWToken`.
     - Lập kế hoạch gọi các phương thức hợp đồng (ví dụ: `transferToken`, `mintToken`, `balances`).
- **Phụ thuộc**:
  - Git, Go 1.21, Docker, tài khoản Infura/Alchemy, MongoDB, Redis.
- **Kết quả**:
  - Kho Git với cấu trúc ban đầu.
  - Tệp `config.yaml` và thiết lập môi trường.
  - Cấu hình Docker (`Dockerfile`, `docker-compose.yml`).
- **Cột mốc**: Khung dự án sẵn sàng để phát triển.

### Giai đoạn 2: Phát triển Cốt lõi (Tuần 2-5, 8/7 - 4/8/2025)
**Mục tiêu**: Triển khai chức năng API cốt lõi cho token, quản lý người dùng và sự kiện đa chữ ký.
- **Nhiệm vụ**:
  1. **Tuần 2: Cấu hình và Cơ sở dữ liệu (8/7 - 14/7)**:
     - Triển khai `internal/config/config.go` để tải cấu hình bằng `viper`.
     - Triển khai `pkg/db/mongodb.go` cho kết nối MongoDB và CRUD cơ bản (ví dụ: `SaveUser`, `GetUserByEmail`).
     - Triển khai `pkg/db/redis.go` cho kết nối Redis và thiết lập bộ nhớ đệm.
     - Kiểm tra kết nối cơ sở dữ liệu và tải cấu hình.
  2. **Tuần 3: Helper và Middleware (15/7 - 21/7)**:
     - Triển khai `internal/helpers/jwt_helper.go` để tạo và xác thực JWT.
     - Triển khai `internal/helpers/crypto_helper.go` để mã hóa/giải mã khóa riêng bằng AES-256.
     - Triển khai `internal/helpers/eth_helper.go` để xác thực địa chỉ Ethereum.
     - Triển khai `internal/helpers/response_helper.go` để định dạng phản hồi JSON/lỗi.
     - Triển khai `internal/middleware/auth_middleware.go` để xác thực JWT.
     - Triển khai `internal/middleware/rate_limit.go` để giới hạn tỷ lệ bằng Redis.
  3. **Tuần 4: Dịch vụ Blockchain và Token (22/7 - 28/7)**:
     - Triển khai `internal/services/blockchain_service.go` để kết nối Ethereum qua `go-ethereum` và gọi các phương thức hợp đồng `WoWToken` (ví dụ: `GetBalance`, `TransferToken`).
     - Triển khai `internal/services/token_service.go` cho các hoạt động token (ví dụ: `/balances`, `/transfer`, `/approve`, `/burn`).
     - Triển khai `internal/controllers/token_controller.go` cho các endpoint liên quan đến token.
     - Xác định `internal/models/token.go` cho các cấu trúc dữ liệu token.
  4. **Tuần 5: Dịch vụ Người dùng và Sự kiện (29/7 - 4/8)**:
     - Triển khai `internal/services/user_service.go` cho đăng ký, đăng nhập và quản lý ví custodial.
     - Triển khai `internal/services/event_service.go` cho các sự kiện đa chữ ký (ví dụ: `/events/mint`, `/events/sign`).
     - Triển khai `internal/controllers/user_controller.go` và `event_controller.go` cho các endpoint tương ứng.
     - Xác định `internal/models/user.go` và `event.go` cho các cấu trúc dữ liệu.
- **Phụ thuộc**:
  - ABI và địa chỉ hợp đồng `WoWToken`.
  - Các phiên bản MongoDB và Redis.
  - Khóa API Infura/Alchemy.
- **Kết quả**:
  - Các module cấu hình, cơ sở dữ liệu và helper hoạt động.
  - Các endpoint liên quan đến token (`/balances`, `/transfer`, v.v.).
  - Các endpoint quản lý người dùng (`/users/register`, `/users/login`, v.v.).
  - Các endpoint sự kiện đa chữ ký (`/events/mint`, `/events/sign`, v.v.).
- **Cột mốc**: Chức năng API cốt lõi được triển khai và có thể kiểm thử cục bộ.

### Giai đoạn 3: Kiểm thử và Xác nhận (Tuần 6-7, 5/8 - 18/8/2025)
**Mục tiêu**: Đảm bảo độ tin cậy, bảo mật và tuân thủ thông số kỹ thuật API.
- **Nhiệm vụ**:
  1. **Kiểm thử Đơn vị (Tuần 6, 5/8 - 11/8)**:
     - Viết kiểm thử đơn vị cho `controllers` (ví dụ: `token_controller_test.go` để kiểm tra `/balances`, `/transfer`).
     - Viết kiểm thử đơn vị cho `services` (ví dụ: giả lập blockchain trong `token_service_test.go`).
     - Viết kiểm thử đơn vị cho `helpers` (ví dụ: `jwt_helper_test.go` để kiểm tra xác thực JWT).
     - Đạt độ phủ kiểm thử ít nhất 80%.
  2. **Kiểm thử Tích hợp (Tuần 6-7, 5/8 - 18/8)**:
     - Kiểm tra luồng end-to-end (ví dụ: đăng ký người dùng, đăng nhập, chuyển token, tạo sự kiện mint).
     - Sử dụng testnet (ví dụ: Sepolia) để mô phỏng tương tác blockchain.
     - Xác minh các trường hợp lỗi (ví dụ: JWT không hợp lệ, số dư không đủ, hợp đồng bị tạm dừng).
  3. **Kiểm thử Bảo mật (Tuần 7, 12/8 - 18/8)**:
     - Kiểm tra xác thực JWT và giới hạn tỷ lệ.
     - Xác minh bảo mật mã hóa/giải mã khóa riêng.
     - Kiểm tra TLS và xác thực đầu vào (ví dụ: định dạng địa chỉ, số lượng).
     - Thực hiện kiểm thử thâm nhập cơ bản (ví dụ: SQL injection, XSS).
- **Phụ thuộc**:
  - Truy cập testnet (ví dụ: Sepolia qua Infura).
  - Khung kiểm thử (`testing`, `testify`).
- **Kết quả**:
  - Bộ kiểm thử đơn vị và tích hợp trong `tests/`.
  - Báo cáo độ phủ kiểm thử.
  - Báo cáo kiểm tra bảo mật cho xác thực và mã hóa.
- **Cột mốc**: API được kiểm thử và xác nhận đầy đủ theo yêu cầu.

### Giai đoạn 4: Triển khai (Tuần 8, 19/8 - 25/8/2025)
**Mục tiêu**: Triển khai API lên môi trường sản xuất.
- **Nhiệm vụ**:
  1. **Chuẩn bị Triển khai**:
     - Hoàn thiện `scripts/deploy.sh` để xây dựng và triển khai API.
     - Tạo hình ảnh Docker với cấu hình sản xuất.
     - Thiết lập hạ tầng đám mây (ví dụ: AWS EC2, ECS hoặc Kubernetes).
     - Cấu hình MongoDB, Redis và Infura/Alchemy cho sản xuất.
  2. **Triển khai API**:
     - Đẩy hình ảnh Docker lên registry (ví dụ: Docker Hub, AWS ECR).
     - Triển khai lên đám mây với bộ cân bằng tải và tự động mở rộng.
     - Thiết lập giám sát (ví dụ: AWS CloudWatch) và ghi log (ví dụ: Logrus).
  3. **Kiểm thử Sau Triển khai**:
     - Chạy kiểm thử smoke trên các endpoint sản xuất.
     - Xác minh kết nối blockchain và ký giao dịch.
     - Kiểm tra đăng ký người dùng và chuyển token trong sản xuất.
- **Phụ thuộc**:
  - Tài khoản nhà cung cấp đám mây (ví dụ: AWS).
  - Các phiên bản MongoDB/Redis sản xuất.
  - Tên miền cho `https://api.wowtoken.example.com`.
- **Kết quả**:
  - API triển khai tại `https://api.wowtoken.example.com/v1`.
  - Thiết lập giám sát và ghi log.
  - Báo cáo kiểm thử sau triển khai.
- **Cột mốc**: API hoạt động trong sản xuất.

### Giai đoạn 5: Bảo trì và Giám sát (Liên tục, Sau 25/8/2025)
**Mục tiêu**: Đảm bảo tính ổn định, hiệu suất và khả năng mở rộng của API.
- **Nhiệm vụ**:
  1. **Giám sát Hiệu suất**:
     - Theo dõi thời gian phản hồi API và tỷ lệ lỗi bằng CloudWatch.
     - Giám sát thời gian hoạt động của nhà cung cấp blockchain (Infura/Alchemy).
  2. **Xử lý Sự cố**:
     - Xử lý lỗi do người dùng báo cáo (ví dụ: qua `support@wowtoken.example.com`).
     - Vá các lỗ hổng bảo mật (ví dụ: cập nhật phụ thuộc).
  3. **Mở rộng Hạ tầng**:
     - Thêm bộ nhớ đệm cho các endpoint đọc nhiều (ví dụ: `/balances/{address}`) bằng Redis.
     - Mở rộng số lượng phiên bản API dựa trên lưu lượng (ví dụ: tự động mở rộng Kubernetes).
  4. **Cải tiến Tính năng**:
     - Thêm hỗ trợ xác thực hai yếu tố (2FA) cho các hành động nhạy cảm.
     - Triển khai các endpoint bổ sung (ví dụ: `/users/deposit`, `/users/withdraw`) nếu cần.
- **Phụ thuộc**:
  - Công cụ giám sát (ví dụ: CloudWatch, Prometheus).
  - Hệ thống ticket hỗ trợ.
- **Kết quả**:
  - Bảng điều khiển giám sát liên tục.
  - Bản sửa lỗi và cập nhật tính năng.

---

## 3. Quản lý Rủi ro
- **Rủi ro**: Nhà cung cấp blockchain bị gián đoạn (ví dụ: Infura).
  - **Giảm thiểu**: Sử dụng nhiều nhà cung cấp (Infura và Alchemy) với dự phòng.
- **Rủi ro**: Vi phạm bảo mật trong lưu trữ khóa riêng.
  - **Giảm thiểu**: Sử dụng AWS KMS để mã hóa khóa trong sản xuất; áp dụng kiểm soát truy cập nghiêm ngặt.
- **Rủi ro**: Lỗi hợp đồng thông minh hoặc trạng thái tạm dừng.
  - **Giảm thiểu**: Kiểm tra tương tác hợp đồng trên testnet; xác minh trạng thái `paused` trong dịch vụ.
- **Rủi ro**: Lưu lượng cao làm quá tải API.
  - **Giảm thiểu**: Triển khai giới hạn tỷ lệ và bộ nhớ đệm; sử dụng hạ tầng tự động mở rộng.

---

## 4. Yêu cầu Tài nguyên
- **Đội ngũ**:
  - 2 Lập trình viên Backend (Go, phát triển API).
  - 1 Lập trình viên Blockchain (Ethereum, `go-ethereum`).
  - 1 Kỹ sư DevOps (Docker, AWS, giám sát).
  - 1 Quản lý Dự án (điều phối, theo dõi tiến độ).
- **Công cụ**:
  - **Phát triển**: Go 1.21, VS Code, Git.
  - **Hạ tầng**: AWS (EC2/ECS, KMS), MongoDB Atlas, Redis Labs.
  - **Blockchain**: Infura/Alchemy, testnet Sepolia.
  - **Kiểm thử**: `testing`, `testify`, Postman để kiểm thử API.
- **Ngân sách**:
  - Chi phí đám mây: ~$100-$200/tháng cho AWS, MongoDB Atlas, Redis Labs.
  - Đăng ký Infura/Alchemy: ~$50/tháng cho truy cập testnet/mainnet.
  - Thời gian lập trình viên: ~$10,000-$15,000 cho 8 tuần (dựa trên 2-3 lập trình viên).

---

## 5. Tóm tắt Thời gian
| Giai đoạn             | Thời gian             | Kết quả Chính                        |
|-----------------------|-----------------------|--------------------------------------|
| Thiết lập và Lập kế hoạch | 1/7 - 7/7         | Kho mã, cấu hình, thiết lập Docker   |
| Phát triển Cốt lõi    | 8/7 - 4/8           | Endpoint API, dịch vụ, helper        |
| Kiểm thử và Xác nhận  | 5/8 - 18/8          | Bộ kiểm thử, báo cáo bảo mật         |
| Triển khai            | 19/8 - 25/8         | API sản xuất, thiết lập giám sát     |
| Bảo trì               | Sau 25/8            | Giám sát, sửa lỗi, cập nhật tính năng|

---

## 6. Tiêu chí Thành công
- Tất cả endpoint API (`/balances`, `/transfer`, `/users/register`, v.v.) hoạt động đúng theo tài liệu.
- Tương tác blockchain (ví dụ: `transferToken`, `mintToken`) an toàn và đáng tin cậy.
- Ví custodial của người dùng được quản lý an toàn với khóa riêng mã hóa.
- API đạt thời gian hoạt động 99.9% và xử lý 100 yêu cầu/giây với bộ nhớ đệm.
- Độ phủ kiểm thử ≥ 80% và không có lỗ hổng bảo mật nghiêm trọng.
- Triển khai thành công tại `https://api.wowtoken.example.com/v1` trước 25/8/2025.

---

## 7. Các Bước Tiếp Theo
- **Ngay lập tức (1/7/2025)**:
  - Thiết lập kho Git và đẩy cấu trúc thư mục ban đầu.
  - Chuẩn bị tài khoản MongoDB, Redis và Infura/Alchemy.
  - Phân công nhiệm vụ cho đội ngũ trong Giai đoạn 1.
- **Liên tục**:
  - Xem xét tiến độ hàng tuần để theo dõi phát triển so với thời gian.
  - Thực hiện đánh giá mã thường xuyên để đảm bảo chất lượng và tuân thủ thông số kỹ thuật.
  - Thiết lập tích hợp liên tục (CI) cho kiểm thử tự động.

Kế hoạch này đảm bảo một cách tiếp cận có cấu trúc để triển khai WoWToken Web API, cân bằng giữa tốc độ phát triển, chất lượng mã và sự sẵn sàng cho sản xuất.