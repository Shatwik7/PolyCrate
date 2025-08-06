# 🧩 Polycrate – 3D Asset Sharing Community Platform

Polycrate is a scalable, microservice-based backend system built in Go, designed to manage, convert, and deliver 3D assets for web-based visualization platforms. It tightly integrates with a React.js frontend using Three.js and GLTFLoader for interactive, compressed 3D previews.

---

## 🧱 Features

- 👤 **User authentication and account management**
- 📁 **Upload 3D assets:** `.obj`, `.gltf`, `.glb`, or `.flx`
- 🔄 **Automatic format conversion and preview generation**
- 📉 **GLB preview optimization and GZIP compression** for fast delivery
- ☁️ **S3-compatible object storage**
- 🔍 **Search** (planned: Elasticsearch)
- 📨 **Notification queue** (planned: NATS or RabbitMQ)
- 💬 **gRPC microservices** with protobuf interfaces
- 🐳 **Docker-based development infrastructure**
- 🧪 **Full test coverage** with reporting

---

## 🧰 Tech Stack

| Layer            | Technology                                      |
|------------------|-------------------------------------------------|
| Backend Services | Go (gRPC, PostgreSQL, Redis)                    |
| Asset Worker     | Go / Node.js / Python (convert & compress)      |
| Frontend         | React.js + Three.js (`@react-three/fiber`)      |
| Database         | PostgreSQL                                      |
| Object Storage   | S3 (AWS or MinIO)                               |
| Queue            | Redis / NATS / RabbitMQ                         |
| Infrastructure   | Docker, Docker Compose                          |
| File Conversion  | Blender CLI, gltf-pipeline, gzip                |
| Testing          | `go test`, Testify                              |

---
## 🚀 Getting Started

### 1️⃣ Start Local Dev Infrastructure

This launches Postgres, Redis (optional), and mounts `schema.sql` into the container.

task dev_infra

text

> **Note:**  
> Make sure `schema.sql` exists under `infra/dev/schema.sql` if you're using Docker volume mounting.

---

### 2️⃣ Run the User Service

task run:user_service


Or directly:
```bash
go run ./cmd/user_service/main.go
```


---

### 3️⃣ Generate Protobuf Code

This will generate Go code from your `.proto` files for gRPC services:

task gen:protobuf


Make sure you have the necessary tools installed:

```bash

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
---

### 4️⃣ Run Tests and Generate Coverage

task test

text

This generates:
- `coverage.out` — raw test coverage file
- `coverage.html` — annotated HTML coverage report

## 🔗 Useful Commands

- **Start All Services**:  
  `docker-compose up -d`

- **Clean Containers/Volumes**:  
  `docker-compose down -v`

---