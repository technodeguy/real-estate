package config

func (allowedFormats FileWhiteListType) IsAcceptedMimeType(mt string) bool {
	for _, val := range allowedFormats {
		if val == mt {
			return true
		}
	}

	return false
}
