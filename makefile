PHP_SERVICE=php
PHP_CLI="docker-compose run -it --rm $(PHP_SERVICE) "

bash:
	@"$(PHP_CLI)"bash

composer:
	@"$(PHP_CLI)"frankenphp php-cli /usr/local/bin/composer ${c}

ext-init:
	GEN_STUB_SCRIPT=/usr/local/sconcur/php-src/build/gen_stub.php frankenphp extension-init /app/ext/build/extension/sconcur.go