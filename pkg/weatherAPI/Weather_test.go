package weather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockTranslator — заглушка для перевода
type MockTranslator struct{}

// RuToEng — реализуем метод перевода для теста
func (m MockTranslator) RuToEng(city string) string {
	return "Moscow" // всегда возвращает "Moscow"
}

// Мок-ответ от API для тестов
func mockWeatherAPI(w http.ResponseWriter, r *http.Request) {
	weatherResponse := WeatherResponse{
		Location: struct {
			Name string `json:"name"`
		}{
			Name: "Moscow",
		},
		Current: struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
			WindKph  float64 `json:"wind_kph"`
			Humidity int     `json:"humidity"`
		}{
			TempC: 15.0,
			Condition: struct {
				Text string `json:"text"`
			}{
				Text: "Облачно",
			},
			WindKph:  10.0,
			Humidity: 60,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResponse)
}

func TestGetWeather(t *testing.T) {
	// Создаем тестовый HTTP сервер
	server := httptest.NewServer(http.HandlerFunc(mockWeatherAPI))
	defer server.Close()

	// Мокаем URL API для теста
	mockAPIURL := server.URL + "?key=%s&q=%s&lang=ru"
	apiKey := "testKey"

	// Тестируем
	tests := []struct {
		city             string
		expectedCity     string
		expectedTemp     int
		expectedPrecip   string
		expectedWind     float64
		expectedHumidity int
	}{
		{
			city:             "Москва",
			expectedCity:     "Moscow",
			expectedTemp:     15,
			expectedPrecip:   "Облачно",
			expectedWind:     10.0,
			expectedHumidity: 60,
		},
	}

	// Создаем мок-транслятор
	translator := MockTranslator{}

	for _, tt := range tests {
		weather := GetWeather(tt.city, mockAPIURL, apiKey, translator)

		if weather.City != tt.expectedCity {
			t.Errorf("expected city %s, got %s", tt.expectedCity, weather.City)
		}
		if weather.Temperature != tt.expectedTemp {
			t.Errorf("expected temperature %d, got %d", tt.expectedTemp, weather.Temperature)
		}
		if weather.Precipitation != tt.expectedPrecip {
			t.Errorf("expected precipitation %s, got %s", tt.expectedPrecip, weather.Precipitation)
		}
		if weather.Wind != tt.expectedWind {
			t.Errorf("expected wind %.1f, got %.1f", tt.expectedWind, weather.Wind)
		}
		if weather.Humidity != tt.expectedHumidity {
			t.Errorf("expected humidity %d, got %d", tt.expectedHumidity, weather.Humidity)
		}
	}
}
