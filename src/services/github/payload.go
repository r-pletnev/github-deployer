package github

import "strings"

type PushPayload struct {
	Ref string `json:"ref"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func(p *PushPayload) BranchName() string {
	sl := strings.Split(p.Ref, "/")
	return sl[len(sl)-1]
}

func (p *PushPayload)IsValidBranch(branchName string) bool {
	return p.BranchName() == branchName
}

