# Entity Relationship Diagram (ERD)

Berikut adalah struktur database utama yang digunakan pada platform CollabForge.

ERD ini dirancang untuk mendukung:
- project collaboration
- contribution tracking
- portfolio automation
- AI feedback system
- organization management

---

# Core Entities

## users

Menyimpan data akun pengguna platform.

### Fungsi
- register & login
- identitas contributor
- pemilik organization
- anggota project

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key user |
| first_name | varchar | nama depan |
| last_name | varchar | nama belakang |
| mid_name | varchar | nama tengah |
| username | varchar | username unik |
| email | varchar | email user |
| password | varchar | password hash |
| provider | enum(local, google) | metode login |
| provider_id | varchar | id provider OAuth |
| bio | text | deskripsi user |
| avatar_url | varchar | foto profil |
| created_at | datetime | waktu dibuat |
| update_at | datetime | waktu diperbarui |

---

## organizations

Menyimpan data organisasi/tim/project owner.

### Fungsi
- wadah project
- manajemen member
- representasi komunitas/tim

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key organization |
| owner_id | varchar | owner organization |
| name | varchar | nama organization |
| slug | varchar | slug unik organization |
| description | text | deskripsi organization |
| logo_url | varchar | logo organization |
| create_at | datetime | waktu dibuat |
| update_at | datetime | waktu diperbarui |

---

## organization_members

Relasi antara user dan organization.

### Fungsi
- menyimpan anggota organization
- menyimpan role user dalam organization

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| organization_id | varchar FK | relasi ke organizations |
| user_id | varchar FK | relasi ke users |
| role_id | varchar | role member |
| joined_at | varchar | tanggal bergabung |

---

## role_desc

Menyimpan deskripsi role dalam organization.

### Fungsi
- manajemen role organization
- custom permission role

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| name | varchar | nama role |
| description_job | varchar | deskripsi pekerjaan |
| organization_id | varchar FK | relasi ke organizations |

---

## projects

Menyimpan data project collaboration.

### Fungsi
- project collaboration
- contribution tracking
- task management

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key project |
| organization_id | varchar FK | relasi ke organizations |
| created_by | varchar | creator project |
| title | varchar | judul project |
| slug | varchar FK | slug project |
| description | varchar | deskripsi project |
| status | enum | status project |
| visibility | enum | public/private project |
| start_date | datetime | tanggal mulai |
| end_date | datetime | tanggal selesai |
| created_at | datetime | waktu dibuat |
| update_at | datetime | waktu diperbarui |

### Status Project

```text
draft
open
ongoing
completed
archive
```

### Visibility

```text
public
private
```

---

## project_member

Relasi antara user dan project.

### Fungsi
- contributor project
- role contributor project

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| project_id | varchar FK | relasi ke projects |
| user_id | varchar FK | relasi ke users |
| role_id | varchar FK | relasi ke role_desc |
| joined_at | datetime | tanggal bergabung |

---

## tasks

Menyimpan task dalam project.

### Fungsi
- workflow project
- assignment contributor
- progress tracking

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| project_id | varchar | relasi ke project |
| created_by | varchar | creator task |
| assigned_to | varchar | assignee task |
| title | varchar | judul task |
| description | varchar | deskripsi task |
| status | enum | status task |
| priority | varchar | prioritas task |
| due_date | varchar | deadline task |
| update_at | string | update terakhir |
| created_at | string | waktu dibuat |

### Status Task

```text
todo
in_progress
review
done
```

---

## task_submission

Menyimpan hasil pengerjaan task oleh contributor.

### Fungsi
- submission hasil kerja
- review hasil task
- contribution tracking

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| task_id | varchar FK | relasi ke tasks |
| user_id | varchar FK | relasi ke users |
| submission_text | varchar | deskripsi submission |
| submission_url | varchar | link submission |
| status | varchar | status submission |
| feedback | varchar | feedback reviewer |
| submitted_at | datetime | waktu submit |
| reviewed_at | datetime | waktu review |

---

## ai_feedback

Menyimpan feedback AI terhadap submission contributor.

### Fungsi
- AI evaluation
- scoring contribution
- automated review

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| submission_id | varchar FK | relasi ke task_submission |
| feedback | varchar | feedback AI |
| score | int | skor AI |
| created_at | datetime | waktu dibuat |

---

## portfolios

Menyimpan portofolio otomatis hasil kontribusi project.

### Fungsi
- generated portfolio
- contribution showcase
- portfolio tracking

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| user_id | varchar FK | relasi ke users |
| project_id | varchar FK | relasi ke projects |
| title | varchar | judul portfolio |
| description | varchar | deskripsi portfolio |
| role | varchar | role contributor |
| generate_at | varchar | waktu generate portfolio |

---

## slugs

Menyimpan slug unik untuk URL entity.

### Fungsi
- URL clean
- SEO friendly URL
- unique routing

### Fields

| Field | Type | Keterangan |
|---|---|---|
| id | varchar | primary key |
| name | varchar | slug unik |

---

# Relasi Utama

```text
users
 ├── organization_members
 ├── project_member
 ├── task_submission
 └── portfolios

organizations
 ├── organization_members
 ├── role_desc
 └── projects

projects
 ├── project_member
 ├── tasks
 └── portfolios

tasks
 └── task_submission

task_submission
 └── ai_feedback
```

---

# Tujuan Desain ERD

ERD ini dirancang agar:
- scalable
- modular
- AI-agent friendly
- mudah di-maintain
- mendukung raw SQL
- mendukung contribution tracking
- mendukung portfolio automation
- mendukung collaborative workflow