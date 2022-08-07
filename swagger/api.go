package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var db *gorm.DB

func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"data":    posts,
	})
}

func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "post not found",
			"data":    "",
		})
	}
	return post
}

// @Summary      查询
// @Description  查询
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "pid"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts/{id} [get]
func Show(c *gin.Context) {
	post := getById(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"data":    post,
	})
}

// @Summary      添加post
// @Description  添加post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param content body string true "json"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts [post]
func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	post.Status = "Active"
	db.Create(&post)

	c.JSON(http.StatusOK, Response{
		Msg:  "",
		Data: post,
	})
}

func Delete(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		return
	}

	db.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
		"data":    "",
	})
}

func Update(c *gin.Context) {
	oldpost := getById(c)
	var newpost Post

	if err := c.ShouldBindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    "",
		})

		return
	}

	oldpost.Title = newpost.Title
	oldpost.Des = newpost.Des
	if newpost.Status != "" {
		oldpost.Status = newpost.Status
	}

	db.Save(&oldpost)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post has been updated",
		"data":    oldpost,
	})
}
