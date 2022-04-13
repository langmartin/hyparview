package simulation

import "math/rand"

func simulation(c WorldConfig) *World {
	w := &World{
		config: &c,
		nodes:  make(map[string]*Client, c.peers),
		morgue: make(map[string]*Client),
	}

	// log.Printf("debug: make all the nodes")
	for i := 0; i < c.peers; i++ {
		id := makeID(i)
		w.nodes[id] = makeClient(w, id)
	}

	// log.Printf("debug: connect all the nodes")
	ns := w.randNodes()
	w.bootstrap = ns[0].Self
	for _, me := range ns[1:] {
		// boot := w.nodes[fmt.Sprintf("n%d", h.Rint(i))]
		me.SendJoin(w.bootstrap)
	}

	for i := 0; i < c.gossipMessages; i++ {
		for _, c := range w.randNodes() {
			c.next()
		}

		// gossip drains all the hyparview messages and sends all the gossip
		// messages before returning. Also maintains the active view
		node := w.get(makeID(rand.Intn(len(w.nodes))))
		p := i + 1
		node.sendGossip(p)
		w.traceRound(p)
	}

	return w
}
