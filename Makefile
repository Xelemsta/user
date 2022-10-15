run:
	docker-compose rm --stop && docker-compose build && docker-compose up --force-recreate

stop:
	docker-compose rm --stop

build:
	docker-compose build

test:
	# XXX: improve
	export POSTGRES_USER=postgres \
	export POSTGRES_PASSWORD=postgres \
	export POSTGRES_DB=postgres \
	export POSTGRES_PORT=5432 \
	export POSTGRES_HOST=localhost \
	export BROKER_KAFKA=127.0.0.1:29092 \
	export KAFKA_TOPIC=user && \
	docker-compose rm --stop && \
	docker-compose build && \
	docker-compose up -d --force-recreate && \
	go clean -testcache && \
	go test -cover -p 1 ./...
