package aisvc

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func (g *AiSvcImpl) UploadReport(c *gin.Context) (utils.BaseResponse, error) {
	var baseRes utils.BaseResponse
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	file, _ := c.FormFile("report")

	// open file
	src, err := file.Open()
	if err != nil {
		return baseRes, err
	}
	defer src.Close()

	//Write data from received file to temp file
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return baseRes, err
	}

	uploadedUrl, err := newCatBox(nil).rawUpload(fileBytes, file.Filename)
	if err != nil {
		return baseRes, err
	}

	// Create new report
	var report tables.Reports
	report.Url = uploadedUrl

	report, err = g.reportsGorm.CreateNewReport(c, report)

	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "Success"
	baseRes.Data = report

	return baseRes, err

}

type catbox struct {
	Client   *http.Client
	Userhash string
}

func newCatBox(client *http.Client) *catbox {
	if client == nil {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &catbox{
		Client: client,
	}
}

func (cat *catbox) rawUpload(b []byte, name string) (string, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("userhash", cat.Userhash)
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(name))
		if err != nil {
			return
		}
		if _, err = io.Copy(part, bytes.NewBuffer(b)); err != nil {
			return
		}
	}()
	ENDPOINT := "https://catbox.moe/user/api.php"
	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
