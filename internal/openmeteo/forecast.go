package openmeteo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ForecastResponse struct {
	Latitude             float64       `json:"latitude"`
	Longitude            float64       `json:"longitude"`
	Elevation            float64       `json:"elevation"`
	GenerationTimeMS     float64       `json:"generationtime_ms"`
	UtcOffsetSeconds     int           `json:"utc_offset_seconds"`
	Timezone             string        `json:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbr"`
	Current              *Current      `json:"current,omitempty"`
	CurrentUnits         *CurrentUnits `json:"current_units,omitempty"`
	Hourly               *Hourly       `json:"hourly,omitempty"`
	HourlyUnits          *HourlyUnits  `json:"hourly_units,omitempty"`
	// Daily                *Daily        `json:"daily,omitempty"`
	// DailyUnits           *DailyUnits   `json:"daily_units,omitempty"`
}

type Current struct {
	Time                *time.Time `json:"time,omitempty"`
	Interval            *int       `json:"interval,omitempty"`
	Temperature2m       *float64   `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *float64   `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *float64   `json:"apparent_temperature,omitempty"`
	IsDay               *bool      `json:"is_day,omitempty"`
	Precipitation       *float64   `json:"precipitation,omitempty"`
	Rain                *float64   `json:"rain,omitempty"`
	Showers             *float64   `json:"showers,omitempty"`
	Showfall            *float64   `json:"snowfall,omitempty"`
	WeatherCode         *int       `json:"weathercode,omitempty"`
	CloudCover          *float64   `json:"cloud_cover,omitempty"`
	PressureMsl         *float64   `json:"pressure_msl,omitempty"`
	SurfacePressure     *float64   `json:"surface_pressure,omitempty"`
	WindSpeed10m        *float64   `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *float64   `json:"wind_direction_10m,omitempty"`
	WindGusts10m        *float64   `json:"wind_gusts_10m,omitempty"`
}

type CurrentUnits struct {
	Time                *string `json:"time,omitempty"`
	Interval            *string `json:"interval,omitempty"`
	Temperature2m       *string `json:"temperature_2m,omitempty"`
	RelativeHumidity2m  *string `json:"relative_humidity_2m,omitempty"`
	ApparentTemperature *string `json:"apparent_temperature,omitempty"`
	IsDay               *string `json:"is_day,omitempty"`
	Precipitation       *string `json:"precipitation,omitempty"`
	Rain                *string `json:"rain,omitempty"`
	Showers             *string `json:"showers,omitempty"`
	Showfall            *string `json:"snowfall,omitempty"`
	WeatherCode         *string `json:"weathercode,omitempty"`
	CloudCover          *string `json:"cloud_cover,omitempty"`
	PressureMsl         *string `json:"pressure_msl,omitempty"`
	SurfacePressure     *string `json:"surface_pressure,omitempty"`
	WindSpeed10m        *string `json:"wind_speed_10m,omitempty"`
	WindDirection10m    *string `json:"wind_direction_10m,omitempty"`
	WindGusts10m        *string `json:"wind_gusts_10m,omitempty"`
}

type Hourly struct {
	Time                     []time.Time `json:"time,omitempty"`
	Temperature2m            []float64   `json:"temperature_2m,omitempty"`
	RelativeHumidity2m       []float64   `json:"relative_humidity_2m,omitempty"`
	DevPoint2m               []float64   `json:"dew_point_2m,omitempty"`
	ApparentTemperature      []float64   `json:"apparent_temperature,omitempty"`
	PrecipitationProbability []float64   `json:"precipitation_probability,omitempty"`
	Precipitation            []float64   `json:"precipitation,omitempty"`
	Rain                     []float64   `json:"rain,omitempty"`
	Showers                  []float64   `json:"showers,omitempty"`
	Snowfall                 []float64   `json:"snowfall,omitempty"`
	SnowDepth                []float64   `json:"snow_depth,omitempty"`
	WeatherCode              []int       `json:"weathercode,omitempty"`
	PressureMsl              []float64   `json:"pressure_msl,omitempty"`
	SurfacePressure          []float64   `json:"surface_pressure,omitempty"`
	CloudCover               []float64   `json:"cloud_cover,omitempty"`
	CloudCoverLow            []float64   `json:"cloud_cover_low,omitempty"`
	CloudCoverMid            []float64   `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh           []float64   `json:"cloud_cover_high,omitempty"`
	Visibility               []float64   `json:"visibility,omitempty"`
	Evapotranspiration       []float64   `json:"evapotranspiration,omitempty"`
	Et0FaiEvapotranspiration []float64   `json:"et0_fai_evapotranspiration,omitempty"`
	VapourPressureDeficit    []float64   `json:"vapour_pressure_deficit,omitempty"`
	WindSpeed10m             []float64   `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             []float64   `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            []float64   `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            []float64   `json:"wind_speed_180m,omitempty"`
	WindDirection10m         []float64   `json:"wind_direction_10m,omitempty"`
	WindDirection80m         []float64   `json:"wind_direction_80m,omitempty"`
	WindDirection120m        []float64   `json:"wind_direction_120m,omitempty"`
	WindDirection180m        []float64   `json:"wind_direction_180m,omitempty"`
	WindGusts10m             []float64   `json:"wind_gusts_10m,omitempty"`
	Temerature80m            []float64   `json:"temperature_80m,omitempty"`
	Temerature120m           []float64   `json:"temperature_120m,omitempty"`
	Temperature180m          []float64   `json:"temperature_180m,omitempty"`
	SoilTemperature0cm       []float64   `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       []float64   `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      []float64   `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      []float64   `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture0to1cm       []float64   `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm       []float64   `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9cm       []float64   `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm      []float64   `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm     []float64   `json:"soil_moisture_27_to_81cm,omitempty"`
	UvIndex                  []float64   `json:"uv_index,omitempty"`
	UvIndexClearSky          []float64   `json:"uv_index_clear_sky,omitempty"`
}

type HourlyUnits struct {
	Time                     *string `json:"time,omitempty"`
	Temperature2m            *string `json:"temperature_2m,omitempty"`
	RelevantHumidity2m       *string `json:"relative_humidity_2m,omitempty"`
	DevPoint2m               *string `json:"dew_point_2m,omitempty"`
	ApparentTemperature      *string `json:"apparent_temperature,omitempty"`
	PrecipitationProbability *string `json:"precipitation_probability,omitempty"`
	Precipitation            *string `json:"precipitation,omitempty"`
	Rain                     *string `json:"rain,omitempty"`
	Showers                  *string `json:"showers,omitempty"`
	Snowfall                 *string `json:"snowfall,omitempty"`
	SnowDepth                *string `json:"snow_depth,omitempty"`
	WeatherCode              *string `json:"weathercode,omitempty"`
	PressureMsl              *string `json:"pressure_msl,omitempty"`
	CloudCover               *string `json:"cloud_cover,omitempty"`
	CloudCoverLow            *string `json:"cloud_cover_low,omitempty"`
	CloudCoverMid            *string `json:"cloud_cover_mid,omitempty"`
	CloudCoverHigh           *string `json:"cloud_cover_high,omitempty"`
	Visibility               *string `json:"visibility,omitempty"`
	Evapotranspiration       *string `json:"evapotranspiration,omitempty"`
	Et0FaiEvapotranspiration *string `json:"et0_fai_evapotranspiration,omitempty"`
	VapourPressureDeficit    *string `json:"vapour_pressure_deficit,omitempty"`
	WindSpeed10m             *string `json:"wind_speed_10m,omitempty"`
	WindSpeed80m             *string `json:"wind_speed_80m,omitempty"`
	WindSpeed120m            *string `json:"wind_speed_120m,omitempty"`
	WindSpeed180m            *string `json:"wind_speed_180m,omitempty"`
	WindDirection10m         *string `json:"wind_direction_10m,omitempty"`
	WindDirection80m         *string `json:"wind_direction_80m,omitempty"`
	WindDirection120m        *string `json:"wind_direction_120m,omitempty"`
	WindDirection180m        *string `json:"wind_direction_180m,omitempty"`
	WindGusts10m             *string `json:"wind_gusts_10m,omitempty"`
	Temerature80m            *string `json:"temperature_80m,omitempty"`
	Temerature120m           *string `json:"temperature_120m,omitempty"`
	Temperature180m          *string `json:"temperature_180m,omitempty"`
	SoilTemperature0cm       *string `json:"soil_temperature_0cm,omitempty"`
	SoilTemperature6cm       *string `json:"soil_temperature_6cm,omitempty"`
	SoilTemperature18cm      *string `json:"soil_temperature_18cm,omitempty"`
	SoilTemperature54cm      *string `json:"soil_temperature_54cm,omitempty"`
	SoilMoisture0to1cm       *string `json:"soil_moisture_0_to_1cm,omitempty"`
	SoilMoisture1to3cm       *string `json:"soil_moisture_1_to_3cm,omitempty"`
	SoilMoisture3to9cm       *string `json:"soil_moisture_3_to_9cm,omitempty"`
	SoilMoisture9to27cm      *string `json:"soil_moisture_9_to_27cm,omitempty"`
	SoilMoisture27to81cm     *string `json:"soil_moisture_27_to_81cm,omitempty"`
	UvIndex                  *string `json:"uv_index,omitempty"`
	UvIndexClearSky          *string `json:"uv_index_clear_sky,omitempty"`
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
