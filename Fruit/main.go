package main

import (
	"log"
	"os"
	"time"

	"github.com/arijitnayak92/taskAfford/Fruit/appcontext"
	"github.com/arijitnayak92/taskAfford/Fruit/db"
	"github.com/arijitnayak92/taskAfford/Fruit/domain"
	"github.com/arijitnayak92/taskAfford/Fruit/handler"
	"github.com/arijitnayak92/taskAfford/Fruit/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Postgres driver
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
}

func main() {
	server := gin.Default()
	appContext := appcontext.NewAppContext(postgresURI)

	pg, err := db.NewPostgres(appContext)
	if err != nil {
		log.Println(err)
		return
	}
	defer pg.Close()
	mongouri := os.Getenv("MONGODB_URI")
	mongoClient, mongoError := db.MongoDBConection(mongouri)
	if mongoError != nil {
		log.Println(mongoError)
	}
	appDB := db.NewDB(pg, mongoClient)
	d := domain.NewDomain(appContext, appDB)
	h := handler.NewHandler(appContext, d)
	r := routes.NewRoutes(h)

	routes.AttachRoutes(server, r)

	log.Printf("the server is now running on port %s", port)
	go func() {
		errc <- server.Run(port)
	}()

	select {
	case err := <-errc:
		log.Printf("ListenAndServe error: %v", err)
	case <-done:
		log.Println("shutting down server ...")
	}
	time.AfterFunc(1*time.Second, func() {
		close(done)
		close(errc)
	})
}
