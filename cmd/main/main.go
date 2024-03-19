package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pidanou/lifetrack/internal/api"
	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/types"
)

func main() {

	datastore := datastore.NewPostgresDatastore()

	datastore.Db.AutoMigrate(&types.Hydration{}, &types.Sleep{})

	port := flag.String("port", ":3000", "Port to serve")
	flag.Parse()

	fmt.Println("Server running on port", *port)
	server := api.NewServer(datastore, *port)
	log.Fatal(server.Start())

}
