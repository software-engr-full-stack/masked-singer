## Architectures

Client <=> Envoy <=> Go gRPC <=> Go Kafka client <=> Kafka

Client <=> Go/HTTP/Websockets <=> Go Kafka client <=> Kafka

Client <=> Confluent Rest API <=> Kafka

## Usage

1. Clone this repo

2. `cd` into clone repo

3. docker-compose up -d

4. `cd` into `backend`

5. `make competition competition=COMPETITION-NAME` to create a Kafka topic named "COMPETITION-NAME"

6. `make vote competition=COMPETITION-NAME singer=SINGER` to vote for specified singer competing in specified competition

7. `make get-votes competition=COMPETITION-NAME` to consume the votes for the specified competition (TODO: count the votes)

8. `make delete-competition competition=COMPETITION-NAME` to delete the Kafka topic named "COMPETITION-NAME"
