package github

type PushEvent struct {
	//Repository json_repo `json:"repository"`
	//HC         json_head_commit `json:"head_commit"`
	Commits    []*commit `json:"commits"`
	Repository repo      `json:"repository"`
	Ref        string    `json:"ref"`
}

type commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    author `json:"author"`
}

type author struct {
	Username string `json:"username"`
	RealName string `json:"name"`
}

type repo struct {
	URL string `json:"url"`
}
