package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/valeno12/MiradaAnimal/config/database"
	"github.com/valeno12/MiradaAnimal/config/logger"
)

// @title Peluquería API
// @version 1.0
// @description API para la gestión de una peluquería
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Inicializar Logger
	logger.InitLogger()
	logger.Log.Info("Inicializando la aplicación...")

	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error al cargar el archivo .env: ", err)
	}

	// Inicializar el servidor en modo normal
	database.InitializeDatabase()
	logger.Log.Info("Base de datos inicializada correctamente")

	e := echo.New()
	routes.RegisterRoutes(e)
	logger.Log.Info("Rutas registradas correctamente")

	// Arrancar el servidor
	port := ":8080"
	logger.Log.Infof("Servidor iniciado en %s", port)
	if err := e.Start(port); err != nil {
		logger.Log.Fatal("Error al iniciar el servidor: ", err)
	}
}
