GOLANG_SERVICE=golang
GOLANG_CLI="docker-compose run -it --rm $(GOLANG_SERVICE) "

bash:
	@"$(GOLANG_CLI)"bash

go:
	@"$(GOLANG_CLI)"go ${c}

test:
	@"$(GOLANG_CLI)"go test ./internal/...
