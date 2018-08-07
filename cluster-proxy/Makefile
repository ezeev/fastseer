build:
	@docker build -t ezeev/fastseer:proxy-v1 -f Dockerfile .

deploy:
	docker push ezeev/fastseer:proxy-v1
	kubectl delete -f deploy.yaml
	kubectl create -f deploy.yaml

replace:
	kubectl delete -f deploy.yaml
	kubectl create -f deploy.yaml

expose:
	kubectl expose deployment nginx --target-port=80 --type=LoadBalancer