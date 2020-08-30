package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-blog/models"
	"github.com/gin-blog/pkg/e"
	"github.com/gin-blog/pkg/setting"
	"github.com/gin-blog/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必須大於0")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("狀態只允許0或1")
		var tagId int = -1
		if arg := c.Query("tag_id"); arg != "" {
			tagId = com.StrTo(arg).MustInt()
			valid.Min(tagId, 1, "tag_id").Message("標籤ID必須大於0")
		}
		code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			code = e.SUCCESS
			data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
			data["total"] = models.GetArticleTotal(maps)
		} else {
			for _, err := range valid.Errors {
				log.Print(err.Key, err.Message)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

}
func AddArticle(c *gin.Context) {

}
func EditArticle(c *gin.Context) {

}
func DeleteArticle(c *gin.Context) {

}
