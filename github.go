package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Post　...postされれてくるJSONの型をここで決めているらしい
type Post struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int    `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Body              string      `json:"body"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		ClosedAt          interface{} `json:"closed_at"`
		MergedAt          interface{} `json:"merged_at"`
		MergeCommitSha    interface{} `json:"merge_commit_sha"`
		Assignee          interface{} `json:"assignee"`
		Milestone         interface{} `json:"milestone"`
		CommitsURL        string      `json:"commits_url"`
		ReviewCommentsURL string      `json:"review_comments_url"`
		ReviewCommentURL  string      `json:"review_comment_url"`
		CommentsURL       string      `json:"comments_url"`
		StatusesURL       string      `json:"statuses_url"`
		Head              struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
					Login             string `json:"login"`
					ID                int    `json:"id"`
					AvatarURL         string `json:"avatar_url"`
					GravatarID        string `json:"gravatar_id"`
					URL               string `json:"url"`
					HTMLURL           string `json:"html_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					OrganizationsURL  string `json:"organizations_url"`
					ReposURL          string `json:"repos_url"`
					EventsURL         string `json:"events_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"owner"`
				Private          bool        `json:"private"`
				HTMLURL          string      `json:"html_url"`
				Description      string      `json:"description"`
				Fork             bool        `json:"fork"`
				URL              string      `json:"url"`
				ForksURL         string      `json:"forks_url"`
				KeysURL          string      `json:"keys_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				TeamsURL         string      `json:"teams_url"`
				HooksURL         string      `json:"hooks_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				EventsURL        string      `json:"events_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BranchesURL      string      `json:"branches_url"`
				TagsURL          string      `json:"tags_url"`
				BlobsURL         string      `json:"blobs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				TreesURL         string      `json:"trees_url"`
				StatusesURL      string      `json:"statuses_url"`
				LanguagesURL     string      `json:"languages_url"`
				StargazersURL    string      `json:"stargazers_url"`
				ContributorsURL  string      `json:"contributors_url"`
				SubscribersURL   string      `json:"subscribers_url"`
				SubscriptionURL  string      `json:"subscription_url"`
				CommitsURL       string      `json:"commits_url"`
				GitCommitsURL    string      `json:"git_commits_url"`
				CommentsURL      string      `json:"comments_url"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				ContentsURL      string      `json:"contents_url"`
				CompareURL       string      `json:"compare_url"`
				MergesURL        string      `json:"merges_url"`
				ArchiveURL       string      `json:"archive_url"`
				DownloadsURL     string      `json:"downloads_url"`
				IssuesURL        string      `json:"issues_url"`
				PullsURL         string      `json:"pulls_url"`
				MilestonesURL    string      `json:"milestones_url"`
				NotificationsURL string      `json:"notifications_url"`
				LabelsURL        string      `json:"labels_url"`
				ReleasesURL      string      `json:"releases_url"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				PushedAt         time.Time   `json:"pushed_at"`
				GitURL           string      `json:"git_url"`
				SSHURL           string      `json:"ssh_url"`
				CloneURL         string      `json:"clone_url"`
				SvnURL           string      `json:"svn_url"`
				Homepage         interface{} `json:"homepage"`
				Size             int         `json:"size"`
				StargazersCount  int         `json:"stargazers_count"`
				WatchersCount    int         `json:"watchers_count"`
				Language         interface{} `json:"language"`
				HasIssues        bool        `json:"has_issues"`
				HasDownloads     bool        `json:"has_downloads"`
				HasWiki          bool        `json:"has_wiki"`
				HasPages         bool        `json:"has_pages"`
				ForksCount       int         `json:"forks_count"`
				MirrorURL        interface{} `json:"mirror_url"`
				OpenIssuesCount  int         `json:"open_issues_count"`
				Forks            int         `json:"forks"`
				OpenIssues       int         `json:"open_issues"`
				Watchers         int         `json:"watchers"`
				DefaultBranch    string      `json:"default_branch"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
					Login             string `json:"login"`
					ID                int    `json:"id"`
					AvatarURL         string `json:"avatar_url"`
					GravatarID        string `json:"gravatar_id"`
					URL               string `json:"url"`
					HTMLURL           string `json:"html_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					OrganizationsURL  string `json:"organizations_url"`
					ReposURL          string `json:"repos_url"`
					EventsURL         string `json:"events_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"owner"`
				Private          bool        `json:"private"`
				HTMLURL          string      `json:"html_url"`
				Description      string      `json:"description"`
				Fork             bool        `json:"fork"`
				URL              string      `json:"url"`
				ForksURL         string      `json:"forks_url"`
				KeysURL          string      `json:"keys_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				TeamsURL         string      `json:"teams_url"`
				HooksURL         string      `json:"hooks_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				EventsURL        string      `json:"events_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BranchesURL      string      `json:"branches_url"`
				TagsURL          string      `json:"tags_url"`
				BlobsURL         string      `json:"blobs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				TreesURL         string      `json:"trees_url"`
				StatusesURL      string      `json:"statuses_url"`
				LanguagesURL     string      `json:"languages_url"`
				StargazersURL    string      `json:"stargazers_url"`
				ContributorsURL  string      `json:"contributors_url"`
				SubscribersURL   string      `json:"subscribers_url"`
				SubscriptionURL  string      `json:"subscription_url"`
				CommitsURL       string      `json:"commits_url"`
				GitCommitsURL    string      `json:"git_commits_url"`
				CommentsURL      string      `json:"comments_url"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				ContentsURL      string      `json:"contents_url"`
				CompareURL       string      `json:"compare_url"`
				MergesURL        string      `json:"merges_url"`
				ArchiveURL       string      `json:"archive_url"`
				DownloadsURL     string      `json:"downloads_url"`
				IssuesURL        string      `json:"issues_url"`
				PullsURL         string      `json:"pulls_url"`
				MilestonesURL    string      `json:"milestones_url"`
				NotificationsURL string      `json:"notifications_url"`
				LabelsURL        string      `json:"labels_url"`
				ReleasesURL      string      `json:"releases_url"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				PushedAt         time.Time   `json:"pushed_at"`
				GitURL           string      `json:"git_url"`
				SSHURL           string      `json:"ssh_url"`
				CloneURL         string      `json:"clone_url"`
				SvnURL           string      `json:"svn_url"`
				Homepage         interface{} `json:"homepage"`
				Size             int         `json:"size"`
				StargazersCount  int         `json:"stargazers_count"`
				WatchersCount    int         `json:"watchers_count"`
				Language         interface{} `json:"language"`
				HasIssues        bool        `json:"has_issues"`
				HasDownloads     bool        `json:"has_downloads"`
				HasWiki          bool        `json:"has_wiki"`
				HasPages         bool        `json:"has_pages"`
				ForksCount       int         `json:"forks_count"`
				MirrorURL        interface{} `json:"mirror_url"`
				OpenIssuesCount  int         `json:"open_issues_count"`
				Forks            int         `json:"forks"`
				OpenIssues       int         `json:"open_issues"`
				Watchers         int         `json:"watchers"`
				DefaultBranch    string      `json:"default_branch"`
			} `json:"repo"`
		} `json:"base"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		Merged         bool        `json:"merged"`
		Mergeable      interface{} `json:"mergeable"`
		MergeableState string      `json:"mergeable_state"`
		MergedBy       interface{} `json:"merged_by"`
		Comments       int         `json:"comments"`
		ReviewComments int         `json:"review_comments"`
		Commits        int         `json:"commits"`
		Additions      int         `json:"additions"`
		Deletions      int         `json:"deletions"`
		ChangedFiles   int         `json:"changed_files"`
	} `json:"pull_request"`
	Repository struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		Private          bool        `json:"private"`
		HTMLURL          string      `json:"html_url"`
		Description      string      `json:"description"`
		Fork             bool        `json:"fork"`
		URL              string      `json:"url"`
		ForksURL         string      `json:"forks_url"`
		KeysURL          string      `json:"keys_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		TeamsURL         string      `json:"teams_url"`
		HooksURL         string      `json:"hooks_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BranchesURL      string      `json:"branches_url"`
		TagsURL          string      `json:"tags_url"`
		BlobsURL         string      `json:"blobs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		TreesURL         string      `json:"trees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		StargazersURL    string      `json:"stargazers_url"`
		ContributorsURL  string      `json:"contributors_url"`
		SubscribersURL   string      `json:"subscribers_url"`
		SubscriptionURL  string      `json:"subscription_url"`
		CommitsURL       string      `json:"commits_url"`
		GitCommitsURL    string      `json:"git_commits_url"`
		CommentsURL      string      `json:"comments_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		ContentsURL      string      `json:"contents_url"`
		CompareURL       string      `json:"compare_url"`
		MergesURL        string      `json:"merges_url"`
		ArchiveURL       string      `json:"archive_url"`
		DownloadsURL     string      `json:"downloads_url"`
		IssuesURL        string      `json:"issues_url"`
		PullsURL         string      `json:"pulls_url"`
		MilestonesURL    string      `json:"milestones_url"`
		NotificationsURL string      `json:"notifications_url"`
		LabelsURL        string      `json:"labels_url"`
		ReleasesURL      string      `json:"releases_url"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         time.Time   `json:"pushed_at"`
		GitURL           string      `json:"git_url"`
		SSHURL           string      `json:"ssh_url"`
		CloneURL         string      `json:"clone_url"`
		SvnURL           string      `json:"svn_url"`
		Homepage         interface{} `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         interface{} `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorURL        interface{} `json:"mirror_url"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		Forks            int         `json:"forks"`
		OpenIssues       int         `json:"open_issues"`
		Watchers         int         `json:"watchers"`
		DefaultBranch    string      `json:"default_branch"`
	} `json:"repository"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}

//Merged マージした後にfilelistを取ってくるよう
type Merged struct {
	Sha         string `json:"sha"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	Changes     int    `json:"changes"`
	BlobURL     string `json:"blob_url"`
	RawURL      string `json:"raw_url"`
	ContentsURL string `json:"contents_url"`
	Patch       string `json:"patch"`
}

type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type Issue struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels []struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Color   string `json:"color"`
		Default bool   `json:"default"`
	} `json:"labels"`
	State       string        `json:"state"`
	Locked      bool          `json:"locked"`
	Assignee    interface{}   `json:"assignee"`
	Assignees   []interface{} `json:"assignees"`
	Milestone   interface{}   `json:"milestone"`
	Comments    int           `json:"comments"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	ClosedAt    time.Time     `json:"closed_at"`
	PullRequest struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
	} `json:"pull_request"`
	Body     string `json:"body"`
	ClosedBy struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"closed_by"`
}

func GetMergedFileList(accountname string, reponame string, numSTR string, token string) []Merged {

	//ファイルリストのURL
	getURL := "https://api.github.com/repos/" + accountname + "/" + reponame + "/pulls/" + numSTR + "/files"

	req, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		// handle err
	}

	token = "token " + token
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Print Body")
	fmt.Println(string(body))

	var merged []Merged
	err = json.Unmarshal(body, &merged)
	if err != nil {
		log.Fatal(err)
	}

	return merged
}

//GetContent is benri kansuu
//githubのマージ後に叩いたファイルリストAPIのレスポンスからcontents_urlを使って
//downloadリンクを取得、リンクURLを返す
func GetContent(accountname string, reponame string, contentURL string, token string) Content {

	req, err := http.NewRequest("GET", contentURL, nil)
	if err != nil {
		// handle err
	}

	token = "token " + token
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Print Body")
	fmt.Println(string(body))

	var content Content
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

//DownloadGithubRawFile は渡されたGithubのファイルダウンロードリンクと
//トークンを使ってダウンロードして、ファイルパスを返す
//一応ダウンロード先も指定できるようにしてある
func DownloadGithubRawFile(contentURL string, token string, downloadDir string, fileName string) string {
	//実ファイルのダウンロード
	req, err := http.NewRequest("GET", contentURL, nil)
	if err != nil {
		// handle err
	}

	token = "token " + token
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	article, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filePath := downloadDir + "/" + fileName
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	file.Write(article)
	file.Close()

	return filePath
}

//GetCategoryList このソフトウェアはGithubのラベル＝ブログ記事のカテゴリとして扱う機能を
//有したいと思っていたので実装
func GetCategoryList(accountName string, numSTR string, repoName string, token string) []string {

	url := "https://api.github.com/repos/" + accountName + "/" + repoName + "/issues/" + numSTR

	varDump(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}

	token = "token " + token
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Get Label")
	fmt.Println(string(body))

	var issue Issue
	err = json.Unmarshal(body, &issue)
	if err != nil {
		log.Fatal(err)
	}

	varDump(issue)

	var labels []string
	for i := 0; i < len(issue.Labels); i++ {
		fmt.Println(issue.Labels[i].Name)
		labels = append(labels, issue.Labels[i].Name)

	}

	return labels
}
