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

// GetTags 獲取標籤列表
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTags 新建文章標籤
func AddTags(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名稱不能為空")
	valid.MaxSize(name, 100, "name").Message("名稱最長為100個字元")
	valid.Required(createdBy, "created_by").Message("創建人不能為空")
	valid.MaxSize(createdBy, 100, "created_by").Message("創建人最長為100個字元")
	valid.Range(state, 0, 1, "state").Message("狀態只能是0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
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

// PutTags 更新指定文章標籤
func PutTags(c *gin.Context) {

}

// DelTags 刪除指定文章標簽
func DelTags(c *gin.Context) {

}
