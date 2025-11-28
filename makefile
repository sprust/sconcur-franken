PHP_SERVICE=php
PHP_CLI="docker-compose run -it --rm $(PHP_SERVICE) "

GOLANG_SERVICE=golang
GOLANG_CLI="docker-compose run -it --rm $(GOLANG_SERVICE) "

bash-php:
	@"$(PHP_CLI)"bash

bash-golang:
	@"$(GOLANG_CLI)"bash

composer:
	@"$(PHP_CLI)"frankenphp php-cli /usr/local/bin/composer ${c}
