# ---- build stage ----
FROM golang:1.27-alpine AS build

# Set workspace directory
WORKDIR /app

# Copy dependency definitions first (enables Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/api ./cmd/api

# ---- runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /bin/api /api

ENV PORT=8080 \
    ENV=production

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]
