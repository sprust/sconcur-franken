GOLANG_SERVICE=golang
GOLANG_CLI="docker-compose exec -it $(GOLANG_SERVICE) "

up:
	docker-compose up -d

stop:
	docker-compose stop

bash:
	@"$(GOLANG_CLI)"bash

go:
	@"$(GOLANG_CLI)"go ${c}

test:
	@"$(GOLANG_CLI)"go test ./internal/...

ext-init:
	GEN_STUB_SCRIPT=/usr/local/lib/php/build/gen_stub.php frankenphp extension-init sconcur.go

ext-build:
	rm -rf build/frankenphp && \
	CGO_ENABLED=1 \
    XCADDY_GO_BUILD_FLAGS="-ldflags='-w -s' -tags=nobadger,nomysql,nopgx,nowatcher" \
    CGO_CFLAGS=$(php-config --includes) \
    CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
    xcaddy build \
        --output build/frankenphp \
        --with github.com/dunglas/frankenphp/caddy \
        --with github.com/dunglas/mercure/caddy \
        --with github.com/dunglas/vulcain/caddy \
        --with github.com/sprust/sconcur-franken=.