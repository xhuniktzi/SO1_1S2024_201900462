package main

import (
	"context"
	"fmt"
	"log"

	pb "client/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
}

func insertData(c *fiber.Ctx) error {
	fmt.Print("Inserting data\n")
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	fmt.Print("Data: ", data, "\n")

	rank := Data{
		Album:  data["album"],
		Year:   data["year"],
		Artist: data["artist"],
		Ranked: data["ranked"],
	}

	fmt.Print("Rank: ", rank, "\n")

	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Artist: rank.Artist,
		Album:  rank.Album,
		Year:   rank.Year,
		Ranked: rank.Ranked,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetInfo())

	return nil
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
