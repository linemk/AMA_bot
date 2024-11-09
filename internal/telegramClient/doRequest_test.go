package telegramClient

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

type MockClient struct {
	mock.Mock
}

// Реализация метода Updates для мока
func (m *MockClient) Updates(offset, limit int) ([]Update, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]Update), args.Error(1)
}

func (m *MockClient) SendMessage(chatID int, text string) error {
	args := m.Called(chatID, text)
	return args.Error(0)
}
