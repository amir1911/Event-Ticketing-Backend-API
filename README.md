# Secure the Crowd
## _Secure Event Ticketing Backend API_

Backend REST API untuk sistem **pemesanan tiket event online** dengan fokus pada **keamanan**, **autentikasi**, dan **validasi stok tiket secara real-time**.

Project ini dikembangkan menggunakan **Golang + Fiber + GORM** sebagai implementasi **Studi Kasus II – Secure the Crowd**.

---

## ✨ Overview

**Secure the Crowd** adalah backend service yang menangani:

- Autentikasi user menggunakan **JWT**
- Otorisasi berbasis role (**Admin & User**)
- Manajemen Event (Admin only)
- Pemesanan tiket oleh user yang sudah login
- Validasi stok tiket secara **real-time**
- Dokumentasi API menggunakan **Postman (Public)**

API ini dirancang agar **scalable**, **maintainable**, dan **aman** untuk traffic tinggi.

---

## 🚀 Features

- User Registration & Login (JWT Authentication)
- Role-based Access Control (RBAC)
- Admin-only Event Management
- Ticket Booking (Login Required)
- Real-time Ticket Stock Validation
- RESTful API Design

---

## 🛠 Tech Stack

- **Golang**
- **Fiber**
- **GORM**
- **MySQL**
- **JWT (Authentication)**
- **Postman (API Documentation)**

---

## ⚙️ Environment Configuration
Buat file `.env` di root project:

## ▶️ Running the Project
`go run main.go`

## 📮 API Documentation
Untuk melihat Endpoint dapat klik link ini:
https://documenter.getpostman.com/view/50460563/2sBXVZpaPy#d9cb811f-fffe-4a5b-bca2-43e3d3ddb188

## 👨‍💻 Author
**Amirullah**


## 📂 Project Structure

```text
go-belajar/
├── controllers/
│   ├── auth_controller.go
│   ├── user_controller.go
│   ├── event_controller.go
│   └── booking_controller.go
├── middleware/
│   ├── auth.go
│   └── admin.go
├── models/
│   ├── user.go
│   ├── event.go
│   └── booking.go
├── database/
│   └── database.go
├── routes/
│   └── api.go
├── config/
│   └── env.go
├── .env
├── main.go
├── go.mod
└── go.sum



