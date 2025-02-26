package main

import (
	"golang-voucher-api/pkg/database"
	"golang-voucher-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi Database
	database.Connect()

	// Inisialisasi Router
	router := gin.Default()

	// Registrasi Routes
	routes.RegisterRoutes(router)

	// Jalankan Server
	router.Run(":8080")
}
