package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	assert := assert.New(t)
	file := "./snippets/example.ll"
	var chars []byte

	content, err := os.ReadFile(file)
	check(err)

	streamer := stream(file)
	for !seof(&streamer) {
		char := speek(&streamer)

		chars = append(chars, char)

		spop(&streamer)
	}

	assert.Equal(content, chars, "File content is correct")
}

