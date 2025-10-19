# Log Analyzer

This project is a backend service built with **Golang** using the **Clean Architecture** pattern.  
It provides secure authentication with **JWT** and **Google Authenticator (TOTP)**, along with APIs for uploading and analyzing log files concurrently.

---

## Features

- User registration and login with password  
- Optional 2FA using **Google Authenticator (TOTP)**  
- JWT authentication for protected endpoints  
- Concurrent log file parsing using goroutines  
- CRUD operations for log analysis  
- Clean Architecture with clear separation of layers  

---



## Authentication Flow

The authentication uses **JWT** and optional **Google Authenticator (TOTP)** for two-factor authentication.

### 1. Register User
**Endpoint:**
POST /api/register

**Body:**
json
{
  "email": "user@example.com",
  "password": "mypassword",
  "name": "John Doe"
}

Response:

{
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}

2. Setup TOTP (Google Authenticator)
Endpoint:
POST /api/totp/setup/:id


Response:
{
  "secret": "JBSWY3DPEHPK3PXP",
  "otpauth_uri": "otpauth://totp/LogAnalyzer:user@example.com?secret=JBSWY3DPEHPK3PXP&issuer=LogAnalyzer"
}
Scan the otpauth_uri using Google Authenticator app to generate your 6-digit code.

3. Verify TOTP

Endpoint:

POST /api/totp/verify/:id


Body:

{
  "code": "123456"
}


Response:

{
  "enabled": true
}

4. Login

Endpoint:

POST /api/login


Body (with TOTP):

{
  "email": "user@example.com",
  "password": "mypassword",
  "totp": "123456"
}


Response:

{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "totp_enabled": true
  }
}


Use this token for all protected endpoints.

Log Upload & Analysis

Once logged in, you can upload and analyze log files.

Upload Log File

Endpoint:

POST /api/upload/


Header:

Authorization: Bearer <JWT_TOKEN>


Body (form-data):

file: <your-log-file.log>


Response:

{
  "message": "file uploaded and analyzed"
}

Example Log Format

Each line in the log file is expected to follow this simplified format:

192.168.1.10 200 0.123
192.168.1.11 404 0.045
192.168.1.12 500 0.320


The system will count:

Total requests

Number of errors (non-200 responses)

Unique IP addresses

Average response time

Technologies Used

Go (Golang)

Gin Web Framework

GORM ORM

PostgreSQL

JWT Authentication

Google Authenticator (TOTP)

Clean Architecture

Running the Project

Clone the repository:

git clone https://github.com/ifs21014-itdel/log-analyzer.git
cd log-analyzer


Install dependencies:

go mod tidy


Create .env:

DB_USER=postgres
DB_PASS=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=log_analyzer
JWT_SECRET=your_secret
APP_NAME=LogAnalyzer


Run:

go run cmd/main.go


Server runs on:

http://localhost:8080

License

MIT License © 2025 Dedi Panggabean
