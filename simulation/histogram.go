package simulation

import (
	"fmt"
	"os"
	"path/filepath"
)

type buckets map[int]int

type histogram struct {
	name    string
	buckets buckets
	max     int
}

type histograms map[string]*histogram

func newHistograms(names []string) histograms {
	coll := make(histograms, len(names))
	for _, name := range names {
		coll[name] = newHistogram(name)
	}
	return coll
}

func newHistogram(name string) *histogram {
	return &histogram{
		name:    name,
		buckets: make(buckets),
	}
}

func (h *histogram) inc(value int) {
	if value > h.max {
		h.max = value
	}

	h.buckets[value] += 1
}

func (h *histogram) add(other *histogram) {
	if other.max > h.max {
		h.max = other.max
	}

	for k, count := range other.buckets {
		h.buckets[k] += count
	}
}

func (h *histogram) plot(filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("plot error: %v", err)
		return
	}

	defer f.Close()

	for value := 0; value <= h.max; value++ {
		count := h.buckets[value]
		if count == 0 {
			continue
		}
		f.WriteString(fmt.Sprintf("%d %d\n", value, count))
	}
}

func (hs histograms) add(other histograms) {
	for key, theirs := range other {
		ours, ok := hs[key]
		if !ok {
			hs[key] = theirs
			continue
		}

		ours.add(theirs)
	}
}

func (hs histograms) plot(directory string) {
	for key, h := range hs {
		h.plot(filepath.Join(directory, key))
	}
}
