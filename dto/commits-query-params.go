package dto

type CommitQueryParams struct {
	SHA  string
	Date string
}

func (p CommitQueryParams) String() string {
	queryString := ""
	if p.SHA != "" {
		queryString += "sha=" + p.SHA
	}
	if p.Date != "" {
		if queryString != "" {
			queryString += "&"
		}
		queryString += "date=" + p.Date
	}
	if queryString != "" {
		return "?" + queryString
	}
	return queryString
}
