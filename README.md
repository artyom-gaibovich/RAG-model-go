# rag-model-go

### Stack
**Database**: _PostgreSQL_  
**Programming language**: _Go_  
**DI**: _google/wire_

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

### Database structure

base_models  
```
CREATE TABLE IF NOT EXISTS base_models (
id SERIAL PRIMARY KEY NOT NULL,
created_at TIMESTAMPTZ DEFAULT current_timestamp,
updated_at TIMESTAMPTZ DEFAULT current_timestamp,
deleted_at TIMESTAMPTZ
);
```

### google/wire

Windows: `go run github.com/google/wire/cmd/wire .\config\` where `.\config\` is folder