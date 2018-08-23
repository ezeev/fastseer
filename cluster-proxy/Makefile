VERSION=v4

build:
	@docker build -t ezeev/fastseer:proxy-$(VERSION) -f Dockerfile .


deploy:
	@docker build -t ezeev/fastseer:proxy-$(VERSION) -f Dockerfile .
	@docker push ezeev/fastseer:proxy-$(VERSION)
	kubectl set image deployment/nginx nginx=ezeev/fastseer:proxy-$(VERSION)


expose:
	kubectl expose deployment nginx --target-port=80 --type=LoadBalancer