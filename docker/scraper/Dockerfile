FROM golang:alpine

RUN apk add openssl

WORKDIR /usr/src/app

COPY . .

RUN openssl genrsa -out /usr/src/app/keys/private.pem 2048
RUN openssl rsa -in /usr/src/app/keys/private.pem -outform PEM -pubout -out /usr/src/app/keys/public.pem

EXPOSE 80

CMD ["go", "run", "main.go"]