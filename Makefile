build:
	go build -o bin/microbe ./cmd/microbe

docker:
	docker build -t jakuboboza/microbe:latest .
	docker push jakuboboza/microbe:latest