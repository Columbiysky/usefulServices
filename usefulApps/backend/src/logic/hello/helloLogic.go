package logic

import (
	"fmt"
	"io"
	"log"
	"main/logic/hello/params"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, "hello")
}

func HelloObject(context *gin.Context, obj params.HelloObject) {
	context.JSON(http.StatusOK, obj)
}

func HelloFromPyServ(context *gin.Context) {
	resp, err := http.Get("http://localhost:8000/")

	if err != nil {
		print(err)
	}

	body, err_ := io.ReadAll(resp.Body)

	if err_ != nil {
		print(err_)
	}

	resp.Body.Close()
	context.JSON(http.StatusOK, string(body))
}

func HelloFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	// context.SaveUploadedFile(file, dst)

	context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
