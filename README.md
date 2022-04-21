# Golange-grpc-gateway example app

### How to run

Run the grpc server
```bash
cd server
go run main.go
```

Run the proxy http
```bash
cd proxy
go run main.go
```

```bash
curl -X 'POST' \
  'http://localhost:8081/echo' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "value": "Peter"
}'
```