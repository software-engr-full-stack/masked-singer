## Usage

1. Clone this repo

2. `cd` into clone repo

3. docker-compose up -d

4. `make topic` to create a "votes" Kafka topic

5. `make vote singer=NAME-OF-SINGER` to vote for specified singer

6. `make get-votes` to consume the votes (TODO: count the votes)

7. `make delete-topic` to delete the "votes" Kafka topic

