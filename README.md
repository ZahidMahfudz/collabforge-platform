## About This
Ini adalah aplikasi backend dari platform CollabForge — sebuah platform portofolio berbasis kontribusi nyata dalam proyek.

CollabFOrge menggabungkan:
* Project Management
* Portfolio Automation
* Contribution Tracking

Dengan tujuan menghasilkan portofolio yang jujur, terukur, dan dapat diverifikasi.

<br />

## Struktur Folder
```
collabforge-platform/
|
|__ cmd/
|   |__ app/
|       |__ main.go              # entry point aplikasi (init config, DI, run server)
|
|__ config/                     # konfigurasi global aplikasi
|   |__ database.go             # koneksi PostgreSQL (pgx)
|   |__ env.go                  # load environment variables
|   |__ logger.go               # inisialisasi logger
|
|__ internal/                   # core logic aplikasi (tidak boleh diakses dari luar)
|   |
|   |__ domain/                 # layer paling murni (tidak tergantung framework)
|   |   |__ entity/             # representasi data utama (User, Project, Task, ActivityLog)
|   |   |__ repository/         # interface repository (kontrak ke database)
|
|   |__ usecase/                # business logic utama (otak aplikasi)
|
|   |__ repository/             # implementasi repository (query ke database pakai pgx)
|
|   |__ delivery/               # layer komunikasi dengan client
|   |   |__ http/
|   |       |__ controller/     # handler request (parse, validasi, panggil usecase)
|   |       |__ route/          # definisi endpoint API (routing)
|
|   |__ dto/                    # data transfer object (kontrak API)
|   |   |__ request/            # struktur input dari client + validasi
|   |   |__ response/           # struktur output ke client
|
|   |__ middleware/             # middleware (auth, logging, error handler)
|
|   |__ service/                # integrasi eksternal (bukan core logic)
|       |__ oauth/              # login OAuth (Google, dll)
|       |__ ai/                 # AI generator (project, task, workflow)
|
|__ pkg/                        # shared utilities (reusable)
|   |__ utils/                  # helper umum
|   |__ response/               # format standar response API
|   |__ errors/                 # custom error handling
|   |__ constant/               # enum/konstanta (project_type, role, dll)
|
|__ migrations/                # file migration database (versioning schema)
|
|__ docs/                      # dokumentasi (ERD, business rules, API contract)
|
|__ .env                       # konfigurasi environment
|__ go.mod                     # dependency management
|__ go.sum
```