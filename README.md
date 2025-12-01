# build extension
```shell
cp /app/ext/main.go /app/ext/sconcur.go && \
GEN_STUB_SCRIPT=/usr/local/lib/php/build/gen_stub.php frankenphp extension-init /app/ext/sconcur.go
rm -rf /app/ext/build/frankenphp-sconcur && \
CGO_ENABLED=1 \
XCADDY_GO_BUILD_FLAGS="-ldflags='-w -s' -tags=nobadger,nomysql,nopgx,nowatcher" \
CGO_CFLAGS="-D_GNU_SOURCE $(php-config --includes)" \
CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
xcaddy build \
    --output /app/ext/build/frankenphp-sconcur \
    --with github.com/dunglas/frankenphp/caddy \
    --with github.com/dunglas/mercure/caddy \
    --with github.com/dunglas/vulcain/caddy \
    --with github.com/sprust/sconcur-franken=/app/ext
```