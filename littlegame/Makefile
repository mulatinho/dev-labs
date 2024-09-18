all: build

clean:
	rm -fv littlegame data.db *~ *.swp

build:
	go get ./...
	GOOS=linux GOARCH=amd64 go build -o littlegame .

image:
	docker build -t littlegame:latest .

push:
	docker build -t littlegame:latest .
	docker tag littlegame:latest mulatinho/littlegame:latest
	docker push mulatinho/littlegame:latest

run:
	docker run -p 8181:8181 littlegame:latest

tests:
	GIN_MODE=debug go test ./...
