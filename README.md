# crud package
只需要一个 gorm 的 model。自动生成 REST 的 CRUD 接口。快速实现后台的增删改查功能。
需要你的项目使用的gin+gorm， 并且golang版本 > 1.18 需要支持范型。


同时查询接口实现了基于 pocketbase 的查询协议 https://pocketbase.io/docs/api-rules-and-filters/

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

### Create

```
POST /api/product
BODY:
{
    "code": "aaa112",
    "price": 100
}
RESPONSE:
{
    "code": 0,
    "message": "success",
    "data": {
        "ID": 1,
        "CreatedAt": "2025-03-04T00:10:29.826052+08:00",
        "UpdatedAt": "2025-03-04T00:10:29.826052+08:00",
        "DeletedAt": null,
        "Code": "aaa112",
        "Price": 100
    }
}
```

### Read

```
GET /api/product/:id
response:
{
    "code": 0,
    "message": "success",
    "data": {
        "ID": 1,
        "CreatedAt": "2025-03-04T00:10:29.826052+08:00",
        "UpdatedAt": "2025-03-04T00:10:29.826052+08:00",
        "DeletedAt": null,
        "Code": "aaa112",
        "Price": 100
    }
}
```

### Update

```
PUT /api/product/:id
BODY:
{
    "code": "aaa112",
    "price": 200
}
RESPONSE:
{
    "code": 0,
    "message": "success",
    "data": {}
}
```

### List

```
GET /api/product
RESPONSE:
{
    "code": 0,
    "message": "success",
    "data": {
        "list": [
            {
                "ID": 1,
                "CreatedAt": "2025-03-04T00:10:29.826052+08:00",
                "UpdatedAt": "2025-03-04T00:12:01.923514+08:00",
                "DeletedAt": null,
                "Code": "aaa112",
                "Price": 200
            }
        ],
        "pagination": {
            "pageNum": 1,
            "pageSize": 20,
            "totalPages": 1,
            "totalCount": 1
        }
    }
}
```

### Delete

```
DELETE /api/product/:id
RESPONSE:
{
    "code": 0,
    "message": "success",
    "data": {}
}
```

## License

MIT
