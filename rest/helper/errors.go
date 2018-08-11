package helper

import (
	"fmt"
	"encoding/json"

	"lang-detector/rest/dto"

	"github.com/emicklei/go-restful"
)

func HandleRESTException(response *restful.Response, err error, httpStatus int) {
	if err == nil {
		err = fmt.Errorf("internal error")
	}

	errEntity := &dto.ErrorResponse{
		Message: err.Error(),
	}

	jsonString, _ := json.Marshal(errEntity)

	response.WriteErrorString(httpStatus, string(jsonString))
}
