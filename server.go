package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

//BlogConfig is はてなだけに対応
type HatenaBlogConfig struct {
	Blogs []struct {
		Blog     string `json:"blog"`
		RepoName string `json:"repo_name"`
		HatenaID string `json:"hatena_id"`
		BlogID   string `json:"blogID"`
		APIKey   string `json:"api_key"`
		Username string `json:"username,omitempty"`
	} `json:"blogs"`
}

type GithubConfig struct {
	Github []struct {
		AccountName string `json:"account_name"`
		RepoName string `json:"repo_name"`
		AccessToken string `json:"access_token"`
	} `json:"github"`
}



func main() {


/*
	////////定数的なもの
	accountName := "j-o-lantern0422"
	//repoName := "test" //切り替えに対応したい
	token := "66faec6dba6db47a71f6a2abf9b17c46fe781071"
*/
	//環境変数から動作環境を取得する
	env := os.Getenv("env")
	varDump(env)
	////////////////

	m := martini.Classic()
	m.Use(render.Renderer())

	//hogehoge.com/apiにJSONがポストされたら動き出す
	m.Post("/api", binding.Json(Post{}), binding.ErrorHandler, func(post Post, r render.Render) {

		//varDump(post)
		action := post.Action
		number := post.Number
		htmlURL := post.PullRequest.HTMLURL
		repoName := post.Repository.Name

		varDump(action)
		varDump(number)
		varDump(htmlURL)
		varDump(repoName)

		//ブログとリポジトリのヒモ付設定を読み込む
		configFile, err := ioutil.ReadFile("./config.json")
		if err != nil {
			panic(err)
		}
		var config HatenaBlogConfig
		json.Unmarshal(configFile, &config)

/*
		var blogIs string

		//Webhookで受け取ったリポジトリ名と同じリポジトリ名を持つ設定を見つける
		for i := 0; i < len(config.Blogs); i++ {
			if config.Blogs[i].RepoName == repoName {
				blogIs = config.Blogs[i].Blog
				blogUserID := config.Blogs[i].HatenaID //変えたい… TODO
				blogID := config.Blogs[i].BlogID
				apiKey := config.Blogs[i].APIKey
			}
		}
		if blogIs == "" {
			fmt.Printf("レポジトリに該当するブログが設定されていません")
			os.Exit(1)
		}
*/
		//Gtihubの設定をよみこむ
		githubCofnig, err := ioutil.ReadFile("./config/githubconfig.json")
		if err != nil {
			panic(err)
		}
		var githubAccount GithubConfig
		json.Unmarshal(githubCofnig, &githubAccount)

		var accountName,token string

		//Webhookで受け取ったリポジトリ名からGithubの設定ファイルを特定する
		for i := 0; i < len(githubAccount.Github); i++ {
			if githubAccount.Github[i].RepoName == repoName {
				accountName = githubAccount.Github[i].AccountName
				token = githubAccount.Github[i].AccessToken
			}
		}

		if accountName == "" || token== ""{
			fmt.Printf("Githubアカウント設定を見なおしてください")
			os.Exit(1)
		}



		//Mergeした時だけ実行される
		if action == "closed" {
			numSTR := strconv.FormatUint(uint64(number), 10)
			varDump(accountName)
			varDump(numSTR)
			varDump(repoName)
			varDump(token)
			var merged []Merged
			merged = GetMergedFileList(accountName, repoName, numSTR, token)

			var labels []string

			labels = GetCategoryList(accountName, numSTR, repoName, token)
			varDump(labels)
			//すべてのファイルにアプローチしていく
			for i := 0; i < len(merged); i++ {
				fmt.Println(merged[i].ContentsURL)

				contentURL := merged[i].ContentsURL

				//contentURLから実ファイルのダウンロード先をもらってくる
				var content Content
				content = GetContent(accountName, repoName, contentURL, token)

				//ダウンロード
				donwloadURL := content.DownloadURL
				downloadDir := "/tmp"
				fileName := content.Path

				filePath := DownloadGithubRawFile(donwloadURL, token, downloadDir, fileName)

				//var blogConfig BlogConfig
				postStatus, err := PostToHatenaBlog(repoName, filePath, labels)

				if err != nil {
					os.Exit(1)
				}

				varDump(postStatus)

				fmt.Print(filePath)

			}

		}

	})

	m.Run()
}

//var_dump 便利関数var_dump
func varDump(v ...interface{}) {
	for _, vv := range v {
		fmt.Printf("%#v\n", vv)
	}
}

//きっと使われることのなくなる関数
func jsonWriter(v ...interface{}) {
	file, err := os.Create(`./hogehoge.txt`)
	if err != nil {
		fmt.Printf("ファイルが開けぬ")
	}
	for _, vv := range v {
		fmt.Printf("%#v\n", vv)
	}
	defer file.Close()
}
