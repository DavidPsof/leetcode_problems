package aggregate

import "github.com/DavidPsof/leetcode_problems/backend/pkg/api_models"

type Post struct {
	Title string
	Link  string
}

func (p *Post) ToApiResp() api_models.PostResp {
	return api_models.PostResp{
		Title: p.Title,
		Link:  p.Link,
	}
}
