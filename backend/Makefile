_port := $$(cat ./kafka.properties | grep bootstrap\.servers | cut -f 2 -d '='| cut -f 2 -d ':')

_build_file := ./build/masked-singer

_build:
	go build -o ${_build_file} *.go

serve: _build
	if type nodemon >/dev/null; then \
		nodemon --signal SIGTERM --ext go --exec "${_build_file} 'serve'"; \
	else \
		${_build_file} 'serve'; \
	fi

http-vote:
	curl --header "Content-Type: application/json" \
			 --request POST \
			 --data '{"action": "vote", "competition_name": "${competition}", "singer_name": "${singer}"}' \
			 http://localhost:8082/vote

vote: _build
	${_build_file} 'vote' ${competition} ${singer}

get-votes: _build
	${_build_file} 'get-votes' ${competition} ${singer}

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

reset-competition:
	make delete-competition competition=${competition}; \
	make competition competition=${competition}

sh:
	docker exec --interactive --tty broker bash

.PHONY: vote get-votes competition delete-competition list-topics sh _build
