# Lessonshare API
REST API for surfacing Lessonshare data; written in Golang.

- [x] Introduce Postgresql
- [x] Config
- [ ] Abstract response logic from handlers
- [ ] User auth (jwt, oauth 2 with google/fb login)
- [x] Logging
- [ ] Error handling
- [ ] Expand on API functionality

## Installation
* To start the REST API run ```go run main.go``` in a terminal window
* To start the development postgresql server (vagrant box), ```cd db``` and ```vagrant up```.
    * Thanks to jackdb for providing a postgresql [vagrant vm](https://github.com/jackdb/pg-app-dev-vm)
* Execute ```./db/seed.sql``` to create the tables in postgresql and fill with seed data.