package v1

import (
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/pkg/r"
	"github.com/Walk2future/bi-chatgpt-golang-python/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListUserTable godoc
//
//	@Summary	ListUserTable
//	@Produce	json
//	@Tags		TableApi
//	@Param		pageRequest	body	requests.Page	true	"登录请求参数"
//	@Accept		json
//	@Success	0		{object}	[]models.UserChart	"成功"
//	@Failure	40002	{object}	r.Response			"参数错误"
//	@Failure	40003	{object}	r.Response			"系统错误"
//	@Router		/table/list [post]
func ListUserTable(c *gin.Context) {
	var req requests.Page
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, r.PARAMS_ERROR)
		return
	}
	var tableService service.TableService
	list, err := tableService.ListTable(&req)
	if err != nil {
		c.JSON(http.StatusOK, r.FAIL.WithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, r.OK.WithData(list))
}
