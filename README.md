## Architectures

- [ ] Client <=> Envoy <=> Go gRPC <=> Go Kafka client <=> Kafka

- [x] Client <=> Go/HTTP/Websockets <=> Go Kafka client <=> Kafka

- [ ] Client <=> Confluent Rest API <=> Kafka

## Usage

1. Clone this repo

2. `cd` into clone repo

3. docker-compose up -d

4. `cd` into `backend`

5. `make competition competition=COMPETITION-NAME` to create a Kafka topic named "COMPETITION-NAME"

6. `make delete-competition competition=COMPETITION-NAME` to delete the Kafka topic named "COMPETITION-NAME"

#### HTTP

1. `make serve` to run the backend server. This will run in the foreground.

2. `make http-get-votes competition=COMPETITION-NAME` to listen for votes using a cURL Websockets connection. This will run in the foreground.

3. `make http-vote competition=COMPETITION-NAME singer=SINGER` to vote using cURL POST request

#### Command line

1. `make vote competition=COMPETITION-NAME singer=SINGER` to vote for specified singer competing in specified competition

2. `make get-votes competition=COMPETITION-NAME` to consume the votes for the specified competition (TODO: count the votes)
