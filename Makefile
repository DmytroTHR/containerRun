STARTING_APP := "service_a"
CLIENT_TO_START := "service_b"
STOP_LIST := "docker ps --filter name=${CLIENT_TO_START}* -q"

run-app:
	@docker-compose pull ${STARTING_APP} ${CLIENT_TO_START}
	@docker-compose run --rm -T ${STARTING_APP}
	@echo "removing pattern service_b container.."
	@docker rm ${CLIENT_TO_START}

stop-containers:
	@docker stop $$(exec "${STOP_LIST}")

stop-app: stop-containers
	@docker-compose down