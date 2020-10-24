package main

import (
	"github.com/ovasquezbrito/candyshop/pkg/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/ovasquezbrito/candyshop/pkg/http/rest"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatal("error al configurar conexión con base de datos", err)
	}

	c, err := r.GetAllCandyNames()
	if err != nil {
		log.Fatal("error al consultar lista de dulces", err)
	}


	log.Println(c)
	fmt.Println("starting server on port 8080")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}

/* CREATE KEYSPACE candy_shop_db WITH replication = {'class' : 'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE IF NOT EXISTS candies (
	candy_id uuid,
	category text,
	name text,
	price float,
	PRIMARY KEY (candy_id, category)
);

INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Chocolate', 'Kikat', 1.99);
INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Toronto', 'Saboy', 1.99);
INSERT INTO candies (candy_id, category, name, price) VALUES (now(), 'Galleta Maria', 'Pui', 1.99);
.*/