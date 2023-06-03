package counselai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/BearTS/go-gin-monolith/config"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Query(c *gin.Context, reqBody QueryReq) (utils.BaseResponse, error) {
	var finalRes utils.BaseResponse
	var err error

	finalRes.StatusCode = http.StatusInternalServerError
	finalRes.Message = "Internal Server Error"
	finalRes.Success = false

	apiUrl := config.CounselAi.URL + "/cases/query"

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return finalRes, errors.Wrap(err, "[][Marshal]")
	}

	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqJson))

	// Add Headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		finalRes.StatusCode = http.StatusInternalServerError
		finalRes.Message = "Internal Server Error"
		finalRes.Success = false
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return finalRes, errors.Wrap(err, "[NameMatch][ReadAll]")
	}

	// Json Unmarshal
	var data QueryRes

	err = json.Unmarshal(body, &data)
	if err != nil {
		return finalRes, errors.Wrap(err, "[NameMatch][Unmarshal]")
	}

	finalRes.StatusCode = http.StatusOK
	finalRes.Data = data
	finalRes.Success = true
	finalRes.Message = "Success"

	return finalRes, err
}
