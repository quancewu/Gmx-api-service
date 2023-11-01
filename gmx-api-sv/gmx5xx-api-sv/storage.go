package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetDatas(*DataReqestDate) ([]*Data, error)
	GetLatestData() (*Data, error)
	InsertData(*Data) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	serverName := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	if serverName == "" {
		serverName = "localhost"
	}
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", serverName, dbName, dbName, dbPassword)
	log.Printf("Postgres SQL hostname: %s", serverName)
	db, err := sql.Open("postgres", connStr)
	// defer db.Close()
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createDataTable()
}

// 4ch   string 40001-02 Node
// 32bit float 40003-04 Direction
// 32bit float 40005-06 Speed
// 32bit float 40007-08 Corrected Direction
// 32bit float 40009-10 Pressure
// 32bit float 40011-12 Relative Humidity
// 32bit float 40013-14 Temperature
// 32bit float 40015-16 Dewpoint
// 32bit float 40017-18 Total Precipitation
// 32bit float 40019-20 Precipitation Intensity
// 16ch  string 40021-28 Date
// 16ch  string 40029-36 Time
// 32bit float 40037-38 Supply Voltage
// 32bit UINT 40039-40 Status

func (s *PostgresStore) createDataTable() error {
	query := `create table if not exists gmx5xx (
		id serial primary key,
		temperature float,
		relative_humidity float,
		dewpoint float,
		pressure float,
		wind_direction float,
		wind_speed float,
		wind_corrected_direction float,
		total_precipitation float,
		precipitation_intensity float,
		gmx_supply_voltage float,
		gmx_status integer,
		timestamp timestamp with time zone,
		created_at timestamp with time zone,
		CONSTRAINT gmx5xx_timestamp_key UNIQUE ("timestamp")
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) InsertData(data *Data) error {
	query := `insert into gmx5xx
	(temperature, relative_humidity, dewpoint, pressure, wind_direction, wind_speed, wind_corrected_direction,
	total_precipitation, precipitation_intensity, gmx_supply_voltage, gmx_status, timestamp, created_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	rows, err := s.db.Query(
		query,
		data.Temperature,
		data.RelativeHumitiy,
		data.Dewpoint,
		data.Pressure,
		data.WindDirection,
		data.WindSpeed,
		data.WindCorrectedDirection,
		data.TotalPrecipitation,
		data.PrecipitationIntensity,
		data.GmxSupplyVoltage,
		data.GmxStatus,
		data.Timestamp,
		data.CreatAt)

	defer rows.Close()

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetDatas(date *DataReqestDate) ([]*Data, error) {
	rows, err := s.db.Query(`select temperature, relative_humidity, dewpoint, pressure, wind_direction, wind_speed,
	wind_corrected_direction,	total_precipitation, precipitation_intensity, gmx_supply_voltage, gmx_status,
	timestamp, created_at from gmx5xx where timestamp between $1 and $2 order by timestamp`, date.After, date.Before)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	datas := []*Data{}
	for rows.Next() {
		data, err := scanIntoData(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}

func (s *PostgresStore) GetLatestData() (*Data, error) {
	rows, err := s.db.Query(`select temperature, relative_humidity, dewpoint, pressure, wind_direction, wind_speed,
	wind_corrected_direction,	total_precipitation, precipitation_intensity, gmx_supply_voltage, gmx_status,
	timestamp, created_at from gmx5xx ORDER BY timestamp DESC LIMIT 1`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	// datas := []*Data{}
	for rows.Next() {
		data, err := scanIntoData(rows)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, fmt.Errorf("GMX Get Last data not found %s", err)
}

func scanIntoData(rows *sql.Rows) (*Data, error) {
	data := new(Data)
	err := rows.Scan(
		&data.Temperature,
		&data.RelativeHumitiy,
		&data.Dewpoint,
		&data.Pressure,
		&data.WindDirection,
		&data.WindSpeed,
		&data.WindCorrectedDirection,
		&data.TotalPrecipitation,
		&data.PrecipitationIntensity,
		&data.GmxSupplyVoltage,
		&data.GmxStatus,
		&data.Timestamp,
		&data.CreatAt)

	return data, err
}
