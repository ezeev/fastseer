VERSION=indexworker-v4

clean:
	rm -f fastseer

run:
	go build -o indexworker handlers.go main.go model.go routes.go server.go
	./indexworker


deploy:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o indexworker handlers.go main.go model.go routes.go server.go
	@docker build -t ezeev/fastseer:$(VERSION) -f Dockerfile .
	@docker push ezeev/fastseer:$(VERSION)
	kubectl set image deployment/fs-index-worker fs-index-worker=ezeev/fastseer:$(VERSION)

expose:
	kubectl expose deployment fs-index-worker --target-port=8083 --type=NodePort



