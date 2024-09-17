package container

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"github.com/Lazaro-Barros/boilerplate-golang/command/application"
	"github.com/Lazaro-Barros/boilerplate-golang/command/domain/entities"
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driven/postgres"
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http"
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http/handler"
	"github.com/Lazaro-Barros/boilerplate-golang/command/infra/driver/http/router"
	"github.com/Lazaro-Barros/boilerplate-golang/config"
	db_sqlc "github.com/Lazaro-Barros/boilerplate-golang/queries_sqlc/db"
)

var container sync.Map

func Init() {
	container = sync.Map{}
}

func getService(alias string, f func() interface{}) interface{} {
	/* Reuse instance */
	if val, ok := container.Load(alias); ok {
		return val
	}

	service := f()
	/* Save instance in container */
	container.Store(alias, service)
	return service
}

func GetConfig() config.Config {
	return getService("config", func() interface{} {
		return config.Get()
	}).(config.Config)
}

func GetPostgresConnection() *sql.DB {
	return getService("postgres", func() interface{} {
		cfg := GetConfig()
		// Configuração da string de conexão
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Postgres.Username,
			cfg.Postgres.Password,
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.Database,
		)

		// Abre a conexão com o banco de dados
		dbConn, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("cannot open database connection:", err)
		}

		// Configurando o pool de conexões
		dbConn.SetMaxOpenConns(25)                 // Máximo de 25 conexões abertas
		dbConn.SetMaxIdleConns(25)                 // Máximo de 25 conexões ociosas
		dbConn.SetConnMaxLifetime(5 * time.Minute) // Conexões duram no máximo 5 minutos

		// Testa a conexão com o banco de dados
		if err := dbConn.Ping(); err != nil {
			log.Fatal("cannot connect to db:", err)
		}
		return dbConn
	}).(*sql.DB)
}

func GetSqlcDb() *db_sqlc.Queries {
	return getService("sqlc", func() interface{} {
		dbConn := GetPostgresConnection()
		return db_sqlc.New(dbConn)
	}).(*db_sqlc.Queries)
}

func GetTodoRepository() entities.TodoRepository {
	return getService("todoRepository", func() interface{} {
		db := GetSqlcDb()
		return postgres.NewTodoRepository(db)
	}).(entities.TodoRepository)
}

func GetTodoService() *application.TodoService {
	return getService("todoService", func() interface{} {
		repo := GetTodoRepository()
		return application.NewTodoService(repo)
	}).(*application.TodoService)
}

func GetTodoHandler() *handler.TodoHandler {
	return getService("todoHandler", func() interface{} {
		service := GetTodoService()
		return handler.NewTodoHandler(service)
	}).(*handler.TodoHandler)
}

func GetRouter() *http.Router {
	return getService("router", func() interface{} {
		Todohandler := GetTodoHandler()
		return router.GetRouter(Todohandler)
	}).(*http.Router)
}
