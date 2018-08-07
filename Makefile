VERSION=v3

clean:
	rm -f fastseer

run:
	go build -o fastseer fastseer-server/fastseer.go
	./fastseer config-stage.yaml


deploy:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o fastseer fastseer-server/fastseer.go
	@docker build -t ezeev/fastseer:$(VERSION) -f Dockerfile .
	@docker push ezeev/fastseer:$(VERSION)
	kubectl set image deployment/fastseer fastseer=ezeev/fastseer:$(VERSION)




