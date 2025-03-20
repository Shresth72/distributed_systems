# ProtoBuf
compile_protoc:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

# Logger
VALUE ?= TGV0J3MgR28gIzIK
OFFSET ?= 0
post_log:
	curl -X POST localhost:6969 -d \
	'{"record": {"value": "${VALUE}"}}'

get_log:
	curl -X GET localhost:6969 -d '{"offset": ${OFFSET}}'

# Run Tests
CLEAR ?= true
test_log:
	cd internal/log && \
	go test .

clean_test_metrics:
	rm test/logs/metrics-*.log test/logs/traces-*.log

test_server:
	cd internal/server && \
	go test . -v -debug=true
ifneq ($(CLEAR), false)
	$(MAKE) clean_test_metrics
endif

# Generate CA Certs
CERT_PATH=.certs
CONFIG_DIR=test
CONFIG_FILES := model.conf policy.csv

.PHONY: init
init:
	mkdir -p ${CERT_PATH}

.PHONY: copy-configs
copy-configs: init
	for file in $(CONFIG_FILES); do \
		cp -u ${CONFIG_DIR}/$$file ${CERT_PATH}/$$file; \
	done

.PHONY: gencert
gencert: init
	@test -f ${CERT_PATH}/ca.pem || cfssl gencert \
		-initca ${CONFIG_DIR}/ca-csr.json | cfssljson -bare ${CERT_PATH}/ca
	@test -f ${CERT_PATH}/server.pem || cfssl gencert \
		-ca=${CERT_PATH}/ca.pem \
		-ca-key=${CERT_PATH}/ca-key.pem \
		-config=${CONFIG_DIR}/ca-config.json \
		-profile=server \
		${CONFIG_DIR}/server-csr.json | cfssljson -bare ${CERT_PATH}/server
	@test -f ${CERT_PATH}/root-client.pem || cfssl gencert \
		-ca=${CERT_PATH}/ca.pem \
		-ca-key=${CERT_PATH}/ca-key.pem \
		-config=${CONFIG_DIR}/ca-config.json \
		-profile=client \
		-cn="root" \
		${CONFIG_DIR}/client-csr.json | cfssljson -bare ${CERT_PATH}/root-client
	@test -f ${CERT_PATH}/nobody-client.pem || cfssl gencert \
		-ca=${CERT_PATH}/ca.pem \
		-ca-key=${CERT_PATH}/ca-key.pem \
		-config=${CONFIG_DIR}/ca-config.json \
		-profile=client \
		-cn="nobody" \
		${CONFIG_DIR}/client-csr.json | cfssljson -bare ${CERT_PATH}/nobody-client

# Test
.PHONY: test
test: init copy-configs gencert
	go test -race ./...
