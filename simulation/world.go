package simulation

type World struct {
	config        *WorldConfig
	nodes         map[string]*Client
	morgue        map[string]*Client
	totalMessages int
	totalPayloads int

	gossipTotal *gossipRound
	gossipRound []*gossipRound

	spinCount  int
	spinCountM map[string]int
}

type WorldConfig struct {
	rounds      int
	peers       int
	mortality   int
	payloads    int
	iteration   int // count rounds for plot filenames
	shuffleFreq int
	failureRate int
}

func (w *World) get(id string) *Client {
	return w.nodes[id]
}

func (w *World) nodeKeys() []string {
	m := w.nodes
	ks := make([]string, len(m))
	i := 0
	for k, _ := range m {
		ks[i] = k
		i++
	}
	return ks
}

func (w *World) randNodes() (ns []*Client) {
	for _, k := range w.nodeKeys() {
		ns = append(ns, w.get(k))
	}
	return ns
}

// TODO: maybe accept the message we're deciding for and do different things?
func (w *World) shouldFail() bool {
	return false
}
