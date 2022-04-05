## Architecture

Client => Envoy => gRPC => Kafka?

## Usage

1. Clone this repo

2. `cd` into clone repo

3. docker-compose up -d

4. `make competition competition=COMPETITION-NAME` to create a Kafka topic named "COMPETITION-NAME"

5. `make vote competition=COMPETITION-NAME singer=SINGER` to vote for specified singer competing in specified competition

6. `make get-votes competition=COMPETITION-NAME` to consume the votes for the specified competition (TODO: count the votes)

7. `make delete-competition competition=COMPETITION-NAME` to delete the Kafka topic named "COMPETITION-NAME"
