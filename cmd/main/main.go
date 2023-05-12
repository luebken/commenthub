package main

import (
	"github.com/gin-gonic/gin"

	"github.com/luebken/commenthub/pkg/db"
	"github.com/luebken/commenthub/pkg/web"
)

func main() {
	db := db.Database{}
	db.Open()
	defer db.Close()

	r := gin.Default()

	r.GET("/comments/:owner/:repo/issues/:issuenr", web.GetCommentsHandler(db))
	r.POST("/comments", web.PostCommentsHandler(db))

	r.Run()
}
