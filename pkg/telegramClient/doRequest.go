package telegramClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host    string       // host API tg -> api.telegram.org/
	baseUrl string       // "bot" + <token>
	client  *http.Client // для создания запросов (do)
}

// получает номер хоста + токен и формирует запрос
func NewClient(host string, token string) *Client {
	return &Client{
		host:    host,           //хост
		baseUrl: "bot" + token,  //path
		client:  &http.Client{}, //помогает осуществить запрос
	}
}

// в таком виде приходит ответ от тг
type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// единичное сообщение
type Update struct {
	Id      int     `json:"update_id"` //Id сообщения
	Message Message `json:"message"`
}

// структура сообщения
type Message struct {
	Text string `json:"text"` // текст
	Chat Chat   `json:"chat"` // id чата
	User User   `json:"from"`
}

type User struct {
	FirstName string `json:"first_name"`
}

type Chat struct {
	Id int `json:"id"` // Unique identifier for the chat
}

// метод для получения сообщений
func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset)) // номер сообщения
	q.Add("limit", strconv.Itoa(limit))   // лимит по отдаче

	// производим запрос при помощи метода "getUpdates"
	data, err := c.doRequest("getUpdates", q) // = GET
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse // слайс ответов
	// подставляем в json значения
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// сборка и отправка запроса
func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https", // протокол
		Host:   c.host,
		Path:   path.Join(c.baseUrl, method),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %v", err)
	}

	// добавляем offset + limit/ id + text
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %v", err)
	}
	return body, nil
}
