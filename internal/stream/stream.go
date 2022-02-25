package stream

import (
	"os"

	"github.com/keezeden/lance/pkg/utils"
)

type Stream struct {
	index int
	buffer []byte
}

func (s *Stream) Peek() byte {
	char := s.buffer[s.index]
	return char
}

func (s *Stream) Pop() {
	s.index++
}

func (s *Stream) Eof() bool {
	return s.index == len(s.buffer) 
}


func BuildStream(file string) Stream {
	content, err := os.ReadFile(file)
    utils.Check(err)

	return Stream{
		index: 0,
		buffer: content,
	}
}