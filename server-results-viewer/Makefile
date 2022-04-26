_port := $$(cat ./kafka.properties | grep bootstrap\.servers | cut -f 2 -d '='| cut -f 2 -d ':')

_build_file := ./build/masked-singer

_http_port := 8082

_build:
	go build -o ${_build_file} *.go

serve:
	if type nodemon >/dev/null; then \
		nodemon --signal SIGTERM --ext go --exec "make _build && ${_build_file} 'serve' ${_http_port}"; \
	else \
		make _build && ${_build_file} 'serve' ${_http_port}; \
	fi

http-vote:
	curl --header "Content-Type: application/json" \
			 --request POST \
			 --data '{"action": "vote", "competition_name": "${competition}", "singer_name": "${singer}"}' \
			 http://localhost:${_http_port}/vote

http-get-votes:
	curl --include \
	     --no-buffer \
	     --header "Connection: Upgrade" \
	     --header "Upgrade: websocket" \
	     --header "Host: localhost:${_http_port}" \
	     --header "Origin: http://localhost:${_http_port}" \
	     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
	     --header "Sec-WebSocket-Version: 13" \
	     http://localhost:${_http_port}/get-votes\?competition_name\=${competition}

vote: _build
	${_build_file} 'vote' ${competition} ${singer}

get-votes: _build
	${_build_file} 'get-votes' ${competition}

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

.PHONY: vote get-votes competition delete-competition list-topics sh _build serve http-vote http-get-votes
