package main

import (
	"time"
)

type InsertDataRequest struct {
	Temperature            *float64  `json:"temperature"`
	RelativeHumitiy        *float64  `json:"relative_humitiy"`
	Dewpoint               *float64  `json:"dewpoint"`
	Pressure               *float64  `json:"pressure"`
	WindDirection          *float64  `json:"wind_direction"`
	WindSpeed              *float64  `json:"wind_speed"`
	WindCorrectedDirection *float64  `json:"wind_corrected_direction"`
	TotalPrecipitation     *float64  `json:"total_precipitation"`
	PrecipitationIntensity *float64  `json:"precipitation_intensity"`
	GmxSupplyVoltage       *float64  `json:"gmx_supply_voltage"`
	GmxStatus              *int      `json:"gmx_status"`
	Timestamp              time.Time `json:"timestamp"`
}

type DataReqestDate struct {
	After  time.Time
	Before time.Time
}

type Data struct {
	Temperature            *float64  `json:"temperature"`
	RelativeHumitiy        *float64  `json:"relative_humitiy"`
	Dewpoint               *float64  `json:"dewpoint"`
	Pressure               *float64  `json:"pressure"`
	WindDirection          *float64  `json:"wind_direction"`
	WindSpeed              *float64  `json:"wind_speed"`
	WindCorrectedDirection *float64  `json:"wind_corrected_direction"`
	TotalPrecipitation     *float64  `json:"total_precipitation"`
	PrecipitationIntensity *float64  `json:"precipitation_intensity"`
	GmxSupplyVoltage       *float64  `json:"gmx_supply_voltage"`
	GmxStatus              *int      `json:"gmx_status"`
	Timestamp              time.Time `json:"timestamp"`
	CreatAt                time.Time `json:"creat_at"`
}

func NewData(req InsertDataRequest) (*Data, error) {
	return &Data{
		Temperature:            req.Temperature,
		RelativeHumitiy:        req.RelativeHumitiy,
		Dewpoint:               req.Dewpoint,
		Pressure:               req.Pressure,
		WindDirection:          req.WindDirection,
		WindSpeed:              req.WindSpeed,
		WindCorrectedDirection: req.WindCorrectedDirection,
		TotalPrecipitation:     req.TotalPrecipitation,
		PrecipitationIntensity: req.PrecipitationIntensity,
		GmxSupplyVoltage:       req.GmxSupplyVoltage,
		GmxStatus:              req.GmxStatus,
		Timestamp:              req.Timestamp,
		CreatAt:                time.Now().UTC(),
	}, nil
}
