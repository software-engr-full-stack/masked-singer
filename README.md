Masked Singer
=============

## How to run

```bash
# Run Kafka on docker (See docker-compose.yml for details)
docker-compose up

# Run 'localhost:8082' server
cd server-votes
go run -tags dynamic .

# Try a vote
curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"action": "vote", "competition_name": "contest-1", "singer_name": "michael"}' \
	http://localhost:8082/vote
```

## Kafka commands on docker

```bash
# add a topic (See kafka.properties for the port number)
docker-compose exec broker \
  kafka-topics --create \
    --topic contest-1 \
    --bootstrap-server localhost:9092 \
    --replication-factor 1 \
    --partitions 1

# delete a topic
docker-compose exec broker \
  kafka-topics --delete \
    --topic contest-1 \
    --bootstrap-server localhost:9092

# list topics
docker-compose exec broker \
  kafka-topics --list \
    --bootstrap-server localhost:9092

# list topic messages
docker-compose exec broker \
    kafka-console-consumer --bootstrap-server localhost:9092 \
    --topic contest-1 --from-beginning
```
