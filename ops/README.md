# Deploying a Kubernetes Cluster for FastSeer

Pre-requisites:

- Digital Ocean Account
- Containership Account
- Private Dockerhub Repo (I created https://hub.docker.com/r/ezeev/fastseer/)

Initial Steps:

- Follow the Containership steps to deploy a cluster on Digital Ocean - https://cloud.containership.io
- Run create-*-secret.sh - this will create some secrets:
    - `fastseer-docker-repo-key` - used for pulling images from the private repo

Deployment Tips For New Containers:

- You will need this section in your yaml when specifying containers in the private repo:

```
apiVersion: v1
kind: Pod
metadata:
  name: whatever
spec:
  containers:
    - name: whatever
      image: ezeev/fastseer:v1
      imagePullPolicy: Always
  imagePullSecrets:
    - name: fastseer-docker-repo-key
```

# Networking

In the cluster-proxy directory that contains all of the resources for an nginx service. It is exposed as a LoadBalancer. The nginx service serves as a router to other services in the cluster. Rather than go through the hassle of configuring letsencrypt, I opted to use Cloudflare. Following their steps was very simple. I just had to change the Nameservers at my registrar (Namecheap.com). A major bonus of Cloudflare is that it looks like you can setup as many CNAME records as you want. I have shopify-app.fastseer.com point to the app. I plan on having the public site go to fastseer.com. 


# Development Workflow

- Start ngrok to serve the API. `ngrok http 8082` - i.e. https://d3b7936c.ngrok.io
- Start admin-ui react app on port 3000 (or 3001 if taken) . From shopify-admin-ui: `npm start`
- Open react app in browser, add appDomain param, set it to equal the ngrok url i.e. ?appDomain=https://d3b7936c.ngrok.io. NOTE: appDomain is only required when not running in the shopify iframe
  - the app also requires auth request params, for fastseer-staging, you can add this: `?appDomain=https://d3b7936c.ngrok.io&shop=fastseer-staging.myshopify.com&hmac=0c58f35abe62e10fecf92a396b047a3848af95a61b31d15d2435be7b4d687468&timestamp=1533675516&locale=en`
  - So, a full URL for doing react dev work will look like: http://localhost:3001/?appDomain=https://d3b7936c.ngrok.io&shop=fastseer-staging.myshopify.com&hmac=0c58f35abe62e10fecf92a396b047a3848af95a61b31d15d2435be7b4d687468&timestamp=1533675516&locale=en
  - If these params aren't working you can get new ones by opening the app through the fastseer-staging shop. The request logs will contain a full, valid hmac query string. 

See https://docs.google.com/drawings/d/14urnrOMoVUtb9pKHA5hyZW47QpcTMiPG1KZcoIsJKBo/edit?usp=sharing for visual illustration of the above.

# Deployment Workflow

So, you've just finished making changes to API server and the app, and wish to push them.

1. Run tests - `go test -v`
  - if a test fails, be sure to check that the auth params are up to date. Update testAuthParams constant in global_test.go if necessary.
2. If any React changes were made, build them. cd into `shopify-admin-ui` and run `npm run build`
3. In the Makefile in root, increment the version number at the top.
4. Run `make deploy`. This will build the binary, docker image, push to docker hub, then update the container image in the running K8s deployment.

