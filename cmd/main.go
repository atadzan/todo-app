package main

import(
  _ "github.com/lib/pq"
  "github.com/joho/godotenv"
  "github.com/spf13/viper"
  "github.com/atadzan/todo-app"
  "github.com/atadzan/todo-app/pkg/handler"
  "github.com/atadzan/todo-app/pkg/repository"
  "github.com/atadzan/todo-app/pkg/service"
  "log"
  "os"
)

func main(){
  if err := initConfig(); err != nil {
    log.Fatalf("error initializing configs: %s", err.Error())
  }

  if err := godotenv.Load(); err != nil {
    log.Fatalf("error loading env variables: %s", err.Error())
  }

  db, err := repository.NewPostgresDB(repository.Config{
    Host:     viper.GetString("db.host"),
    Port:     viper.GetString("db.port"),
    Username: viper.GetString("db.username"),
    DBName:   viper.GetString("db.dbname"),
    SSLMode:   viper.GetString("db.sslmode"),
    Password: os.Getenv("DB_PASSWORD"),
  })
  if err != nil {
    log.Fatalf("failed to initialize db: %s", err.Error())
  }

  repos := repository.NewRepository(db)
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  srv := new(todo.Server)
  if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
    log.Fatalf("error occured while running http server: #{err.Error()}")
  }
}
func initConfig() error{
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}
