clean:
	rm -fv tasks *~ *.swp

build:
	go get ./...
	go build -o ./tasks cmd/main.go

image:
	docker build -t tasks:latest .

push:
	docker tag tasks:latest mulatinho/tasks:latest
	docker push mulatinho/tasks:latest

run:
	docker run -p 8181:8181 tasks:latest

tests:
	GIN_MODE=debug go test
