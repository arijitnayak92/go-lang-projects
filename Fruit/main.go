package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/affordmed/affmed/appcontext"
	"gitlab.com/affordmed/affmed/db"
	"gitlab.com/affordmed/affmed/domain"
	"gitlab.com/affordmed/affmed/handler"
	"gitlab.com/affordmed/affmed/routes"
	"gitlab.com/affordmed/affmed/util"
	"gitlab.com/affordmed/affmed/validation"
	"log"
	"os"
	"os/signal"
	"time"

	// Postgres driver
	_ "github.com/lib/pq"
)

var (
	// Derived from ldflags -X
	buildRevision string
	buildVersion  string
	buildTime     string

	// general options
	versionFlag bool
	helpFlag    bool

	// server port
	port string

	// program controller
	done = make(chan struct{})
	errc = make(chan error)

	postgresURI string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	flag.BoolVar(&versionFlag, "version", false, "show current version and exit")
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.StringVar(&port, "port", ":8080", "server port")

	postgresURI = os.Getenv("POSTGRES_URI")
	if postgresURI == "" {
		log.Fatal("Postgres URI not found!")
	}
}

func setBuildVariables() {
	if buildRevision == "" {
		buildRevision = "dev"
	}
	if buildVersion == "" {
		buildVersion = "dev"
	}
	if buildTime == "" {
		buildTime = time.Now().UTC().Format(time.RFC3339)
	}
}

func parseFlags() {
	flag.Parse()

	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if versionFlag {
		fmt.Printf("%s %s %s\n", buildRevision, buildVersion, buildTime)
		os.Exit(0)
	}
}

func handleInterrupts() {
	log.Println("start handle interrupts")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	sig := <-interrupt
	log.Printf("caught sig: %v", sig)
	// close resource here
	done <- struct{}{}
}

func main() {
	setBuildVariables()
	parseFlags()
	go handleInterrupts()

	server := gin.Default()

	appContext := appcontext.NewAppContext(postgresURI)

	pg, err := db.NewPostgres(appContext)
	if err != nil {
		log.Println(err)
		return
	}
	defer pg.Close()

	u := util.NewUtil()
	v := validation.NewValidation(u)
	d := domain.NewDomain(appContext, pg, u)
	h := handler.NewHandler(appContext, d, v, u)
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
