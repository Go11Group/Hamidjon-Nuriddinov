package postgres

import (
	"database/sql"
	"mymode/model"
)

type StationRepo struct {
	Db *sql.DB
}


func NewStationRepo(db *sql.DB) *StationRepo{
	return &StationRepo{Db: db}
}


func (S *StationRepo) CreateStation(station model.CreateStation)error{
	_, err := S.Db.Exec(`INSERT INTO Stations(name) values($1)`, station.Name)
	return err
}


func (S *StationRepo) GetByIdStation(id string) (model.CreateStation, error){
	station := model.CreateStation{}
	err := S.Db.QueryRow(`SELECT * FROM Stations WHERE station_id = $1`, id).Scan(&station.StationId, &station.Name)
	return station, err
}


func (S *StationRepo) GetAllStations() ([]model.CreateStation, error){
	stations := []model.CreateStation{}
	rows, err := S.Db.Query(`SELECT * FROM Stations`)
	if err != nil{
		return stations, err
	}
	for rows.Next(){
		station := model.CreateStation{}
		err := rows.Scan(&station.StationId, &station.Name)
		if err != nil{
			return stations, err
		}
		stations = append(stations, station)
	}
	return stations, err
}


func (S *StationRepo) UpdateStation(id string, station model.CreateStation) error{
	_, err := S.Db.Exec(`UPDATE Stations SET name = $1 WHERE station_id = $2`, station.Name, station.StationId)
	return err
}


func (S *StationRepo) DeleteStation(id string) error{
	_, err := S.Db.Exec(`DELETE FROM Stations WHERE station_id = $1`, id)
	return err
}
