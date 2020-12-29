package spacestation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yangjian.com/basicarchitecture/api/middleware"
	"yangjian.com/basicarchitecture/api/services"
)

type SpaceStationController struct {
}

type requestDicSpaceStation struct {
	ShortName string  `binding:"-" json:"short_name" form:"short_name"`
	CName     string  `binding:"-" json:"c_name" form:"c_name"`
	Py        string  `binding:"-" json:"py" form:"py"`
	Latitude  float64 `binding:"omitempty,latitude" json:"latitude" form:"latitude"`
	Longitude float64 `binding:"omitempty,longitude" json:"longitude" form:"longitude"`
	KeyWord   string  `form:"key_word" json:"key_word" binding:"-"`
	Page      int     `form:"page" json:"page" binding:"-"`
	PerPage   int     `form:"per_page" json:"per_page" binding:"-"`
}

// @Summary 站点增加
// @Description 增加站点字典表
// @Tags 站点
// @Accept json
// @Produce json
// @Param short_name query string true "站点简称"
// @Param c_name query string false "站点中文名称"
// @Param latitude query float64 false "纬度"
// @Param longitude query float64 false "经度"
// @Success 201 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/dic [post]
func (ssController *SpaceStationController) AddDicSpaceStation(ctx *gin.Context) {

	var rdsStation requestDicSpaceStation
	if err := ctx.ShouldBind(&rdsStation); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	err := services.AddStationDic(rdsStation.ShortName, rdsStation.CName, rdsStation.Latitude, rdsStation.Longitude)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, middleware.MarkSuccess(nil))
}

// @Summary 站点删除
// @Description 删除站点字典表
// @Tags 站点
// @Accept json
// @Produce json
// @Param short_name path string true "站点简称"
// @Success 204 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/dic/{short_name} [delete]
func (ssController *SpaceStationController) DestroyDicSpaceStation(ctx *gin.Context) {

	shortName := ctx.Param("short_name")
	if shortName == "" {
		ctx.JSON(http.StatusBadRequest, middleware.MarkErrorParam("short_name is master"))
		return
	}
	err := services.DestroyStationDic(shortName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, middleware.MarkSuccess(nil))
}

// @Summary 站点更改
// @Description 修改站点字典表
// @Tags 站点
// @Accept json
// @Produce json
// @Param short_name query string true "站点简称"
// @Param c_name query string false "站点中文名称"
// @Param latitude query float64 false "纬度"
// @Param longitude query float64 false "经度"
// @Success 200 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/dic [put]
func (ssController *SpaceStationController) UpdateDicSpaceStation(ctx *gin.Context) {

	var rdsStation requestDicSpaceStation
	if err := ctx.ShouldBind(&rdsStation); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	err := services.UpdateStationDic(rdsStation.ShortName, rdsStation.CName, rdsStation.Latitude, rdsStation.Longitude)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, middleware.MarkSuccess(nil))
}

// @Summary 站点查询
// @Description 查询单个站点信息
// @Tags 站点
// @Accept json
// @Produce json
// @Param short_name path string true "站点简称"
// @Success 200 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/dic/{short_name} [get]
func (ssController *SpaceStationController) SelectDicSpaceStation(ctx *gin.Context) {

	shortName := ctx.Param("short_name")
	if shortName == "" {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError("short_name is master"))
		return
	}
	dic, err := services.SelectStationDic(shortName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, middleware.MarkSuccess(dic))
}

// @Summary 站点列表查询
// @Description 查询符合条件的站点信息
// @Tags 站点
// @Accept json
// @Produce json
// @Param page query int false "当前页数"
// @Param per_page query int false "每页条数"
// @Success 200 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/dic [get]
func (ssController *SpaceStationController) ListDicSpaceStation(ctx *gin.Context) {

	var rdsStation requestDicSpaceStation
	if err := ctx.ShouldBindQuery(&rdsStation); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	if ctx.Query("page") == "" || rdsStation.Page <= 0 {
		rdsStation.Page = 1
	}
	if ctx.Query("per_page") == "" || rdsStation.Page <= 0 {
		rdsStation.PerPage = 10
	}
	dic, err := services.ListStationDic(rdsStation.Page, rdsStation.PerPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, middleware.MarkSuccess(dic))
}

// @Summary 站点搜索
// @Description 搜索地点的基本信息
// @Tags 站点
// @Accept json
// @Produce json
// @Param key_word query string true "关键字，可以是站号、站名、以及站名拼音的简拼和全拼"
// @Param page query int false "当前页数"
// @Param per_page query int false "每页条数"
// @Success 200 {object} middleware.ResponseWrapper
// @Failure 400 {object} middleware.ResponseWrapper
// @Router /station/search [get]
func (ssController *SpaceStationController) SearchDicSpaceStation(ctx *gin.Context) {

	var rdsStation requestDicSpaceStation
	if err := ctx.ShouldBindQuery(&rdsStation); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	if ctx.Query("page") == "" || rdsStation.Page <= 0 {
		rdsStation.Page = 1
	}
	if ctx.Query("per_page") == "" || rdsStation.Page <= 0 {
		rdsStation.PerPage = 10
	}
	dic, err := services.SearchStationDic(rdsStation.KeyWord, rdsStation.Page, rdsStation.PerPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.MarkError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, middleware.MarkSuccess(dic))
}
