package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL para `database/sql`
)

var DB *sql.DB

func InitializeDatabase() error {
	// Construir el DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Intentos de conexión
	maxRetries := 5
	retryDelay := 10 * time.Second

	var db *sql.DB
	var err error

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Intentando conectar a la base de datos (Intento %d/%d)...\n", i, maxRetries)
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			// Probar la conexión
			err = db.Ping()
			if err == nil {
				fmt.Println("Conexión a la base de datos exitosa")
				DB = db
				return nil
			}
		}

		fmt.Printf("Error al conectar con la base de datos: %v\n", err)
		if i < maxRetries {
			fmt.Printf("Reintentando en %v...\n", retryDelay)
			time.Sleep(retryDelay)
		} else {
			fmt.Println("No se pudo conectar a la base de datos después de varios intentos")
			return err
		}
	}

	return fmt.Errorf("error inesperado al intentar conectar a la base de datos")
}
