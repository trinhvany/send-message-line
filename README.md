# 📬 Message System: Publisher & Subscriber (Golang + SQS + S3 + Redis + RDS)

## 🧩 Mục tiêu

Hệ thống này gồm 2 thành phần:

1. **Publisher**: 
   - Lấy dữ liệu message từ RDS
   - Lấy danh sách userid từ S3
   - Build và push message lên SQS

2. **Subscriber**:
   - Đọc message từ SQS
   - Sử dụng Redis để kiểm tra trùng lặp user
   - Gửi tin nhắn tới người dùng (qua LINE API)

---

## 📦 Cấu trúc thư mục

```bash
.
├── cmd/
│   ├── main.go
│
├── internal/
│   ├── rds/
│   │   ├── client.go
│   │   ├── message.go
│   │   └── message_detail.go
│   ├── redis/
│   │   └── client.go
│   ├── s3/
│   │   ├── client.go
│   │   └── data.go
│   ├── sqs/
│   │   └── client.go
│
│   ├── model/
│   │   ├── message.go
│   │   └── message_detail.go
│
│   ├── publisher/
│   │   └── publisher.go
│
│   ├── subscriber/
│   │   └── subscriber.go
│
│   ├── system/
│   │   ├── config.go
│   │   └── const.go
│
│   └── utils/
│       └── line.go
│
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Run dự án
1. **Run docker**

```base
docker compose up -d
```

2. **Tạo RDS**
(image.png)

3. **Cài gói cần thiết**
```base
go mod tidy
```

3. **Run source**
```base
go run cmd/main.go
```
