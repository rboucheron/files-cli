package stringutils

func FindWeight(fontType string) [2]string {
	switch fontType {
	case "Thin":
		return [2]string{"100", "normal"}
	case "ExtraLight":
		return [2]string{"200", "normal"}
	case "Light":
		return [2]string{"300", "normal"}
	case "Normal":
		return [2]string{"400", "normal"}
	case "Medium":
		return [2]string{"500", "normal"}
	case "SemiBold":
		return [2]string{"600", "normal"}
	case "Bold":
		return [2]string{"700", "normal"}
	case "ExtraBold":
		return [2]string{"800", "normal"}
	case "Black":
		return [2]string{"900", "normal"}
	case "ThinItalic":
		return [2]string{"100", "italic"}
	case "ExtraLightItalic":
		return [2]string{"200", "italic"}
	case "LightItalic":
		return [2]string{"300", "italic"}
	case "NormalItalic":
		return [2]string{"400", "italic"}
	case "MediumItalic":
		return [2]string{"500", "italic"}
	case "SemiBoldItalic":
		return [2]string{"600", "italic"}
	case "BoldItalic":
		return [2]string{"700", "italic"}
	case "ExtraBoldItalic":
		return [2]string{"800", "italic"}
	case "BlackItalic":
		return [2]string{"900", "italic"}
	default:
		return [2]string{"unknown", "unknown"}
	}
}
