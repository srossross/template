from alpine

RUN apk update && apk add curl openssl
RUN curl --fail https://srossross.github.io/template/get.sh | sh
