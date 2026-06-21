test:
	docker-compose -f docker-compose-test.yml up -d
	go test -tags=integration ./...
	docker-compose -f docker-compose-test.yml down