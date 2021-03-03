package RestResponse

import "github.com/gin-gonic/gin"

type RestResponse struct {
	Count   int         `json:"count"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//TODO: @Error Response 404
func NewErrorNotFoundResponse(Mode string, Status string, Message string) *RestResponse {
	// Mode Production = true
	if Mode == gin.ReleaseMode {
		return &RestResponse{
			Count:   0,
			Status:  Status,
			Data:    nil,
			Message: Message,
		}
	}

	// Mode Production = false
	return &RestResponse{
		Count:   0,
		Status:  Status,
		Data:    nil,
		Message: "Not_Found",
	}
}

//TODO: @Error Response 500
func NewErrorInternalResponse(Mode string, Status string, Message string) *RestResponse {
	// Mode Production = true
	if Mode == gin.ReleaseMode {
		return &RestResponse{
			Count:   0,
			Status:  Status,
			Data:    nil,
			Message: "Internal_Server_Error",
		}
	}

	// Mode Production = false
	return &RestResponse{
		Count:   0,
		Status:  Status,
		Data:    nil,
		Message: Message,
	}
}

//TODO: @Success Response
func NewSuccessResponse(Count int, Status string, Data interface{}, Message string) *RestResponse {
	return &RestResponse{
		Count:   Count,
		Status:  Status,
		Data:    Data,
		Message: Message,
	}
}
