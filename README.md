# Backend Booking Ticket

Aplikasi backend untuk sistem pemesanan tiket, dibangun menggunakan **Go (Golang)** dengan pendekatan **Clean Architecture** dan **Docker**.

## Tech Stack

- **Language**: Go 1.24
- **Framework**: [Fiber](https://gofiber.io/) (Express-like performance)
- **Database**: PostgreSQL (via Neon Tech)
- **ORM**: [GORM](https://gorm.io/)
- **Authentication**: Bcrypt (untuk hashing password)
- **Containerization**: Docker & Docker Compose

## Struktur Project (Clean Architecture)

```
.
├── cmd/api             # Entry point aplikasi (main.go)
├── config              # Load .env dan koneksi Database
├── internal
│   ├── entity          # Definisi Struct (User, Booking)
│   ├── handler         # HTTP Handlers (Menerima Request)
│   ├── repository      # Akses ke Database (Query SQL)
│   ├── service         # Logika Bisnis (Hitung harga, Hash password)
│   └── routes          # Konfigurasi URL Routing API
├── pkg                 # Helper / Utilities
└── docker-compose.yml  # Orchestration Container
```

## Cara Menjalankan

### Prasyarat
- Docker & Docker Compose
- (Opsional) Go 1.22+ jika ingin run tanpa Docker

### Setup & Run (Rekomendasi)

1. **Clone Repository**
   ```bash
   git clone https://github.com/abimanyupewe/backend-boking-ticket.git
   cd backend-boking-ticket
   ```

2. **Setup Environment**
   Buat file `.env` di root folder:
   ```env
   DATABASE_URL=postgresql://user:password@host/dbname?sslmode=require
   ```

3. **Jalankan dengan Docker**
   ```bash
   docker compose up -d --build
   ```
   Aplikasi akan berjalan di `http://localhost:3000`.

## API Endpoints

### Autentikasi (/api/auth)

| Method | Endpoint | Deskripsi | Body (JSON) |
| :--- | :--- | :--- | :--- |
| POST | `/register` | Mendaftar user baru | `{"name": "...", "email": "...", "password": "..."}` |
| POST | `/login` | Masuk ke aplikasi | `{"email": "...", "password": "..."}` |

### Booking (/api/bookings)

| Method | Endpoint | Deskripsi | Body (JSON) |
| :--- | :--- | :--- | :--- |
| POST | `/` | Membuat booking baru | `{"user_id": "...", "ticket_id": "...", "amount": 1}` |
| GET | `/` | Melihat semua booking | - |
| GET | `/:id` | Detail booking | - |

## Cara Test Manual (Curl)

**Register:**
```bash
curl -X POST http://localhost:3000/api/auth/register \
-H "Content-Type: application/json" \
-d '{"name": "User", "email": "user@test.com", "password": "123"}'
```

**Booking:**
```bash
curl -X POST http://localhost:3000/api/bookings \
-H "Content-Type: application/json" \
-d '{"user_id": "ID_USER_DARI_REGISTER", "ticket_id": "TICKET-01", "amount": 2}'
```
