package simulation

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"testing"

	h "github.com/hashicorp/hyparview"
)

// TestSimulation is the test entry point
func TestSimulation(t *testing.T) {
	count := envInt("SIMULATION_COUNT", 5)
	config := WorldConfig{
		peers:          envInt("SIMULATION_PEERS", 1000),
		failureRate:    envInt("SIMULATION_NETWORK_FAILURE_PERCENT", 5),
		mortality:      envInt("SIMULATION_NODE_FAILURE_PERCENT", 5),
		gossipMessages: envInt("SIMULATION_MESSAGES", 200),
	}
	config.stabilizationRounds = config.peers / 10

	totalStats := newStats()

	for i := 1; i <= count; i++ {
		config.seed = envInt64("SIMULATION_SEED", h.Rint64Crypto(math.MaxInt64-1))
		config.iteration = i
		w := testSimulation(t, config)
		totalStats.add(w.stats)
	}

	totalStats.plot("../data")
}

func envInt(variable string, otherwise int) int {
	return int(envInt64(variable, int64(otherwise)))
}

func envInt64(variable string, otherwise int64) int64 {
	conv, err := strconv.ParseInt(os.Getenv(variable), 10, 64)
	if err != nil {
		conv = otherwise
	}
	return conv
}

// testSimulation is the entry point to test a single world
// World configuration and assertion goes here
func testSimulation(t *testing.T, config WorldConfig) *World {
	fmt.Printf("world: %d seed: %d peers: %d\n", config.iteration, config.seed, config.peers)

	w := simulation(config)
	w.stats = newStats()

	err := w.Connected()
	if err != nil {
		t.Errorf("world %d: graph disconnected: %s", config.iteration, err.Error())
	}

	// This isn't an error. It's useful for working on symmetry, but because of the
	// failure rate, there's always a tail of asymmetries
	// err = w.isSymmetric()
	// if err != nil {
	// 	t.Logf("run %d: active view asymmetric: %s", i, err.Error())
	// }

	// w.debugQueue()
	// w.plotPeer("n2375")
	w.mkdir()
	w.plotSeed(config.seed)

	w.plotBootstrapCount()
	w.plotInDegree()
	w.plotOutDegree()
	w.plotGossip()
	w.plotGraphs()
	w.stats.plot(w.dir())

	return w
}

func newStats() histograms {
	return newHistograms([]string{
		"bootstrap",
		"in_degree_active",
		"in_degree_passive",
		"out_degree_active",
		"out_degree_passive",
		// "hyparview_messages",
		// "gossip_seen",
		// "gossip_waste",
	})
}
