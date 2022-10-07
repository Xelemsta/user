#!/bin/sh

export POSTGRES_USER=postgres \
       POSTGRES_PASSWORD=postgres \
       POSTGRES_DB=postgres \
       POSTGRES_PORT=5432 \
       POSTGRES_HOST=localhost \
       BROKER_KAFKA=127.0.0.1:29094 \
       KAFKA_TOPIC=user