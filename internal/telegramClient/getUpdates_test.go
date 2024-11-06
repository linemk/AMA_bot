package telegramClient

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPassedGetToken(t *testing.T) {
	// обнуляем флаги, которые могут быть в реальном запуске
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cmd", "--token=token"}

	token := getToken()
	assert.Equal(t, "token", token, "Expected token to be test_token_value")
}
