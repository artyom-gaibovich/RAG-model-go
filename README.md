# rag-model-go

### Stack
**Database**: _PostgreSQL_  
**Programming language**: _Go_  
**DI**: _google/wire_

### Run && Build quick

1. `docker build -t server_generation .`
2. `docker run -p 8080:8080 server_generation`


### How to start

`go run main.go`

### Dependency
`go get -u github.com/antonfisher/nested-logrus-formatter`  
`go get -u github.com/gin-gonic/gin`  
`go get -u golang.org/x/crypto`  
`go get -u gorm.io/gorm`  
`go get -u gorm.io/driver/postgres`  
`go get -u github.com/sirupsen/logrus`  
`go get -u github.com/joho/godotenv`  

### ENV structure

```
PORT=8080
# Application
APPLICATION_NAME=rag-model-restful-api

# Database
DB_DSN="host=localhost user=root password=root dbname=rag-model-api port=5432"

# Logging
LOG_LEVEL=DEBUG
```

### google/wire

Windows: `go run github.com/google/wire/cmd/wire .\config\` where `.\config\` is folder