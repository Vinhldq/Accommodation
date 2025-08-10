# 🏨 Accommodation Booking System

A comprehensive accommodation booking platform built with modern technologies, featuring a Go backend and Angular frontend.

## 📋 Overview

This project is a full-stack accommodation booking system that allows users to:
- Browse and search accommodations
- Make bookings and manage reservations
- Handle payments and reviews
- Admin panel for managing properties and users
- Real-time analytics and statistics

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.23.2
- **Framework**: Gin Gonic
- **Database**: MySQL 8.0
- **Cache**: Redis 7.0
- **Message Queue**: Apache Kafka 4.0
- **Authentication**: JWT
- **Documentation**: Swagger
- **Monitoring**: Prometheus + Grafana
- **Database Migration**: Goose
- **SQL Generation**: SQLC

### Frontend
- **Framework**: Angular 19.2
- **UI Library**: Taiga UI, PrimeNG
- **Maps**: Google Maps API
- **Language**: TypeScript

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Development**: Air (hot reload)

## 🚀 Quick Start

### Prerequisites
- Go 1.23.2+
- Node.js 18+
- Docker & Docker Compose
- Make (for Windows: `choco install make`)

### 1. Clone the Repository
```bash
git clone https://github.com/Vinhldq/DoAnChuyenNganh.git
cd DoAnChuyenNganh
```

### 2. Backend Setup

Navigate to backend directory:
```bash
cd backend
```

Start infrastructure services (MySQL, Redis, Kafka):
```bash
make docker-build
```

Run database migrations:
```bash
make db-up
```

Start the backend server:
```bash
make dev
```

The backend will be available at `http://localhost:8080`

**API Documentation**: `http://localhost:8080/swagger/index.html`

### 3. Frontend Setup

Navigate to frontend directory:
```bash
cd frontend
```

Install dependencies:
```bash
make npm-install
```

Start the development server:
```bash
make dev
```

The frontend will be available at `http://localhost:4200`

## 📁 Project Structure

```
├── backend/
│   ├── cmd/server/           # Application entry point
│   ├── internal/
│   │   ├── controllers/      # HTTP handlers
│   │   ├── services/         # Business logic
│   │   ├── repo/            # Data access layer
│   │   ├── middlewares/     # HTTP middlewares
│   │   └── database/        # Generated SQL code
│   ├── sql/
│   │   ├── queries/         # SQL queries
│   │   └── schemas/         # Database migrations
│   └── docker-compose.yaml  # Infrastructure setup
├── frontend/
│   └── src/
│       ├── app/             # Angular components
│       ├── assets/          # Static assets
│       └── environments/    # Environment configs
└── README.md
```

## 🐳 Docker Services

The application uses the following services:
- **MySQL** (Port 3307): Main database
- **Redis** (Port 6379): Caching and sessions
- **Kafka** (Port 9092/9094): Message queue
- **Prometheus** (Port 9090): Metrics collection
- **Grafana** (Port 3000): Monitoring dashboard

## 📖 Available Make Commands

### Backend Commands
```bash
make dev                    # Start development server
make docker-build          # Build and start all services
make docker-up             # Start existing containers
make docker-down           # Stop all containers
make db-up                 # Run database migrations
make db-down               # Rollback last migration
make db-reset              # Reset database
make sql-gen               # Generate SQL code from queries
make swagger               # Generate API documentation
```

### Frontend Commands
```bash
make npm-install           # Install dependencies
make dev                   # Start development server
```

## 🔧 Development Tools

### Database Migration
This project uses [Goose](https://github.com/pressly/goose) for database migrations:

```bash
# Create new migration
make create-migration name=create_users_table

# Apply migrations
make db-up

# Rollback migration
make db-down
```

### Hot Reload
Backend uses Air for hot reloading during development. Frontend uses Angular CLI's built-in hot reload.

## 🏗️ Key Features

- **User Management**: Registration, authentication, profile management
- **Accommodation Management**: Property listings, rooms, facilities
- **Booking System**: Real-time availability, reservations, order management
- **Payment Integration**: Secure payment processing
- **Review System**: User reviews and ratings
- **Admin Dashboard**: Complete admin panel with analytics
- **File Upload**: Image and document management
- **Real-time Notifications**: Using Kafka for messaging
- **Monitoring**: Prometheus metrics and Grafana dashboards

## 📊 API Endpoints

Key API endpoints include:
- `/api/auth/*` - Authentication
- `/api/accommodations/*` - Property management
- `/api/bookings/*` - Reservation system
- `/api/payments/*` - Payment processing
- `/api/reviews/*` - Review system
- `/api/admin/*` - Admin operations
- `/api/stats/*` - Analytics

Full API documentation is available at `/swagger/index.html` when running the backend.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is for educational purposes as part of a graduation thesis (Đồ Án Chuyên Ngành).

## 📞 Support

For any questions or issues, please contact the development team or create an issue in the repository.