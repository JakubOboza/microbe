build:
	go build -o bin/microbe ./cmd/microbe
	go build -o bin/microscope ./cmd/microscope

docker:
	docker build -t jakuboboza/microbe:latest .
	docker push jakuboboza/microbe:latest
