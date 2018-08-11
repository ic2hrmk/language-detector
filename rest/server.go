package rest

import (
	"github.com/emicklei/go-restful"

	"lang-detector/rest/resources"
	"lang-detector/rest/dto"
	"lang-detector/language/naive-detector"
)

func Init() {
	// Chose service which detector to use
	detector := naive_detector.NewNaiveDetector()

	// Init. service
	service := resources.NewLanguageDetectorService(detector)

	// Init. REST container
	ws := new(restful.WebService)

	ws.Consumes("application/json")
	ws.Produces("application/json")

	ws.Route(ws.
		GET("/which-language").
		To(service.WhichLanguageHandler).
		Returns(200, "OK", dto.WhichLanguageResponse{}).
		Returns(401, "BadRequest", dto.ErrorResponse{}).
		Returns(404, "NotFound", dto.ErrorResponse{}).
		Writes(dto.WhichLanguageResponse{}),
	)

	restful.Add(ws)
}
