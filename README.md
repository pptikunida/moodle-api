# 🎓 Moodle API Gateway - UNIDA Gontor

API ini adalah middleware antara Sistem Informasi Akademik (SIAKAD) dengan Moodle, ditulis menggunakan bahasa Go (Golang), yang mengimplementasikan clean architecture dan RESTful API. Digunakan untuk melakukan sinkronisasi user, pembuatan course, enrolment user, dan manajemen peran.

---

## 🚀 Fitur Utama

- ✅ Sinkronisasi user dari SIAKAD ke Moodle
- ✅ Pembuatan course + enroll otomatis via API
- ✅ Manual enrolment user ke course
- ✅ Assign role (teacher/student)

---

## 📁 Struktur Project

```bash
.
├── apispec.json                  # Dokumentasi Swagger API
├── config/
│   └── env.go                    # Konfigurasi environment (load API key, token Moodle)
├── contracts/
│   └── moodle_user_getter.go     # Interface kontrak untuk pengambilan user
├── controllers/
│   └── moodle_controller.go      # Handler Gin untuk setiap endpoint Moodle
├── main.go                       # Entry point aplikasi
├── middleware/
│   └── auth_middleware.go        # Middleware autentikasi API Key
├── model/
│   └── web/
│       ├── moodle_create_course_request.go
│       ├── moodle_create_course_response.go
│       ├── moodle_create_course_with_enroll_request.go
│       ├── moodle_create_course_with_enroll_response.go
│       ├── moodle_exception.go
│       ├── moodle_manual_enroll_request.go
│       ├── moodle_role_assigment_request.go
│       ├── moodle_status_response.go
│       ├── moodle_user_create_request.go
│       ├── moodle_user_create_response.go
│       ├── moodle_user_getByField_request.go
│       ├── moodle_user_getByField_response.go
│       ├── moodle_user_sync_request.go
│       ├── moodle_user_update_request.go
│       ├── moodle_user_update_response.go
│       └── web_response.go       # Struktur umum response API
├── README.md                     # Dokumentasi proyek
├── routes/
│   └── routes.go                 # Definisi semua endpoint dan group route Gin
├── services/
│   ├── moodle_service.go         # Interface untuk MoodleService
│   └── moodle_service_impl.go    # Implementasi business logic MoodleService
└── utils/
    ├── helpers/
    │   └── moodle_helpers.go     # Helper untuk membuat form Moodle, request builder, dll
    └── validation/
        └── moodle_validation.go  # Validasi business logic (duplikat user, dll)

````

---

## ⚙️ Requirement

* Go versi **1.24+**
* Moodle versi **5.0+**
* Token Web Service aktif di Moodle (Admin > Web services > Manage tokens)
* API key environment variable (`API_SECRET_KEY`)

---

## 🛠️ Instalasi

### 1. Clone Project

```bash
git clone https://github.com/pptikunida/moodle-api
cd moodle-api
```

### 2. Buat File `.env`

Contoh `cp .env.example .env`:

```env
MOODLE_API_URL=https://elearning.unida.gontor.ac.id/webservice/rest/server.php
MOODLE_TOKEN=your_moodle_token_here
APP_PORT=YOUR_PORT
API_SECRET_KEY=your_custom_api_key
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Jalankan Aplikasi

```bash
go run main.go
```

---

## 📦 Dependency Penting

* [`github.com/gin-gonic/gin`](https://github.com/gin-gonic/gin) - HTTP Web Framework
* [`github.com/joho/godotenv`](https://github.com/joho/godotenv) - Load environment variables
* [`github.com/swaggo/gin-swagger`](https://github.com/swaggo/gin-swagger) - Swagger UI untuk dokumentasi API

---

## 🔐 Keamanan / Middleware

* Setiap endpoint diproteksi dengan `API Key` melalui **header**:

```http
X-API-Key: your_custom_api_key
```

---

## 📌 Endpoint Penting

| Method | Endpoint                                    | Deskripsi                               |
| ------ | ------------------------------------------- | --------------------------------------- |
| `GET`  | `/api/moodle/site-info`                     | Informasi Moodle & token                |
| `POST` | `/api/moodle/users/sync`                    | Sinkronisasi user SIAKAD ke Moodle      |
| `POST` | `/api/moodle/roles/assign`                  | Assign role ke user                     |
| `POST` | `/api/moodle/courses/create-with-enrolment` | Buat course dan langsung assign teacher |
| `POST` | `/api/moodle/courses/enrol/manual`          | Enroll user ke course sebagai student   |

---

## 📄 Dokumentasi API Swagger

Swagger UI tersedia di:

```
http://localhost:8080/swagger/index.html
```

Atau import **Swagger JSON** spec: `./apispec.json` ke https://editor.swagger.io/

---

## 🧪 Testing API

Gunakan Postman atau curl:

```bash
curl -X POST http://localhost:8080/api/moodle/users/sync \
  -H 'Content-Type: application/json' \
  -H 'X-API-Key: your_api_key' \
  -d '{ "username": "john.doe", "email": "...", ... }'
```

---

## 👥 Kontributor

* Rizky Cahyono Putra – [@rizkycahyono97](https://github.com/rizkycahyono97)
* pptikunida - [@pptikunida](https://github.com/pptikunida)

---