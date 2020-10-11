package model

type Response struct {
	HTMLVersion            string           `json:"html_version"`
	PageTitle              string           `json:"page_title"`
	HeadingCount           HeadingStructure `json:"heading_count"`
	InternalLinksCount     int64            `json:"internal_links_count"`
	ExternalLinksCount     int64            `json:"external_links_count"`
	InaccessibleLinksCount int64            `json:"inaccessible_links_count"`
	LoginFormExist         bool             `json:"login_form_exist"`
}

//there are six levels of heading according to W3C
//ref : https://www.w3.org/MarkUp/html3/headings.html#:~:text=HTML%20defines%20six%20levels%20of,level%20and%20H6%20the%20least.
type HeadingStructure struct {
	H1Count int64 `json:"h1_count"`
	H2Count int64 `json:"h2_count"`
	H3Count int64 `json:"h3_count"`
	H4Count int64 `json:"h4_count"`
	H5Count int64 `json:"h5_count"`
	H6Count int64 `json:"h6_count"`
}
