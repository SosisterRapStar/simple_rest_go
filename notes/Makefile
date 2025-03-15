BASE_URL := http://localhost:8888/api/v1

build:
	docker-compose up -d

run:
	go run main.go

testOneThread:
ifndef ENDPOINT
	@echo "Testing default endpoint: $(BASE_URL)/notes/?limit=2"
	wrk -t 1 -c 1 -d 300s --latency "$(BASE_URL)/notes/?limit=2"
else
	@echo "Testing custom endpoint: $(BASE_URL)$(ENDPOINT)"
	wrk -t 1 -c 1 -d 300s --latency "$(BASE_URL)$(ENDPOINT)"
endif

testMultiThread:
ifndef ENDPOINT
	@echo "Testing default endpoint: $(BASE_URL)/notes/?limit=2"
	wrk -t 10 -c 400 -d 5m --latency "$(BASE_URL)/notes/?limit=2"
else
	@echo "Testing custom endpoint: $(BASE_URL)$(ENDPOINT)"
	wrk -t 10 -c 400 -d 5m --latency "$(BASE_URL)$(ENDPOINT)"
endif
