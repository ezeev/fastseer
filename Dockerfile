FROM golang:1.10.2

WORKDIR /app
COPY fastseer fastseer
COPY template template
COPY config-prod.yaml config.yaml
COPY secret secret
COPY shopify-admin-ui/build/ shopify-admin-ui/build/ 

#COPY ./public/index.html public/index.html
#COPY ./public/script.js public/script.js
#COPY ./public/style.css public/style.css
#CMD ["/app/fastseer"]
EXPOSE 8082
ENTRYPOINT ["/app/fastseer", "config.yaml"]