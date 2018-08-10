FROM golang:1.10.2

WORKDIR /app
COPY indexworker indexworker

#COPY ./public/index.html public/index.html
#COPY ./public/script.js public/script.js
#COPY ./public/style.css public/style.css
#CMD ["/app/fastseer"]
EXPOSE 8083
ENTRYPOINT ["/app/indexworker"]