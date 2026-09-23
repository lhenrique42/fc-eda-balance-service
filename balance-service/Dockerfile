FROM golang:1.25

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["go", "run", "./cmd/balanceservice"]
