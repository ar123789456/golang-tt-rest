package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest/auth"
	"rest/auth/repository/sqlite"
	"rest/counter"
	"rest/counter/repository/redis"
	"rest/email"
	"rest/hash"
	"rest/substr"
	"time"

	//usecase
	useruc "rest/auth/usecase"
	counteruc "rest/counter/usecase"
	emailuc "rest/email/usecase"
	substruc "rest/substr/usecase"

	//handler
	userhandler "rest/auth/delivery/http"
	counterhandler "rest/counter/delivery/http"
	emailhandler "rest/email/delivery/http"
	subs "rest/substr/delivery/http"

	"github.com/gin-gonic/gin"
	r "github.com/go-redis/redis"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	httpServer *http.Server

	//usecase
	userUc    auth.UseCase
	counterUC counter.ConterUseCase
	emailUC   email.EmailUseCase
	hashUC    hash.HashUseCase
	substrUC  substr.UseCase
}

func NewApp() *App {
	db, client := initDB()
	userrepo := sqlite.NewUserRepository(db)
	counterrepo := redis.NewCounterRepository(client)
	return &App{
		userUc:    useruc.NewAuthUseCase(userrepo),
		counterUC: counteruc.NewCounterUseCase(counterrepo),
		emailUC:   emailuc.NewEmailUseCase(),
		substrUC:  substruc.NewSubstrUseCase(),
	}
}

func (a *App) Run(port string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use(
	// 	gin.Recovery(),
	// 	gin.Logger(),
	// )
	userhandler.RegisterUserEndpoints(router, a.userUc)
	counterhandler.RegisterCounterEndpoints(router, a.counterUC)
	emailhandler.RegisterEmailEndpoints(router, a.emailUC)
	router.POST("/rest/substr/find", subs.NewHandler(a.substrUC).Post)

	//add Http server
	a.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() (*sql.DB, *r.Client) {

	client := r.NewClient(&r.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	db.Exec(`CREATE TABLE user(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		first_name TEXT,
		last_name TEXT
	  );`)
	return db, client
}
