package gowler

func ScrapLinks(s string) (hyperlinks []string, err error) {
	hyperLinks := make([]string, 0)
	hyperLinks = append(hyperLinks, "http://github.com/aishraj/")
	return hyperLinks, nil
}
