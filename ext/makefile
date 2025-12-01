GOLANG_SERVICE=golang
GOLANG_CLI="docker-compose exec -it $(GOLANG_SERVICE) "

up:
	docker-compose up -d

stop:
	docker-compose stop --timeout 1

down:
	docker-compose down --timeout 1

bash:
	@"$(GOLANG_CLI)"bash

go:
	@"$(GOLANG_CLI)"go ${c}

test:
	@"$(GOLANG_CLI)"go test ./internal/...

ext-init:
	GEN_STUB_SCRIPT=/usr/local/lib/php/build/gen_stub.php frankenphp extension-init sconcur.go

ext-build:
	 rm -f build/sconcur.so build/sconcur.h && CGO_CFLAGS=$(php-config --includes) go build -buildmode=c-shared -o build/sconcur.so .

ext-test:
	  php -d extension=./build/sconcur.so -r "echo \SConcur\Extension\ping('${c}') . PHP_EOL;"