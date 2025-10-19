```markdown
# Log Analyzer

A robust backend service built with **Golang** using **Clean Architecture** principles. This project provides secure authentication with **JWT** and **Google Authenticator (TOTP)**, along with powerful APIs for uploading and analyzing log files concurrently.

---

## Features

- User registration and login with secure password handling
- Optional Two-Factor Authentication (2FA) using **Google Authenticator (TOTP)**
- JWT authentication for protected endpoints
- Concurrent log file parsing using goroutines
- Complete CRUD operations for log analysis
- Clean Architecture with clear separation of concerns
- Real-time log analytics and statistics

---

## Authentication Flow

The authentication system uses **JWT tokens** and optional **Google Authenticator (TOTP)** for enhanced security.

### 1. Register User

**Endpoint:**
```http
POST /api/register
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "mypassword",
  "name": "John Doe"
}
```

**Response:**
```json
{
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

---

### 2. Setup TOTP (Google Authenticator)

**Endpoint:**
```http
POST /api/totp/setup/:id
```

**Response:**
```json
{
  "secret": "JBSWY3DPEHPK3PXP",
  "otpauth_uri": "otpauth://totp/LogAnalyzer:user@example.com?secret=JBSWY3DPEHPK3PXP&issuer=LogAnalyzer"
}
```

**Note:** Scan the `otpauth_uri` QR code using the Google Authenticator app to generate your 6-digit verification codes.

---

### 3. Verify TOTP

**Endpoint:**
```http
POST /api/totp/verify/:id
```

**Request Body:**
```json
{
  "code": "123456"
}
```

**Response:**
```json
{
  "enabled": true
}
```

---

### 4. Login

**Endpoint:**
```http
POST /api/login
```

**Request Body (with TOTP):**
```json
{
  "email": "user@example.com",
  "password": "mypassword",
  "totp": "123456"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "totp_enabled": true
  }
}
```

**Important:** Use this token in the Authorization header for all protected endpoints.

---

## Log Upload & Analysis

Once logged in, you can upload and analyze log files through the API.

### Upload Log File

**Endpoint:**
```http
POST /api/upload/
```

**Header:**
```http
Authorization: Bearer <JWT_TOKEN>
```

**Body (form-data):**
```
file: <your-log-file.log>
```

**Response:**
```json
{
  "message": "file uploaded and analyzed"
}
```

---

### Example Log Format

Each line in the log file should follow this format:

```
192.168.1.10 200 0.123
192.168.1.11 404 0.045
192.168.1.12 500 0.320
```

**The system analyzes:**
- Total requests
- Number of errors (non-200 responses)
- Unique IP addresses
- Average response time

---

## Technologies Used

- **Go (Golang)** - Programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Database
- **JWT** - Authentication tokens
- **Google Authenticator (TOTP)** - Two-factor authentication
- **Clean Architecture** - Design pattern

---

## Getting Started

### Prerequisites

- Go 1.19 or higher
- PostgreSQL
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/ifs21014-itdel/log-analyzer.git
cd log-analyzer
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file in the root directory:
```env
DB_USER=postgres
DB_PASS=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=log_analyzer
JWT_SECRET=your_secret_key
APP_NAME=LogAnalyzer
```

4. Run the application:
```bash
go run cmd/main.go
```

5. The server will start at:
```
http://localhost:8080
```

---

## Project Structure

```
log-analyzer/
├── cmd/
│   └── main.go
├── internal/
│   ├── domain/
│   ├── usecase/
│   ├── repository/
│   └── delivery/
├── pkg/
├── .env
├── go.mod
└── README.md
```

---

## License

MIT License © 2025 Dedi Panggabean

---

## Author

**Dedi Panggabean**

GitHub: [@ifs21014-itdel](https://github.com/ifs21014-itdel)

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
```
