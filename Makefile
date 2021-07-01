

all: orch adapter svc_dummy svc_db

proto:
	protoc -I=. --go_out=. --go-grpc_out=. adapter.proto
	protoc -I=. --go_out=. --go-grpc_out=. svc_dummy.proto
	protoc -I=. --go_out=. --go-grpc_out=. svc_db.proto

app: proto
	go mod download
orch: app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/orch/ ./app/orch
adapter: app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/adapter/ ./app/adapter
svc_dummy: app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/svc_dummy/ ./app/svc_dummy
svc_db: app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/svc_db/ ./app/svc_db


run:
	cd build
	docker-compose up