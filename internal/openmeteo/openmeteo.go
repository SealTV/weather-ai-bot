package openmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://geocoding-api.open-meteo.com/v1"

type Client struct {
	hc *http.Client
}

func NewClient() *Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()

	transport.MaxIdleConns = 10

	return &Client{
		hc: &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		},
	}
}

type GeoResponse struct {
	Results []LatLong `json:"results"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *Client) GetLatLong(ctx context.Context, city string) (*LatLong, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/search?name=%s&count=1&language=en&format=json", baseURL, url.QueryEscape(city)), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to Geo API: %w", err)
	}
	defer resp.Body.Close()

	var response GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(response.Results) < 1 {
		return nil, errors.New("no results found")
	}

	return &response.Results[0], nil
}

type ForecastResponse struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	Elevation            float64     `json:"elevation"`
	GenerationTimeMS     float64     `json:"generationtime_ms"`
	UtcOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbr"`
	Hourly               []Hourly    `json:"hourly"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Daily                []Daily     `json:"daily"`
	DailyUnits           DailyUnits  `json:"daily_units"`
}

type Hourly struct {
	Time          []time.Time `json:"time,omitempty"`
	Temperature2m []float64   `json:"temperature_2m,omitempty"`
}

type HourlyUnits struct {
	Temperature2m string `json:"temperature_2m,omitempty"`
}

func (c *Client) GetForecast(ctx context.Context, location *LatLong) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/point?lat=%f&lon=%f", baseURL, location.Latitude, location.Longitude), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return fmt.Errorf("error making request to Geo API: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
