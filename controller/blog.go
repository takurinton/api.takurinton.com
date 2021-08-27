package controller

import (
	"net/http"
	"portfolio/pagination"
	"portfolio/service"
	"strconv"

	"portfolio/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllPosts(c *gin.Context) {
	h := service.Blog{}

	results, err := h.GetAllPosts()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, results)
	}
}

func GetAllPostsReverse(c *gin.Context) {
	h := service.Blog{}

	results, err := h.GetAllPostsReverse()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, results)
	}
}

func GetPosts(c *gin.Context) {
	h := service.Blog{}

	// クエリパラメータの実装
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	_category := c.DefaultQuery("category", "")

	p, err := h.GetPosts(page, _category)

	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		category, next, prev := pagination.GetParams(page, p.TotalPage, p.NextPage, p.PrevPage, _category)
		res := map[string]interface{}{
			"next":      next,
			"previous":  prev,
			"category":  category,
			"results":   p.Records,
			"current":   page,
			"total":     p.TotalPage - 1,
			"page_size": 5,
			"first":     1,
			"last":      p.TotalPage,
		}
		c.JSONP(http.StatusOK, res)
	}
}

func CreatePost(c *gin.Context) {
	h := service.Blog{}

	var post model.Post

	if err := c.BindJSON(&post); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.CreatePost(post); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"title": post.Title})
	}
}

func GetPost(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	results, err := h.GetPost(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, results)
	}
}

func GetPostAdmin(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	results, err := h.GetPostAdmin(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSONP(http.StatusNotFound, nil)
		} else {
			c.JSONP(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSONP(http.StatusOK, results)
	}
}

func UpdateDetailPost(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	var post model.Post

	if err := c.BindJSON(&post); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.UpdatePost(id, post); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"title": post.Title})
	}
}

func DeleteDetailPost(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	title := c.PostForm("title")
	if err := h.DeletePost(id); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"title": title})
	}
}

func PostComment(c *gin.Context) {
	h := service.Blog{}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var comment model.BlogappComment

	c.BindJSON(&comment)

	if err := h.PostComment(id, comment.Name, comment.Contents); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, comment)
	}
}

func CreateCategory(c *gin.Context) {
	h := service.Blog{}
	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.CreateCategory(category.Name); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"name": category.Name})
	}
}

func GetDetailCategory(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	res, err := h.GetDetailCategory(id)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusOK, gin.H{"category": res})
	}
}

func UpdateCategory(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.UpdateCategory(id, category.Name); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"name": category.Name})
	}
}

func GetCategory(c *gin.Context) {
	h := service.Blog{}

	res, err := h.GetCategory()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusOK, gin.H{"category": res})
	}
}

func DeleteCategory(c *gin.Context) {
	h := service.Blog{}
	id := c.Param("id")

	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.DeleteCategory(id); err != nil {
		c.JSONP(http.StatusInternalServerError, nil)
	} else {
		c.JSONP(http.StatusCreated, gin.H{"name": category.Name})
	}
}
