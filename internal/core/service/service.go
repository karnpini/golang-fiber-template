package service

// type ServiceImpl struct {
// 	ikmCore ports.ClientPort
// }

// func NewService(ikmCore ports.ClientPort) ports.ServicePort {
// 	return &ServiceImpl{ikmCore: ikmCore}
// }

// func (s *ServiceImpl) GetContent(contentId string, lang string, version string) (map[string]interface{}, error) {
// 	contentByte, err := s.ikmCore.GetContent(contentId, lang, version)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var content map[string]interface{}
// 	if err := json.Unmarshal(contentByte, &content); err != nil {
// 		return nil, err
// 	}

// 	return content, nil
// }
