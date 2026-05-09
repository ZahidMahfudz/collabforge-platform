## About This

CollabForge adalah platform portofolio berbasis kontribusi nyata dalam proyek.

Platform ini dirancang untuk membantu membangun portofolio yang lebih jujur, terukur, dan dapat diverifikasi berdasarkan aktivitas dan kontribusi pengguna di dalam proyek yang mereka kerjakan.

CollabForge menggabungkan beberapa konsep utama:
- Project Management
- Portfolio Automation
- Contribution Tracking
- Recruitment & Talent Discovery
- AI Assisted Project Workflow

Setiap task, kontribusi, dan aktivitas proyek dapat menjadi bagian dari portofolio pengguna secara otomatis, sehingga profil yang ditampilkan tidak hanya berdasarkan klaim, tetapi berdasarkan bukti kontribusi nyata.

Fokus utama platform ini adalah membangun sistem reputasi dan portofolio berbasis proof of work, sehingga recruiter maupun organisasi dapat menilai kemampuan pengguna berdasarkan kontribusi aktual di dalam proyek.

<br />

## Tech Stack

| Category            | Technology                    |
| ------------------- | ----------------------------- |
| Language            | Golang                        |
| Framework           | Fiber                         |
| Database            | PostgreSQL                    |
| Database Driver     | pgx                           |
| Query Style         | Raw SQL                       |
| Authentication      | PASETO                        |
| Password Hashing    | bcrypt                        |
| Architecture        | Simple Clean Architecture     |
| API Style           | REST API                      |
| Environment Config  | godotenv                      |
| Logger              | Logrus                        |
| AI Development Flow | AI Agent Assisted Development |
| Containerization    | Docker *(planned)*            |

<br />

## Struktur Folder
```
collabforge-platform/
|
|__ 0_docs/                     # dokumentasi 
|   |
|   |__ API contract/           # Dokumentasi API dan API Contract
|   |
|   |__ ERD/                    # Dokumentasi ERD
|   |
|   |__ planning/               # Perintah-perintah eksekusi AI Agent untuk generate code
|
|__ cmd/
|   |__ app/
|       |__ main.go             # entry point aplikasi (init config, DI, run server)
|
|__ config/                     # konfigurasi global aplikasi
|
|__ internal/                   # core logic aplikasi (tidak boleh diakses dari luar)
|   |
|   |__ controller/             # layer yang mengatur request dan response
|   |
|   |__ dto/                    # data transfer object (kontrak API)
|   |   |
|   |   |__ request/            # struktur input dari client + validasi
|   |   |
|   |   |__ response/           # struktur output ke client
|   |
|   |__ entity/                 # Representasi tabel / model database
|   |
|   |__ middleware/             # middleware (auth, logging, error handler)
|   |
|   |__ repository/             # implementasi repository (query ke database pakai pgx)
|   |
|   |__ routes/                 # Daftar endpoint (dalam Group)
|   |
|   |__ service/                # integrasi eksternal (bukan core logic)
|   |
|   |__ usecase/                # Core logic aplikasi
|
|__ utils/                      # shared utilities (reusable)
|
|__ migrations/                 # file migration database (versioning schema)
|
|__ .env                        # konfigurasi environment
|__ go.mod                      # dependency management
|__ go.sum
```