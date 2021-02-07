package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/appcontext"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/db"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/domain"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/handler"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/routes"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/validation"

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
	mongoURI = os.Getenv("MONGO_URI")
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

	u := utils.NewUtil()
	v := validation.NewValidation(u)

	d := domain.NewDomain(appContext, pg, mongoClient, u)
	h := handler.NewHandler(appContext, d, v, u)
	r := routes.NewRouter(h)
	router, _ := r.Routes()

	// log.Printf("Server Running on port %s", port)

	router.Run(port)

}
