package internal

func LoadSlide(path string) (Meta, []string, error) {
	meta, content, err := ParseMetadata(path)
	if err != nil {
		return meta, []string{}, err
	}

	slidescontent := ParseSlides(string(content))

	return meta, slidescontent, err
}
