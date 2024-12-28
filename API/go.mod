module github.com/Joao-lucas-felix/DevBook/API

go 1.18

replace github.com/Joao-lucas-felix/DevBook/API/src/controllers => ./src/controllers

require github.com/gorilla/mux v1.8.1

require github.com/joho/godotenv v1.5.1

require github.com/lib/pq v1.10.9

require github.com/badoux/checkmail v1.2.4

require golang.org/x/crypto v0.31.0
