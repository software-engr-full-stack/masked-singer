_port := $$(cat ./kafka.properties | grep bootstrap\.servers | cut -f 2 -d '='| cut -f 2 -d ':')

vote:
	go run producer.go config.go ${competition} ${singer}

get-votes:
	go run consumer.go config.go ${competition}

competition:
	docker-compose exec broker \
	  kafka-topics --create \
	    --topic ${competition} \
	    --bootstrap-server localhost:${_port} \
	    --replication-factor 1 \
	    --partitions 1

delete-competition:
	docker-compose exec broker \
	  kafka-topics --delete \
	    --topic ${competition} \
	    --bootstrap-server localhost:${_port}

list-topics:
	docker-compose exec broker \
	  kafka-topics --list \
	    --bootstrap-server localhost:${_port}

sh:
	docker exec --interactive --tty broker bash

.PHONY: vote get-votes topic delete-topic sh
