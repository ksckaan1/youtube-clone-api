package subscribe

type SubList struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Count  int64  `json:"count"`
	Subs   []*Sub `json:"subs"`
}

type Sub struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
}
