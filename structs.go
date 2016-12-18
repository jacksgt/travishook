package travishook

/*
* conforming to Travis CI's "Webhook Delivery Format"
* https://docs.travis-ci.com/user/notifications/#Webhooks-Delivery-Format
 */
type WebhookPayload struct {
	ID             int    `json:"id"`
	Number         string `json:"number"`
	Status         int    `json:"status"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	StatusMessage  string `json:"status_message"`
	Commit         string `json:"commit"`
	Branch         string `json:"branch"`
	Message        string `json:"message"`
	CompareURL     string `json:"compare_url"`
	CommittedAt    string `json:"committed_at"`
	CommitterName  string `json:"committer_name"`
	CommitterEmail string `json:"committer_email"`
	AuthorName     string `json:"author_name"`
	AuthorEmail    string `json:"author_email"`
	Type           string `json:"type"`
	BuildURL       string `json:"build_url"`
	Repository     struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		OwnerName string `json:"owner_name"`
		URL       string `json:"url"`
	} `json:"repository"`
	Config struct {
		Notifications struct {
			Webhooks []string `json:"webhooks"`
		} `json:"notifications"`
	} `json:"config"`
	Matrix []struct {
		ID           int         `json:"id"`
		RepositoryID int         `json:"repository_id"`
		Number       string      `json:"number"`
		State        string      `json:"state"`
		StartedAt    interface{} `json:"started_at"`
		FinishedAt   interface{} `json:"finished_at"`
		Config       struct {
			Notifications struct {
				Webhooks []string `json:"webhooks"`
			} `json:"notifications"`
		} `json:"config"`
		Status         interface{} `json:"status"`
		Log            string      `json:"log"`
		Result         interface{} `json:"result"`
		ParentID       int         `json:"parent_id"`
		Commit         string      `json:"commit"`
		Branch         string      `json:"branch"`
		Message        string      `json:"message"`
		CommittedAt    string      `json:"committed_at"`
		CommitterName  string      `json:"committer_name"`
		CommitterEmail string      `json:"committer_email"`
		AuthorName     string      `json:"author_name"`
		AuthorEmail    string      `json:"author_email"`
		CompareURL     string      `json:"compare_url"`
	} `json:"matrix"`
}

/*
* conforming to Travis CI's "Verifying Webhook requests"
* https://docs.travis-ci.com/user/notifications/#Verifying-Webhook-requests
 */
type APIConfig struct {
	Config struct {
		Host        string `json:"host"`
		ShortenHost string `json:"shorten_host"`
		Assets      struct {
			Host string `json:"host"`
		} `json:"assets"`
		Pusher struct {
			Key string `json:"key"`
		} `json:"pusher"`
		Github struct {
			APIURL string   `json:"api_url"`
			Scopes []string `json:"scopes"`
		} `json:"github"`
		Notifications struct {
			Webhook struct {
				PublicKey string `json:"public_key"`
			} `json:"webhook"`
		} `json:"notifications"`
	} `json:"config"`
}
