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
