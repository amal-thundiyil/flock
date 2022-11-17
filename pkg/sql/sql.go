package sql

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func connectDB() {
  err := godotenv.Load(".env") 
  if err != nil {
    log.Fatalf("Error loading env file")
  }
  config, err := pgx.ParseConfig(os.Getenv("COCKROACHDB_URL"))
  if err != nil {
    log.Fatalf("Error connecting to DB")
  }
  config.RuntimeParams["application_name"] = "$ server_flock_db"
  conn, err := pgx.ConnectConfig(context.Background(), config)
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close(context.Background())
}
