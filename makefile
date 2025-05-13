up:
	# docker-compose -f pkg/kafka/docker-compose.yml -p kafka -up -d --build
	docker-compose -f auth/.deploy/local/docker-compose.yml -p auth_service up -d --build
	docker-compose -f learning/.deploy/local/docker-compose.yml -p learning_service up -d --build
	docker-compose -f gateway/.deploy/local/docker-compose.yml -p gateway_service up -d --build
	# ...

down:
	docker-compose -f gateway/.deploy/local/docker-compose.yml down
	docker-compose -f learning/.deploy/local/docker-compose.yml down
	docker-compose -f auth/.deploy/local/docker-compose.yml down
	# docker-compose -f pkg/kafka/docker-compose.yml down
	# ...