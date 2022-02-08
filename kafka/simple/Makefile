# Start the kafka service
kafka-run:
	docker-compose up -d

# Stop it
kafka-stop:
	docker-compose down

# Create events topic
create-events:
	docker exec kafka-go kafka-topics.sh \
		--bootstrap-server localhost:9092 --create \
		--replication-factor 1 --partitions 1 \
		--topic events

# Produce a message for events topic
produce-event:
	docker exec kafka-go /opt/bitnami/kafka/bin/kafka-console-producer.sh \
		--bootstrap-server 127.0.0.1:9092 \
		--topic new-message

run-producer:
	go run producer.go 127.0.0.1:9092 events

run-consumer:
	go run consumer.go 127.0.0.1:9092 group events

tidy:
	go mod tidy -v

# Stop Kafkadrop
kafdrop-stop:
	docker stop  kafka-go-kafdrop

# Run Kafkadrop
kafdrop-run:
	docker run --name kafka-go-kafdrop --rm -d -p 9000:9000 \
		-e KAFKA_BROKERCONNECT="10.5.0.1:9092" \
		-e JVM_OPTS="-Xms32M -Xmx256M" --network kafka-go \
		-e SERVER_SERVLET_CONTEXTPATH="/" \
		obsidiandynamics/kafdrop:latest

# Run Kafkadrop (same as above but for production environments)
kafdrop-production:
	docker run --rm -p 9000:9000 \
		-e KAFKA_BROKERCONNECT=$(node -e "process.stdout.write($(terraform output kafka_bootstrap_brokers))") \
		-e JVM_OPTS="-Xms128M -Xmx2G" \
		-e KAFKA_PROPERTIES=$(echo security.protocol=SSL | base64) \
		-e SERVER_SERVLET_CONTEXTPATH="/" \
		obsidiandynamics/kafdrop:latest

