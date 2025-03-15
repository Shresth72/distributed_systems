# ProtoBuf
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

test:
	go test -race ./

# Logger
VALUE ?= TGV0J3MgR28gIzIK
OFFSET ?= 0
post_log:
	curl -X POST localhost:6969 -d \
	'{"record": {"value": "${VALUE}"}}'

get_log:
	curl -X GET localhost:6969 -d '{"offset": ${OFFSET}}'

# Test
test_log:
	cd internal/log && \
	go test .
