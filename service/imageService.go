package service

import (
	"errors"
	"fmt"
	"images/model"
	"images/primary"
	"images/repository"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ImageService interface {
	Create(ctx *gin.Context, request model.ImageRequest) (*model.ImageResponse, error)
	Delete(ctx *gin.Context, imageID string) error
}

type imageService struct {
	ir repository.ImageRepository
}

func NewImageRepository(ir repository.ImageRepository) ImageService {
	return &imageService{
		ir: ir,
	}
}

func (is *imageService) Create(ctx *gin.Context, request model.ImageRequest) (*model.ImageResponse, error) {
	if err := ctx.ShouldBind(&request); err != nil {
		return nil, err
	}
	validate := validator.New()

	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	//validasi gambar
	ext := strings.ToLower(filepath.Ext(request.Avatar.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return nil, errors.New("unsupport file format")
	}

	fmt.Println(ext)

	if err := ctx.SaveUploadedFile(request.Avatar, "assets/"+request.Avatar.Filename); err != nil {
		ctx.HTML(500, "gagal.html", gin.H{"error": "Failed to save file"})
		return nil, err
	}

	img := model.Images{
		Id:         primary.IdRndm(8),
		Path:       request.Avatar.Filename,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	result, err := is.ir.Create(img)
	if err != nil {
		return nil, err
	}

	return &model.ImageResponse{
		Id:         result.Id,
		Path:       result.Path,
		Created_at: result.Created_at,
		Updated_at: result.Updated_at,
	}, nil
}

func (is *imageService) Delete(ctx *gin.Context, imageID string) error {
	id := ctx.PostForm("id")
	if err := ctx.ShouldBind(id); err != nil {
		return nil
	}
	imgPath, err := is.ir.FindByID(imageID)
	if err != nil {
		return nil
	}
	if err := os.Remove("assets/" + imgPath.Path); err != nil {
		return err
	}
	err = is.ir.Delete(imageID)
	if err != nil {
		return err
	}

	return nil
}
