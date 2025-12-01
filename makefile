FRANKEN_SERVICE=franken
FRANKEN_CLI="docker-compose exec -it $(FRANKEN_SERVICE) "

up:
	docker-compose up -d

stop:
	docker-compose stop --timeout 1

down:
	docker-compose down --timeout 1

bash:
	@"$(FRANKEN_CLI)"bash

composer:
	@"$(FRANKEN_CLI)"frankenphp php-cli /usr/local/bin/composer ${c}
