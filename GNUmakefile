export SIMULATION_PEERS=130
export SIMULATION_COUNT=3
export SIMULATION_SEED=
export SIMULATION_NETWORK_FAILURE_PERCENT=0
export SIMULATION_NODE_FAILURE_PERCENT=85

simulation: ## run the simulation test
	mkdir -p data
	go test -v ./simulation

plot: ## plot the normal output
	mkdir -p plot
	./bin/plot-degree "In Degree" "in-active" $(SIMULATION_COUNT) > plot/in-degree.png
	./bin/plot-stacked $(SIMULATION_COUNT) > plot/gossip.png

plot-slow: ## more plots, but these are slow
	mkdir -p plot
	./bin/plot-all-graphs

test: message-generated.go
	go test

message-generated.go: message.go message.go.genny
	ln message.go.genny tmp.go
	go generate
	mv gen-tmp.go $@
	rm tmp.go

.PHONY: simulation plot test
.DEFAULT_GOAL: test
