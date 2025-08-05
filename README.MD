# 🧩 Polycrate – 3D Asset Sharing Community Platform

Polycrate is a scalable **microservice-based backend system** built in **Go**, designed to manage and deliver **3D assets** for web-based visualization platforms. It integrates tightly with a **React.js** frontend that renders models using **Three.js**.

---

## 🧱 Features

- 👤 User management with authentication and password hashing
- 📦 Asset storage metadata, tagging, and likes
- 🔍 Search (planned via Elasticsearch)
- 📬 Notification queue (planned)
- 🧪 Full test coverage with reporting
- 💬 gRPC microservices with protobuf interfaces
- 🐳 Docker-based development infrastructure

---

## 🧰 Tech Stack

| Layer          | Technology                       |
|----------------|----------------------------------|
| Backend        | Go (gRPC, PostgreSQL)            |
| Frontend       | React.js + Three.js              |
| Database       | PostgreSQL                       |
| Infrastructure | Docker, Docker Compose           |
| Search         | Elasticsearch (planned)          |
| Messaging      | RabbitMQ/NATS (planned)          |
| Caching        | Redis (planned)                  |
| Testing        | `go test`, `Testify`             |

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