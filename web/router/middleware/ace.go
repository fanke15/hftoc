package middleware

import (
	"github.com/fanke15/hftoc/pkg/basic"
	"github.com/fanke15/hftoc/pkg/lib/conf"
	"github.com/fanke15/hftoc/pkg/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

var (
	aceDir = basic.AnySliceToStr(basic.StrNull, conf.New().GetString("project.dir.static"), "ace/")
)

// 模板初始化设置
func InitAce(name, inner string, data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		tpl, err := ace.Load(
			basic.AnySliceToStr(basic.StrNull, aceDir, name),
			basic.AnySliceToStr(basic.StrNull, aceDir, inner),
			&ace.Options{DynamicReload: true})
		if err != nil {
			log.Error(err.Error(), "func", "InitAce")
			return
		}
		_ = tpl.Execute(c.Writer, data)
	}
}
