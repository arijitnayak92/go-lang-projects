package main

import (
	"flag"
	"log"
	"os"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/arijitnayak92/taskAfford/Fruit/routes"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/arijitnayak92/taskAfford/Fruit/validation"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	port        string
	postgresURI string
	mongoURI    string
)

func init() {
	godotenv.Load()

	flag.StringVar(&port, "port", ":8080", "server port")

	postgresURI = os.Getenv("POSTGRES_URI")
	if postgresURI == "" {
		log.Fatal("Postgres URI not found!")
	}

	mongoURI = os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("Mongo URI not found!")
	}
}

func main() {

	appContext := appcontext.NewAppContext(postgresURI, mongoURI)

	pg, err := db.NewPostgres(appContext)
	if err != nil {
		log.Println(err)
		return
	}
	defer pg.Close()

	mongoClient, mongoError := db.NewMongo(appContext)
	if mongoError != nil {
		log.Println(mongoError)
	}
	appDB := db.NewDB(pg, mongoClient)
	u := utils.NewUtil()
	v := validation.NewValidation(u)
	d := domain.NewDomain(appContext, appDB, u)
	h := handler.NewHandler(appContext, d, v, u)
	r := routes.NewRouter(h)
	router, _ := r.Routes()

	log.Printf("Server Running on port %s", port)

	router.Run(port)

}
