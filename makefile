PHP_SERVICE=php
PHP_CLI="docker-compose run --rm $(PHP_SERVICE) "

bash-php:
	@"$(PHP_CLI)"bash