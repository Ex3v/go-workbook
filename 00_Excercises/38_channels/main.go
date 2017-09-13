package main

import (
	"fmt"
	"math"
	"regexp"
)

const LETTER rune = 'i'
const THREADS int = 6

func main() {
	fmt.Println("\n\nCurrent letter to count is:", string(LETTER))
	fmt.Println("Thread count:", THREADS)

	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vestibulum diam nec lacus venenatis, sagittis efficitur sapien rutrum. Proin sit amet mi tortor. Aenean eget ligula a felis varius placerat. Aenean felis tortor, placerat eu auctor a, suscipit blandit nibh. Fusce tincidunt eget elit vitae aliquam. Vivamus lacinia quam leo, ac porta massa ornare id. Nulla facilisi. Curabitur eros sem, rutrum non nibh sed, aliquet tincidunt est. Ut suscipit augue quis diam laoreet, ac porta orci ullamcorper. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus."

	channels := spawnChannels()

	for i, channel := range channels {
		strChunk := getStrChunk(i, str)
		fmt.Printf("\nChunk %d: %s", i, strChunk)
		go channeledSum(strChunk, channel)
	}

	totalSum := getTotalSum(channels)

	fmt.Printf("\nThe total sum of letter '%s' in this string is: %d", string(LETTER), totalSum)
}
func getTotalSum(channels []chan int) int {
	sum := 0
	for _, c := range channels {
		partialSum := <-c
		sum += partialSum
	}

	return sum
}

func getStrChunk(offset int, s string) string {
	sLen := len(s)
	chunkSize := int(math.Ceil(float64(sLen / THREADS)))
	totalChunks := int(math.Floor(float64(sLen / chunkSize)))
	chunkStart := offset * chunkSize
	chunkEnd := chunkStart + chunkSize

	if chunkEnd > sLen || offset == (totalChunks-1) {
		chunkEnd = sLen
	}

	strChunk := s[chunkStart:chunkEnd]

	return strChunk
}

func spawnChannels() []chan int {
	var chans = make([]chan int, THREADS)

	for i := 0; i < THREADS; i++ {
		chans[i] = make(chan int)
	}

	return chans
}

func channeledSum(str string, c chan int) {
	sum := countLetterOccurrences(str)

	c <- sum
}

func countLetterOccurrences(str string) int {
	pattern := fmt.Sprintf("[%s]{1}", string(LETTER))

	r, _ := regexp.Compile(pattern)

	res := r.FindAllString(str, -1)

	return len(res)
}
