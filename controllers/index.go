package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsongyan/wblog/helpers"
	"github.com/wangsongyan/wblog/models"
	"net/http"
	"strconv"
)

func IndexGet(c *gin.Context) {
	posts, err := models.ListPost("")
	if err == nil {
		for _, post := range posts {
			post.Tags, _ = models.ListTagByPostId(strconv.FormatUint(uint64(post.ID), 10))
		}
		c.HTML(http.StatusOK, "index/index.html", gin.H{
			"posts":    posts,
			"tags":     helpers.ListTag(),
			"archives": helpers.ListArchive(),
		})
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func AdminIndex(c *gin.Context) {
	user, _ := c.Get("User")
	/*c.JSON(http.StatusOK, gin.H{
		"data": user,
	})*/
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"pageCount":    models.CountPage(),
		"postCount":    models.CountPost(),
		"tagCount":     models.CountTag(),
		"commentCount": models.CountComment(),
		"user":         user,
	})
}