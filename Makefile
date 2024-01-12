.PHONY: build run lint
build:
	go build -o dist/golang-project-layout ./cmd/golang-project-layout/

run:
	dist/golang-project-layout server

lint:
	golint ./...

.PHONY: api_gen api_install_dep api_clean
api_install_dep:
	go env -w GOPROXY=https://goproxy.cn,direct
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/jstemmer/go-junit-report@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts@latest
	go install github.com/rakyll/statik@latest

api_gen:
	protoc -I . -I third_party \
		--go_out=paths=source_relative:. \
        		--go-grpc_out=paths=source_relative:. \
        		--grpc-gateway_out=paths=source_relative:. \
        		--grpc-gateway-ts_out=paths=source_relative:./dist/sdk/ \
        		--openapiv2_out=logtostderr=true:. \
        		--openapiv2_opt allow_merge=true \
        		--openapiv2_opt output_format=json \
        		--openapiv2_opt merge_file_name="golang-gateway-project-layout." \
		api/golang-project-layout/v1/golang-project-layout.proto api/general/v1/demo.proto
	cp -R *.swagger.json docs/swagger-ui/golang-gateway-project-layout.swagger.json

api_clean:
	rm -f api/*/*/*.pb.go api/*/*/*.pb.gw.go api/*/*/*.swagger.json api/*/*/*.pb.validate.go
	rm -rf dist/sdk/*
	rm -rf docs/swagger-ui/*.swagger.json
