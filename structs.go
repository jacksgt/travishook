package travishook

import (
	"time"
)

/*
* conforming to Travis CI's "Webhook Delivery Format"
* https://docs.travis-ci.com/user/notifications/#Webhooks-Delivery-Format
 */
type WebhookPayload struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Config struct {
		Sudo     bool     `json:"sudo"`
		Dist     string   `json:"dist"`
		Language string   `json:"language"`
		Python   []string `json:"python"`
		Branches struct {
			Only []string `json:"only"`
		} `json:"branches"`
		Cache struct {
			Pip         bool     `json:"pip"`
			Directories []string `json:"directories"`
		} `json:"cache"`
		Deploy struct {
			Provider string `json:"provider"`
			APIKey   struct {
				Secure string `json:"secure"`
			} `json:"api_key"`
			App         string `json:"app"`
			SkipCleanup bool   `json:"skip_cleanup"`
			True        struct {
				Branch []string `json:"branch"`
			} `json:"true"`
		} `json:"deploy"`
		Notifications struct {
			Slack struct {
				Rooms struct {
					Secure string `json:"secure"`
				} `json:"rooms"`
				OnSuccess string `json:"on_success"`
			} `json:"slack"`
			Webhooks string `json:"webhooks"`
		} `json:"notifications"`
		Install      []string `json:"install"`
		BeforeScript []string `json:"before_script"`
		Script       []string `json:"script"`
		Result       string   `json:".result"`
		GlobalEnv    []string `json:"global_env"`
		Group        string   `json:"group"`
	} `json:"config"`
	Type              string      `json:"type"`
	State             string      `json:"state"`
	Status            int         `json:"status"`
	Result            int         `json:"result"`
	StatusMessage     string      `json:"status_message"`
	ResultMessage     string      `json:"result_message"`
	StartedAt         time.Time   `json:"started_at"`
	FinishedAt        time.Time   `json:"finished_at"`
	Duration          int         `json:"duration"`
	BuildURL          string      `json:"build_url"`
	CommitID          int         `json:"commit_id"`
	Commit            string      `json:"commit"`
	BaseCommit        string      `json:"base_commit"`
	HeadCommit        string      `json:"head_commit"`
	Branch            string      `json:"branch"`
	Message           string      `json:"message"`
	CompareURL        string      `json:"compare_url"`
	CommittedAt       time.Time   `json:"committed_at"`
	AuthorName        string      `json:"author_name"`
	AuthorEmail       string      `json:"author_email"`
	CommitterName     string      `json:"committer_name"`
	CommitterEmail    string      `json:"committer_email"`
	PullRequest       bool        `json:"pull_request"`
	PullRequestNumber int         `json:"pull_request_number"`
	PullRequestTitle  string      `json:"pull_request_title"`
	Tag               interface{} `json:"tag"`
	Repository        struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		OwnerName string `json:"owner_name"`
		URL       string `json:"url"`
	} `json:"repository"`
	Matrix []struct {
		ID           int    `json:"id"`
		RepositoryID int    `json:"repository_id"`
		ParentID     int    `json:"parent_id"`
		Number       string `json:"number"`
		State        string `json:"state"`
		Config       struct {
			Sudo     bool   `json:"sudo"`
			Dist     string `json:"dist"`
			Language string `json:"language"`
			Python   string `json:"python"`
			Branches struct {
				Only []string `json:"only"`
			} `json:"branches"`
			Cache struct {
				Pip         bool     `json:"pip"`
				Directories []string `json:"directories"`
			} `json:"cache"`
			Notifications struct {
				Slack struct {
					Rooms struct {
						Secure string `json:"secure"`
					} `json:"rooms"`
					OnSuccess string `json:"on_success"`
				} `json:"slack"`
				Webhooks string `json:"webhooks"`
			} `json:"notifications"`
			Install      []string `json:"install"`
			BeforeScript []string `json:"before_script"`
			Script       []string `json:"script"`
			Result       string   `json:".result"`
			GlobalEnv    []string `json:"global_env"`
			Group        string   `json:"group"`
			Os           string   `json:"os"`
			Addons       struct {
				Deploy struct {
					Provider string `json:"provider"`
					APIKey   struct {
						Secure string `json:"secure"`
					} `json:"api_key"`
					App         string `json:"app"`
					SkipCleanup bool   `json:"skip_cleanup"`
					True        struct {
						Branch []string `json:"branch"`
					} `json:"true"`
				} `json:"deploy"`
			} `json:"addons"`
		} `json:"config"`
		Status         int         `json:"status"`
		Result         int         `json:"result"`
		Commit         string      `json:"commit"`
		Branch         string      `json:"branch"`
		Message        string      `json:"message"`
		CompareURL     string      `json:"compare_url"`
		StartedAt      interface{} `json:"started_at"`
		FinishedAt     interface{} `json:"finished_at"`
		CommittedAt    time.Time   `json:"committed_at"`
		AuthorName     string      `json:"author_name"`
		AuthorEmail    string      `json:"author_email"`
		CommitterName  string      `json:"committer_name"`
		CommitterEmail string      `json:"committer_email"`
		AllowFailure   bool        `json:"allow_failure"`
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
