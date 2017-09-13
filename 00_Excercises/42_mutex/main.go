package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	CHUNK_SIZE int = 20
)

type WordStat struct {
	stats map[string]int
	mux   sync.Mutex
}

type Search func(a string) []string

func (ws *WordStat) String() string {
	var buffer bytes.Buffer

	for k, v := range ws.stats {
		str := fmt.Sprintf("Letter: %s, stat: %d\n", string(k), v)
		buffer.WriteString(str)
	}

	return buffer.String()
}

func (ws *WordStat) AddLetter(r string) {
	if val, ok := ws.stats[r]; ok {
		ws.stats[r] = val + 1
	} else {
		ws.stats[r] = 1
	}
}

func NewWordStat() *WordStat {
	stats := make(map[string]int)
	ws := WordStat{stats, sync.Mutex{}}

	return &ws
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("\n\nMutex example")

	ws := NewWordStat()
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vestibulum diam nec lacus venenatis, sagittis efficitur sapien rutrum. Proin sit amet mi tortor. Aenean eget ligula a felis varius placerat. Aenean felis tortor, placerat eu auctor a, suscipit blandit nibh. Fusce tincidunt eget elit vitae aliquam. Vivamus lacinia quam leo, ac porta massa ornare id. Nulla facilisi. Curabitur eros sem, rutrum non nibh sed, aliquet tincidunt est. Ut suscipit augue quis diam laoreet, ac porta orci ullamcorper. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus."
	chunks := getChunkedString(str)

	for id, chunk := range chunks {
		wg.Add(1)
		go Parser(id, chunk, ws)
	}

	wg.Wait()
	fmt.Println("\nDone! Stats are:\n\n")
	fmt.Print(ws)
}

func Parser(identifier int, s string, ws *WordStat) {
	splitted := strings.Split(s, "")

	for _, letter := range splitted {
		ws.mux.Lock()
		fmt.Printf("\nLocking for parser %d", identifier)
		time.Sleep(time.Duration((rand.Int() % 50) * int(time.Millisecond)))
		ws.AddLetter(letter)
		ws.mux.Unlock()
	}

	wg.Done()
}

func getChunkedString(s string) []string {
	out := []string{}
	maxChunks := int(math.Ceil(float64(len(s) / CHUNK_SIZE)))

	for currentChunk := 0; currentChunk < maxChunks; currentChunk++ {
		chunkStart := currentChunk * CHUNK_SIZE
		chunkEnd := (currentChunk + 1) * CHUNK_SIZE
		chunk := s[chunkStart:chunkEnd]
		out = append(out, chunk)
	}

	return out
}
