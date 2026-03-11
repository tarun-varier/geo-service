# GPS Location Microservice - Specification

## 1. Project Overview

**Project Name:** classroom-gps-service
**Type:** REST API Microservice
**Core Functionality:** Exposes GPS coordinates of the classroom via HTTP API, containerized with Docker, deployed on Minikube, monitored by Prometheus, with CI/CD via GitHub Actions
**Target Users:** Students, instructors, and any application needing classroom location data

---

## 2. Technical Stack

- **Language:** Go (Golang)
- **Framework:** Gin (HTTP router)
- **Container:** Docker
- **Orchestration:** Minikube (Kubernetes)
- **Monitoring:** Prometheus + node-exporter
- **CI/CD:** GitHub Actions

---

## 3. GPS Coordinates

**Classroom Location:** Room 301, Computer Science Building
- **Latitude:** 40.7128
- **Longitude:** -74.0060
- **Location:** New York University, Brooklyn (example classroom)

---

## 4. API Specification

### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/location | Returns GPS coordinates in JSON |
| GET | /health | Health check endpoint |
| GET | /metrics | Prometheus metrics endpoint |

### Response Format (GET /api/v1/location)

```json
{
  "room": "Room 301",
  "building": "Computer Science Building",
  "latitude": 40.7128,
  "longitude": -74.0060,
  "accuracy": 10,
  "timestamp": "2026-03-11T12:00:00Z"
}
```

---

## 5. Kubernetes Deployment

- **Deployment:** Single replica deployment
- **Service:** NodePort service for external access
- **Resources:** CPU/Memory limits defined

---

## 6. Prometheus Monitoring

- Metrics exposed at `/metrics` endpoint
- Default port: 8080
- Scraping interval: 15s

---

## 7. CI/CD Pipeline (GitHub Actions)

- **Build:** Compile Go application
- **Test:** Run unit tests
- **Docker Build:** Build and push Docker image
- **Deploy:** Deploy to Minikube (manual trigger)

---

## 8. File Structure

```
lab7/
├── SPEC.md
├── Dockerfile
├── main.go
├── go.mod
├── go.sum
├── kubernetes/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── configmap.yaml
├── prometheus/
│   └── prometheus.yml
└── .github/
    └── workflows/
        └── ci-cd.yml
```
