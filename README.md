# SafeSchool Data Monitor

## 📖 Project Overview

SafeSchool Data Monitor is a cloud-native microservices system that simulates the real-time monitoring of student online activities.  
Built with Go, GCP, and automated CI/CD pipelines, it reflects a real-world architecture similar to the systems developed by companies like Qoria.

## 🏗 Architecture

- **Backend API** — receives activity data (Go)
- **Pub/Sub Queue** — manages message passing
- **Worker Service** — processes and stores activities in BigQuery
- **Database** — Google BigQuery
- **Cache** — Redis (optional for performance)
- **Deployment** — GCP Cloud Run via Docker and GitHub Actions

## 🚀 Features

- RESTful API (POST/GET activities)
- Asynchronous processing with Pub/Sub
- Unit and Integration Tests
- API Test Automation (Postman / k6)
- Full CI/CD Pipeline with GitHub Actions
- Dockerized Services
- Cloud Deployment (GCP)

## 🛠 Tech Stack

- Golang
- Google Cloud Platform (BigQuery, Pub/Sub, Cloud Run)
- Redis
- GitHub Actions
- Docker

## 📦 Setup

### Prerequisites
- Go 1.21+
- Docker
- GCP account
- Redis (optional for local dev)

### Local Run

```bash
git clone https://github.com/yourusername/safeschool-datamonitor.git
cd backend
go run main.go
