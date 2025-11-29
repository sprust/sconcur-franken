PHP_SERVICE=php
PHP_CLI="docker-compose exec -it $(PHP_SERVICE) "

up:
	docker-compose up -d

stop:
	docker-compose stop --timeout 1

down:
	docker-compose down --timeout 1

bash:
	@"$(PHP_CLI)"bash

composer:
	@"$(PHP_CLI)"ext/build/frankenphp php-cli /usr/local/bin/composer ${c}

php:
	@"$(PHP_CLI)"ext/build/frankenphp php-cli ${c}