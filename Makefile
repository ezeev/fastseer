VERSION=v14

clean:
	rm -f fastseer

run:
	go build -o fastseer exec/main.go
	./fastseer config-stage.yaml

docs:
	swagger -mainApiFile=exec/main.go -output=./docs/API.md -format=markdown -apiPackage=github.com/ezeev/fastseer

deploy:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o fastseer exec/main.go
	@docker build -t ezeev/fastseer:$(VERSION) -f Dockerfile .
	@docker push ezeev/fastseer:$(VERSION)
	kubectl set image deployment/fastseer fastseer=ezeev/fastseer:$(VERSION)