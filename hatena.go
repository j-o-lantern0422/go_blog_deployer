package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//BlogConfig ブログのコンフィグ。複数ブログに対応できるような形にした
type BlogConfig struct {
	Blogs []struct {
		Blog     string `json:"blog"`
		Username string `json:"username"`
		APIKey   string `json:"api_key"`
		HatenaID string `json:"hatena_id"`
		RepoName string `json:"repo_name"`
		BlogID   string `json:"blogID"`
	} `json:"blogs"`
}

/******     はてなブログのエントリー用のXML作成用構造体     ******/
type Author struct {
	Name string `xml:"name"`
}

type HatenaXMLContent struct {
	XMLName      xml.Name `xml:"content"`
	ContentValue string   `xml:",chardata"`
	Type         string   `xml:"type,attr"`
}

type Category struct {
	XMLName xml.Name `xml:"category"`
	Term    string   `xml:"term,attr"`
}

type CategoryList struct {
	Category []Category
}

type App struct {
	XMLName xml.Name `xml:"app:control"`
	Draft   string   `xml:"app:draft"`
}

//Entry is はてなのエントリーXML作成に必要なやつ
type Entry struct {
	XMLName  xml.Name         `xml:"entry"`
	XMLns    string           `xml:"xmlns,attr"`
	XMLnsApp string           `xml:"xmlns:app,attr"`
	Title    string           `xml:"title"`
	Author   Author           `xml:"author"`
	Content  HatenaXMLContent `xml:"content"`
	Updated  string           `xml:"updated"`
	Category []Category       `xml:"category"`
	App      App
}

/******     はてなブログ用Structここまで     ******/

//HatenaConfig 記事作成時に対象のはてなブログの情報だけ保存するためのもの
type HatenaConfig struct {
	Blog     string `json:"blog"`
	Username string `json:"username"`
	APIKey   string `json:"api_key"`
	HatenaID string `json:"hatena_id"`
	RepoName string `json:"repo_name"`
	BlogID   string `json:"blogID"`
}

//CreatePost is Postするためのリクエストをつくる
func CreatePost(articleFile *os.File, author string, categorys []string, publishStatus bool) (postXML string, returnErr error) {
	reader := bufio.NewReader(articleFile)

	// 最初の行はタイトルになる
	title, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	title = strings.TrimRight(title, "\n")

	//1行めをタイトルにしたのでスキップする
	reader.ReadString('\n')

	//３行目以降が実際の記事
	content := ""
	buf := ""
	var fileErr error
	for {
		buf, fileErr = reader.ReadString('\n')
		content += buf
		if fileErr == io.EOF {
			break
		} else if fileErr != nil {
			return "", fileErr
		}
	}

	return CreatePostXML(string(title), author, content, categorys, publishStatus)
}

//CreatePostXML はPostするためのXML(本分とかカテゴリーとかの設定がすべてなされているもの）を作って返す
func CreatePostXML(title string,
	author string,
	contents string,
	categorys []string,
	publishStatus bool) (buildXML string, err error) {
	draftStatusStr := "yes"
	if publishStatus == true {
		draftStatusStr = "no"
	}


	var category []Category

	for i := 0; i < len(categorys); i++ {
		tmp := Category{Term: categorys[i]}
		category = append(category, tmp)
	}

	entry := &Entry{
		XMLns:    "http://www.w3.org/2005/Atom",
		XMLnsApp: "http://www.w3.org/2007/app",
		Title:    title,
		Author:   Author{Name: author},
		Content:  HatenaXMLContent{ContentValue: contents, Type: "text/plain"},
		Category: category,
		App:      App{Draft: draftStatusStr},
	}

	output, marshalErr := xml.MarshalIndent(entry, "", "    ")
	if marshalErr != nil {
		return "", marshalErr
	}

	return xml.Header + string(output), nil
}

//CallHatenaAtomAPI はてなブログのAtomAPIを叩く
func CallHatenaAtomAPI(xml string, config HatenaConfig) error {
	draftPostURL := fmt.Sprintf("%s%s/%s/atom/entry",
		"https://blog.hatena.ne.jp/",
		config.HatenaID,
		config.BlogID)

	req, _ := http.NewRequest(
		"POST",
		draftPostURL,
		bytes.NewBuffer([]byte(xml)))
	req.SetBasicAuth(config.HatenaID, config.APIKey)
	req.Header.Set("Content-Type", "application/atomsvc+xml; charset=utf-8")

	client := new(http.Client)
	res, requestErr := client.Do(req)
	defer res.Body.Close()

	if requestErr != nil {
		return requestErr
	}

	if res.StatusCode != http.StatusCreated {
		fmt.Print("http bad response")
		os.Exit(1)
	}

	return nil
}

//PostToHatenaBlog is はてなブログにポストする時に叩かれるやつ
func PostToHatenaBlog(repoName string, filePath string, category []string) (status int, err error) {

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var blogConfig BlogConfig
	json.Unmarshal(file, &blogConfig)

	for i := 0; i < len(blogConfig.Blogs); i++ {
		fmt.Print(blogConfig.Blogs[i].RepoName)
		if blogConfig.Blogs[i].RepoName == repoName {
			if blogConfig.Blogs[i].Blog != "hatena" {
				fmt.Print("This Blog Configration is not for hatena")
				os.Exit(1)
			}

			articleBody, fperr := os.Open(filePath)
			if fperr != nil {
				// open error
			}

			defer articleBody.Close()

			//postするXMLを作る
			postXML, xmlCreateErr := CreatePost(
				articleBody,
				blogConfig.Blogs[i].HatenaID,
				category,
				false,
			)
			if xmlCreateErr != nil {
				return 1, errors.New("Post XML Create Error")
			}

			//実際にポストする
			apiCallErr := CallHatenaAtomAPI(postXML, blogConfig.Blogs[i])
			if apiCallErr != nil {
				return 1, errors.New("APIコールエラー")

			} else {
				return 0, nil
			}
		}
	}

	return 0, nil

}
