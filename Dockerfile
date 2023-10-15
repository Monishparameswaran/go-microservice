FROM golang:1.20.10-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go mod download 
RUN go build main.go
# FROM gcr.io/distroless/base-debian11
FROM istio/distroless:1.19-2023-10-13T03-27-30
WORKDIR /myapp
COPY --from=builder /app/main ./
EXPOSE 8070 
CMD [ "/myapp/main" ]