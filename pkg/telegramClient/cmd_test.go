package telegramClient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartParser(t *testing.T) {
	var ClientTest = NewClient("23", "23")

	t.Run("If id is empty", func(t *testing.T) {
		expectedError := "id or name is empty"
		err := StartParser(ClientTest, 0, "")

		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err.Error())
		}
	})
}
