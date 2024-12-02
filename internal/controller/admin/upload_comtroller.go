package admin

import (
	"atom-go/internal/common/upload"
	"atom-go/pkg/response"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// 上传控制器
type UploadController struct{}

// 上传文件
func (u *UploadController) UploadFile(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	fileResult, err := upload.New(
		upload.SetLimitType([]string{
			"audio/mpeg",
			"audio/x-m4a",
			"video/mp4",
			"video/x-flv",
			"video/x-m4v",
			"application/x-zip-compressed",
			"application/pdf",
			"application/msword",
			"application/vnd.ms-excel",
			"application/vnd.ms-powerpoint",
			"application/vnd.openxmlformats-officedocument.presentationml.presentation",
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		}),
	).SetFile(&upload.File{
		FileName:    fileHeader.Filename,
		FileType:    fileHeader.Header.Get("Content-Type"),
		FileHeader:  fileHeader.Header,
		FileContent: fileContent,
	}).Save()

	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().SetData(fileResult).Send(ctx)
}

// 上传图片
func (u *UploadController) UploadImage(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	fileResult, err := upload.New(
		upload.SetLimitType([]string{
			"image/jpeg",
			"image/png",
			"image/svg+xml",
		}),
	).SetFile(&upload.File{
		FileName:    fileHeader.Filename,
		FileType:    fileHeader.Header.Get("Content-Type"),
		FileHeader:  fileHeader.Header,
		FileContent: fileContent,
	}).Save()

	if err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().SetData(fileResult).Send(ctx)
}
