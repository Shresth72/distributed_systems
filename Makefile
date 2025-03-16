# ProtoBuf
compile_protoc:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

test_protoc:
	go test -race ./

# Logger
VALUE ?= TGV0J3MgR28gIzIK
OFFSET ?= 0
post_log:
	curl -X POST localhost:6969 -d \
	'{"record": {"value": "${VALUE}"}}'

get_log:
	curl -X GET localhost:6969 -d '{"offset": ${OFFSET}}'

# Run Tests
test_log:
	cd internal/log && \
	go test .

test_server:
	cd internal/server && \
	go test .

# Generate CA Certs
CONFIG_PATH=.proglog/

.PHONY: ca_init
ca_init:
	mkdir -p ${CONFIG_PATH}

.PHONY: ca_gencert
ca_gencert:
	cfssl gencert \
		-initca cert/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=cert/ca-config.json \
		-profile=server \
		cert/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=cert/ca-config.json \
		-profile=client \
		cert/client-csr.json | cfssljson -bare client
	mv *.pem *.csr ${CONFIG_PATH}

.PHONY: ca_test
ca_test:
	go test -race ./...
