build:
	go build -o bin/microbe ./cmd/microbe
	go build -o bin/microscope ./cmd/microscope

test:
	go test -v ./pkg/microbe/
	go test -v ./pkg/microscope/

docker:
	docker build -t jakuboboza/microbe:latest .
	docker push jakuboboza/microbe:latest

