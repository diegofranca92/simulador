FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin/:${PATH}"

# instala a lib pro kafka rodar
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

# deixa o container de pé
CMD ["tail", "-f", "/dev/null"]