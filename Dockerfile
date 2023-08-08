FROM alpine:3.18.3
WORKDIR /app
COPY ./my-app /app/my-app

ENTRYPOINT [ "/app/my-app" ]
