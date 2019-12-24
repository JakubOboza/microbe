GIT_COMMIT := $(shell git rev-parse HEAD)

build:
	go build -o bin/microbe ./cmd/microbe
	go build -o bin/microscope ./cmd/microscope

test:
	go test -v ./pkg/microbe/

docker:
	@echo LATEST COMMIT IS $(GIT_COMMIT)
	docker build -t jakuboboza/microbe:latest -t jakuboboza/microbe:$(GIT_COMMIT) .
	docker push jakuboboza/microbe:latest
	docker push jakuboboza/microbe:$(GIT_COMMIT)

