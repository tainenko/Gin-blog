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
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("tital")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("標籤ID必須大於0")
	valid.Required(title, "title").Message("標題不能為空")
	valid.Required(desc, "desc").Message("簡述不能為空")
	valid.Required(content, "content").Message("內容不能為空")
	valid.Required(createdBy, "created_by").Message("創建人不能為空")
	valid.Range(state, 0, 1, "state").Message("狀態只允許0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("tital")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("狀態只允許0或1")
	}
	valid.Min(id, 1, "id").Message("ID必須大於0")
	valid.MaxSize(title, 100, "title").Message("標題最長為100個字元")
	valid.MaxSize(desc, 255, "desc").Message("敘述最長為255個字元")
	valid.MaxSize(content, 65535, "content").Message("內容最長為65535個字元")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能為空白")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最長為100個字元")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface)
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func DeleteArticle(c *gin.Context) {

}
