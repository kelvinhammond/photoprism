package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMimeType(t *testing.T) {
	t.Run("jpg", func(t *testing.T) {
		filename := ExpandFilename("./testdata/test.jpg")
		mimeType := MimeType(filename)
		assert.Equal(t, "image/jpeg", mimeType)
	})
	t.Run("not existing filename", func(t *testing.T) {
		filename := ExpandFilename("./testdata/xxx.jpg")
		mimeType := MimeType(filename)
		assert.Equal(t, "", mimeType)
	})
}
