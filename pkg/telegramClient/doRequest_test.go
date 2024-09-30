package telegramClient

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"
)

func TestNewClient(t *testing.T) {
	host := "api.telegram.org"
	token := "1234"

	client := NewClient(host, token)

	require.NotNil(t, client, "Expected client to be non-nil")
	assert.Equal(t, host, client.host)
	assert.Equal(t, "bot"+token, client.baseUrl)
}

func TestUpdates(t *testing.T) {
	var token = "1234"
	var client = Client{
		host:    "api.telegram.org",
		baseUrl: "bot" + token,
		client:  &http.Client{},
	}
	t.Run("")
}

func TestDoRequest(t *testing.T) {
	var method = "getUpdates"
	var client = Client{}

	u := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   path.Join(client.baseUrl, method),
	}
	req, err := http.NewServer()
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
}
