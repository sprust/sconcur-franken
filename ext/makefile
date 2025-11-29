GOLANG_SERVICE=golang
GOLANG_CLI="docker-compose run -it --rm $(GOLANG_SERVICE) "

bash:
	@"$(GOLANG_CLI)"bash

go:
	@"$(GOLANG_CLI)"go ${c}

test:
	@"$(GOLANG_CLI)"go test ./internal/...

build-franken:
	CGO_ENABLED=1 \
    XCADDY_GO_BUILD_FLAGS="-ldflags='-w -s' -tags=nobadger,nomysql,nopgx" \
    CGO_CFLAGS=$(php-config --includes) \
    CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
    xcaddy build \
        --output frankenphp \
        --with github.com/dunglas/frankenphp/caddy \
        --with github.com/dunglas/mercure/caddy \
        --with github.com/dunglas/vulcain/caddy \
        --with github.com/sprust/sconcur-franken=.