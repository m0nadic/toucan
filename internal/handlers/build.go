package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"toucan/internal/service"
)

type BuildHandler struct {
	Builder service.BuilderService
}

func NewBuildHandler(builderService service.BuilderService) BuildHandler {
	return BuildHandler{
		Builder: builderService,
	}
}

func (bh BuildHandler) Build(ctx *gin.Context) {
	goExe := ctx.PostForm("executable")
	goOS := ctx.PostForm("os")
	goArch := ctx.PostForm("arch")
	goCode := ctx.PostForm("content")

	exeFile, err := bh.Builder.Build(goExe, goOS, goArch, goCode)
	defer os.Remove(exeFile)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+path.Base(exeFile))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(exeFile)
}
