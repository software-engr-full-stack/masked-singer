_port := $$(cat ./kafka.properties | grep bootstrap\.servers | cut -f 2 -d '='| cut -f 2 -d ':')

_topic_name := $$(cat ./config.sh | grep topic_name | cut -f 2 -d '=')

vote:
	go run producer.go config.go ${singer}

get-votes:
	go run consumer.go config.go

topic:
	docker-compose exec broker \
	  kafka-topics --create \
	    --topic ${_topic_name} \
	    --bootstrap-server localhost:${_port} \
	    --replication-factor 1 \
	    --partitions 1

delete-topic:
	docker-compose exec broker \
	  kafka-topics --delete \
	    --topic ${_topic_name} \
	    --bootstrap-server localhost:${_port}

sh:
	docker exec --interactive --tty broker bash

.PHONY: topic
