package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luebken/commenthub/pkg/db"
)

func GetCommentsHandler(db db.Database) gin.HandlerFunc {

	return func(c *gin.Context) {
		owner := c.Param("owner")
		repo := c.Param("repo")
		issuenr := c.Param("issuenr")

		comments := db.GetComments(owner, repo, issuenr)
		c.JSON(http.StatusOK, comments)
	}
}

func PostCommentsHandler(db db.Database) gin.HandlerFunc {

	return func(c *gin.Context) {

		var comment db.Comment

		if err := c.BindJSON(&comment); err != nil {
			return
		}
		owner := c.Param("owner")
		repo := c.Param("repo")
		issuenr := c.Param("issuenr")

		db.PutComment(owner, repo, issuenr, comment)
		c.JSON(http.StatusOK, "s")
	}
}
