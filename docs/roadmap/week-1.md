# Week 1: Data Core & API Definition

## Tasks

- Define the JSON schemas for dynamic surveys and respondent metadata.
- Boilerplate a Go (Gin/Echo) API for the /api/v1/mobile-hook/ endpoint to handle USSD/SMS sessions.
- Deliverable: A functional API that accepts survey submissions and manages session state.

## Implementation

This is the functional foundation for **Biami**. Following the 13-week roadmap, Week 1 focuses on the **Data Core** and the **Mobile Webhook**.

We will use the **Gin** framework for the API due to its performance and middleware ecosystem, which is highly regarded in the Canadian cloud-native space.

### 1. Project Structure
To demonstrate architectural maturity, we use a standard Go layout that separates the entry point from the internal logic.

```text
biami/
├── cmd/
│   └── api/
│       └── main.go          # Entry point
├── internal/
│   ├── handler/
│   │   └── webhook.go       # USSD/SMS Webhook logic
│   └── model/
│       └── survey.go        # Dynamic JSON schemas
├── Dockerfile               # Multi-stage production build
├── Makefile                 # Developer automation
└── go.mod
```

---

### 2. The Implementation

#### `internal/model/survey.go`
This defines the dynamic nature of Biami. We use `json.RawMessage` to allow for unstructured survey data, a key requirement for an autonomous platform.

```go
package model

import "encoding/json"

// SurveyResponse represents the incoming data from a mobile channel
type SurveyResponse struct {
	SessionID string          `json:"session_id"`
	Mobile    string          `json:"mobile"`
	Channel   string          `json:"channel"` // USSD, SMS, PWA
	Payload   json.RawMessage `json:"payload"` // Dynamic survey answers
}

// MobileHookRequest represents the raw trigger from a gateway like Twilio/Infobip
type MobileHookRequest struct {
	From    string `json:"from"`
	Text    string `json:"text"`
	Session string `json:"session_id"`
}
```

#### `internal/handler/webhook.go`
This handles the logic for the `/api/v1/mobile-hook/` endpoint.

```go
package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"biami/internal/model"
)

func MobileHookHandler(c *gin.Context) {
	var req model.MobileHookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// TODO: Add Logic to check MobileSession state in Redis (Week 4)
	// Example USSD Response format
	response := "CON Welcome to Biami. \n1. Start Survey \n2. Check Balance"

	c.String(http.StatusOK, response)
}
```

#### `cmd/api/main.go`
The application entry point.

```go
package main

import (
	"biami/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/mobile-hook", handler.MobileHookHandler)
	}

	r.Run(":8080")
}
```

---

### 3. Architect’s "Pro" Tooling
To align with the **Solutions Architect** goal, include these two files to show you design for deployment and automation.

#### `Dockerfile` (Multi-stage Build)
```dockerfile
# Build Stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o biami-api ./cmd/api/main.go

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/biami-api .
EXPOSE 8080
CMD ["./biami-api"]
```

#### `Makefile`
```makefile
.PHONY: run build test docker-build

run:
	go run cmd/api/main.go

build:
	go build -o bin/biami cmd/api/main.go

test:
	go test ./...

docker-build:
	docker build -t biami-api:latest .
```

---

### Next Steps for Week 1
1. **Initialize the module:** Run `go mod init biami` and `go get github.com/gin-gonic/gin`.
2. **Test the Webhook:** Use `curl` to simulate a USSD gateway hit:
   ```bash
   curl -X POST http://localhost:8080/api/v1/mobile-hook \
   -H "Content-Type: application/json" \
   -d '{"from": "+123456789", "text": "*123#", "session_id": "xyz123"}'
   ```

By the end of this hour, you have a container-ready Go API that follows professional directory standards. Ready to document the first **ADR (Architecture Decision Record)** for the repository?