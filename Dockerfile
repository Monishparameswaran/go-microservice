FROM golang:latest as builder
WORKDIR /app
COPY go.sum .
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go
FROM gcr.io/distroless/base-debian11
WORKDIR /serve
COPY --from=builder /app/main ./
EXPOSE 8070 
ENTRYPOINT [ "/serve/main" ]