package models

type TaskRequest struct {
	Links     []string `json:"links,omitempty"`
	LinksList []int    `json:"links_list,omitempty"`
}

type TaskResponse struct {
	Links    map[string]string `json:"links,omitempty"`
	LinksNum int               `json:"id,omitempty"`
}
