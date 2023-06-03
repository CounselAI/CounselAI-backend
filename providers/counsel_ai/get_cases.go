package counselai

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/BearTS/go-gin-monolith/config"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

// /cases/ ->get request
func GetCases(c *gin.Context) (utils.BaseResponse, error) {
	var finalRes utils.BaseResponse
	var err error

	finalRes.StatusCode = http.StatusInternalServerError
	finalRes.Message = "Internal Server Error"
	finalRes.Success = false

	url := config.CounselAi.URL + "/cases/"

	req, _ := http.NewRequest("GET", url, nil)

	// Add Headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		finalRes.StatusCode = http.StatusInternalServerError
		finalRes.Message = "Internal Server Error"
		finalRes.Success = false
		return finalRes, err
	}

	finalRes.StatusCode = response.StatusCode
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		finalRes.StatusCode = http.StatusInternalServerError
		finalRes.Message = "Internal Server Error"
		finalRes.Success = false
		return finalRes, err
	}

	// Json Unmarshal
	var data GetCasesRes
	err = json.Unmarshal(body, &data)
	if err != nil {
		finalRes.StatusCode = http.StatusInternalServerError
		finalRes.Message = "Internal Server Error"
		finalRes.Success = false
	}

	finalRes.Data = data
	finalRes.Success = true
	finalRes.StatusCode = response.StatusCode
	finalRes.Message = "Success"

	return finalRes, err
}
