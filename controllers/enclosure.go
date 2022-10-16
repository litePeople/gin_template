package controllers

import (
	"gin_template/consts"
	"gin_template/models"
	"mime/multipart"

	"github.com/pkg/errors"
)

type EnclosureCtl struct {
	BaseCtl
}

// EnclosureUpload 上传附件
// @Summary 上传附件
// @Description 上传附件
// @ID WebEnclosureUpload
// @Tags WEB Client,WEB Client的附件管理
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param file formData file true "附件文件，最大20M"
// @Success 200 {object} ginx.Resp{data=models.Enclosure} "处理成功"
// @Failure 1001 {object} ginx.Resp "参数错误"
// @Failure 1002 {object} ginx.Resp "业务逻辑处理错误"
// @Security ApiKeyAuth
// @Router /v1/web/enclosure/upload [post]
func (t *EnclosureCtl) EnclosureUpload() {
	var (
		upFile    *multipart.FileHeader
		enclosure *models.Enclosure
		err       error
	)
	upFile, err = t.Ctx.FormFile("file")
	if nil != err {
		t.JSONE(consts.ApiErrCodeParam, err)
		return
	}

	if upFile.Size > 20*1024*1024 {
		t.JSONE(consts.ApiErrCodeParam, errors.New("文件太大"))
		return
	}

	t.JSONS(enclosure)
}
