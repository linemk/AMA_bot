package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetToken(t *testing.T) {
	// обнуляем флаги, которые могут быть в реальном запуске
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cmd", "--token=token"}

	token := GetToken()
	assert.Equal(t, "token", token, "Expected token to be test_token_value")
}
