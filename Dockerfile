FROM alpine:3.17.3
WORKDIR /app
COPY ./my-app /app/my-app

ENTRYPOINT [ "/app/my-app" ]
