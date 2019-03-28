.PHONY: generate
generate:
	go generate ./...

.PHONY: unit
unit:
	if [ ! -d coverage ]; then mkdir coverage; fi	
	go test --cover -v -race -coverprofile=coverage/golangphoenix-tests.profile -timeout 30s ./...

# Run make unit first!
.PHONY: coverage_results
coverage_results: 
	go tool cover -html=coverage/golangphoenix-tests.profile