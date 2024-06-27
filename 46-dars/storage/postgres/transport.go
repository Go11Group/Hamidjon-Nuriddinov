package postgres

import (
	"database/sql"
	pb "mymode/protos/transport/protos"
)


type TransportRepo struct{
	Db *sql.DB
}

func NewTransportRepo(db *sql.DB)*TransportRepo{
	return &TransportRepo{Db: db}
}

func(T *TransportRepo) GetBusSchedule(number int)(*pb.Bus, error){
	bus := pb.Bus{}
	err := T.Db.QueryRow(`SELECT * FROM Bus WHERE number = $1`, number).Scan(
		&bus.Number, &bus.From, &bus.To, &bus.Loc, &bus.TrafficStat)
	return &bus, err
}

func(T *TransportRepo) TrackBusLocation(number int)(*pb.Location, error){
	loc := pb.Location{}
	err := T.Db.QueryRow(`SELECT location FROM Bus Where number = $1`, number).Scan(
		&loc)
	return &loc, err
}

func(T *TransportRepo) ReportTrafficJam(loc pb.Location)(*pb.Status, error){
	_, err := T.Db.Exec(`UPDATE SET Bus traffic_stat = true WHERE location = $1`, loc)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}