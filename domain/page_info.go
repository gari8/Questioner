package domain

type PageInfo struct {
	EndCursor   int  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
}
