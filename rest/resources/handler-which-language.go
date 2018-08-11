package resources

import (
	"fmt"
	"net/http"
	"log"

	"github.com/emicklei/go-restful"

	"lang-detector/rest/dto"
	"lang-detector/rest/helper"
)

func (rcv *LanguageDetectorService) WhichLanguageHandler(req *restful.Request, resp *restful.Response) {
	//
	// Validate parameters
	//
	phrase := req.QueryParameter("phrase")
	if phrase == "" {
		err := fmt.Errorf("parameter 'phrase' is empty")

		log.Println("[detect-language]", err.Error())
		helper.HandleRESTException(resp, err, http.StatusBadRequest)
		return
	}

	//
	// Perform action
	//
	language, err := rcv.detector.GetLanguage(phrase)
	if err != nil {
		log.Println("[detect-language]", err.Error())
		helper.HandleRESTException(resp, err, http.StatusInternalServerError)
		return
	}

	resp.WriteEntity(dto.WhichLanguageResponse{Language: language})
}
