clean:
	rm -fv tasks data.db *~ *.swp

build:
	go get ./...
	GOOS=linux GOARCH=amd64 go build -o tasks .

image:
	docker build -t tasks:latest .

push:
	docker build --build-arg GIN_MODE=release -t tasks:latest .
	docker tag tasks:latest mulatinho/tasks:latest
	docker push mulatinho/tasks:latest

run:
	docker run -p 8181:8181 tasks:latest

tests:
	GIN_MODE=debug go test ./...
