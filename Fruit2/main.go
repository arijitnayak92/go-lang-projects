package main

import (
	"log"

	"github.com/arijitnayak92/taskAfford/Fruit2/db"
	"github.com/arijitnayak92/taskAfford/Fruit2/domain"
	"github.com/arijitnayak92/taskAfford/Fruit2/handler"
	"github.com/arijitnayak92/taskAfford/Fruit2/routes"
	"github.com/joho/godotenv"
)

func main() {
	myEnv, enverr := godotenv.Read() // Read Environment Variables and store it in a map
	if enverr != nil {
		log.Println("appcontext: error reading environment vars ", enverr)

	}

	postgresURI := myEnv["POSTGRES_URI"]
	postgresClient, err := db.ConnectToPostgres(postgresURI)
	if err != nil {
		log.Println("appcontext: error connecting to postgres - ", err)

	}

	mongoURI := myEnv["MONGO_URI"]
	mongoClient, err := db.ConnectToMongo(mongoURI)
	if err != nil {
		log.Println("appcontext: error connecting to mongodb - ", err)

	}

	// postgresRepo := db.NewPostgresRepo(postgresClient)
	// mongoRepo := db.NewMongoRepo(mongoClient) //Mongo repo
	appRepository := db.NewRepository(postgresClient, mongoClient)

	userDomain := domain.NewUser(appRepository.Postgres)
	// productDomain := domain.NewProduct(postgresRepo)
	// cartDomain := domain.NewCart(mongoRepo)

	//appDomain := domain.NewAppDomain(postgresRepo, postgresRepo, mongoRepo)
	appDomain := domain.NewDomain(appRepository)

	userHandler := handler.NewUser(userDomain)
	// productHandler := handler.NewProduct(productDomain)
	// cartHandler := handler.NewCart(cartDomain)

	appHandler := handler.NewHandler(appDomain)

	appRouter := routes.NewRouter()
	appRouter.SetupRoutes(appHandler, userHandler)
	appRouter.Router.Run(":8080")

}
