PHP_SERVICE=php
PHP_CLI="docker-compose run -it --rm $(PHP_SERVICE) "

bash:
	@"$(PHP_CLI)"bash

composer:
	@"$(PHP_CLI)"frankenphp php-cli /usr/local/bin/composer ${c}
