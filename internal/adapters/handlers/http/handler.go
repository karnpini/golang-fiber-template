package handlers

// type Handler struct {
// 	service ports.ServicePort
// }

// func NewHandler(service ports.ServicePort) *Handler {
// 	return &Handler{
// 		service: service,
// 	}
// }

// func (h *Handler) GetContentHandler(c *fiber.Ctx) error {
// 	contentId := c.Params("contentId")
// 	lang := c.Query("lang", "th-TH")
// 	version := c.Query("version", "")
// 	if contentId == "" {
// 		interceptors.SendJSONResponse(c, fiber.StatusBadRequest, e.INVALID_PARAMS, nil)
// 		return nil
// 	}

// 	contentData, err := h.service.GetContent(contentId, lang, version)
// 	if err != nil {
// 		interceptors.SendJSONResponse(c, fiber.StatusInternalServerError, e.ERROR_GET_CONTENT_VERINT, err.Error())
// 		return nil
// 	}

// 	interceptors.SendJSONResponse(c, fiber.StatusOK, e.SUCCESS, contentData)
// 	return nil
// }
