STARTING_APP := "service_a"
CLIENT_TO_START := "service_b"
STOP_LIST := "docker ps --filter name=${CLIENT_TO_START}* -q"

run-app: build-app
	@docker-compose run --rm -T ${STARTING_APP}

remove-client-containers:
	@docker stop $$(exec "${STOP_LIST}")

stop-app: remove-client-containers
	@docker-compose down

build-app:
	@docker-compose up --build --no-start --remove-orphans