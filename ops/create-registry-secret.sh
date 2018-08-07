# Create secret to authenticate w/ private repository
DOCKER_REGISTRY_SERVER=docker.io
DOCKER_USER=ezeev
DOCKER_EMAIL=evan.pease@gmail.com
DOCKER_PASSWORD=REPLACE_ME

kubectl create secret docker-registry fastseer-docker-repo-key \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL
