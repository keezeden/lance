package main

import (
	"os"
)

type Stream struct {
	index int
	buffer []byte
}

func speek(s *Stream) byte {
	char := s.buffer[s.index]
	return char
}

func spop(s *Stream) {
	s.index++
}

func seof(s *Stream) bool {
	return s.index == len(s.buffer) 
}


func stream(file string) Stream {
	content, err := os.ReadFile(file)
    check(err)

	return Stream{
		index: 0,
		buffer: content,
	}
}