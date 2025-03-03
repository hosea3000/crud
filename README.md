# crud package
This is a package for creating CRUD (Create, Read, Update, Delete) restful APIs from a gorm model

## Usage

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hosea3000/crud"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "product"
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router := r.Group("/api")

	c := crud.NewCRUD[Product](db, "product")
	c.RegisterRoutes(router)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

```

## Endpoints

### List

```
GET /api/product
```

### Create

```
POST /api/product
```

### Read

```
GET /api/product/:id
```

### Update

```
PUT /api/product/:id
```

### Delete

```
DELETE /api/product/:id
```

## License

MIT
