package stream

import (
	"os"
	"testing"

	"github.com/keezeden/lance/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	assert := assert.New(t)
	file := "../../snippets/example.ll"
	var chars []byte

	content, err := os.ReadFile(file)
	utils.Check(err)

	streamer := BuildStream(file)
	for !streamer.Eof() {
		char := streamer.Peek()

		chars = append(chars, char)

		streamer.Pop()
	}

	assert.Equal(content, chars, "File content is correct")
}

