package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "server/grpc"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc"
)

func connectToDatabase() (*sql.DB, error) {
	// Configura tus parámetros de conexión aquí
	connString := ""
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *server) insertData(data Data) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO datainfo (Album, Year, Artist, Ranked) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, data.Album, data.Year, data.Artist, data.Ranked)

	if err != nil {
		return err
	}
	return nil
}

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetArtist())
	data := Data{
		Year:   in.GetYear(),
		Album:  in.GetAlbum(),
		Artist: in.GetArtist(),
		Ranked: in.GetRanked(),
	}
	fmt.Println(data)

	if err := s.insertData(data); err != nil {
		log.Printf("Error inserting data: %v", err)
		return nil, err
	}

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el álbum y lo inserté en la base de datos"}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
