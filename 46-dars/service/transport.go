package service

import (
	"context"
	"log"
	pb "mymode/protos/transport/protos"
	"mymode/storage/postgres"
	"time"
)

type TransportServer struct {
	pb.UnimplementedTransportServer
	TDb postgres.TransportRepo
}

func NewTransportRepo(t postgres.TransportRepo) *TransportServer {
	return &TransportServer{TDb: t}
}

func (T *TransportServer) GetBusSchedule(ctx context.Context, number *pb.Number)(*pb.Bus, error){
	bus, err := T.TDb.GetBusSchedule(int(number.Number))
	if err != nil{
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return bus, err
}

func (T *TransportServer) TrackBusLocation(ctx context.Context, number *pb.Number)(*pb.Location, error){
	loc, err := T.TDb.TrackBusLocation(int(number.Number))
	if err != nil{
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return loc, err
}

func (T *TransportServer) ReportTrafficJam(ctx context.Context, loc *pb.Location)(*pb.Status, error){
	status, err := T.TDb.ReportTrafficJam(*loc)
	if err != nil{
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return status, err
}
