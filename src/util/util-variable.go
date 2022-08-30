package util

import "go.mongodb.org/mongo-driver/mongo"

// Variable block
var Client *mongo.Client
var CollectionName = [4]string{"users", "merchants", "outlets", "transactions"}

// Constanta block
const DatabaseName string = "majoo"
const Port string = "8080"

// Error block
const Error0 = "Error 0. Mohon cek kembali request anda"             // middleware-read.go
const Error1 = "Error 1. Terjadi kesalahan pada server"              // middleware-read.go
const Error2 = "Error 2. Mohon cek kembali login anda"               // middleware-read.go
const Error3 = "Error 3. Mohon cek kembali token anda"               // middleware-read.go
const Error4 = "Error 4. Terjadi kesalahan pada server"              // middleware-read.go
const Error5 = "Error 5. Limit tidak boleh kosong"                   // middleware-read.go
const Error6 = "Error 6. Pastikan mengisi limit dengan angka"        // middleware-read.go
const Error7 = "Error 7. Terjadi kesalahan pada server"              // middleware-read.go
const Error8 = "Error 8. Terjadi kesalahan pada server"              // middleware-read.go
const Error9 = "Error 9. Mohon cek kembali parameter request anda"   // middleware-read.go
const Error10 = "Error 10. Mohon cek kembali parameter request anda" // middleware-read.go
const Error11 = "Error 11. Terjadi kesalahan pada server"            // middleware-read.go
const Error12 = "Error 12. Terjadi kesalahan pada server"            // middleware-read.go
const Error13 = "Error 13. Terjadi kesalahan pada server"            // middleware-read.go
const Error14 = "Error 14. Terjadi kesalahan pada server"            // middleware-read.go
const Error15 = "Error 15. Terjadi kesalahan pada server"            // middleware-read.go
const Error16 = "Error 16. Terjadi kesalahan pada server"            // middleware-read.go
const Error17 = "Error 17. Terjadi kesalahan pada server"            // middleware-read.go
const Error18 = "Error 18. Terjadi kesalahan pada server"            // middleware-read.go
const Error19 = "Error 19. Terjadi kesalahan pada server"            // middleware-read.go
const Error20 = "Error 20. Terjadi kesalahan pada server"            // middleware-read.go
const Error21 = "Error 21. Terjadi kesalahan pada server"            // middleware-read.go
const Error22 = "Error 22. Terjadi kesalahan pada server"            // middleware-read.go
const Error23 = "Error 23. Terjadi kesalahan pada server"            // middleware-read.go
const Error24 = "Error 24. Terjadi kesalahan pada server"            // middleware-read.go
