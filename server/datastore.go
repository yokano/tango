package tango

import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"strings"
	"fmt"
	"log"
	"encoding/xml"
)

// データストアのデータ型
type Entity struct {
	Words []string
}

/**
 * データストアに単語を追加する
 * ajaxから呼び出すためのAPI
 */
func add(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Text string `xml:"entry>content>properties>Text"`
	}

	var c appengine.Context
	var u *user.User
	var key *datastore.Key
	var err error
	var entity *Entity
	var word string
	var resp *http.Response
	var client *http.Client
	var request *http.Request
	var responseXML []byte
	var result Result
	var targetURL string
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	
	// 単語一覧の読み込み
	key = datastore.NewKey(c, "words", u.ID, 0, nil)
	entity = new(Entity)
	datastore.Get(c, key, entity)
	
	// 単語を追加
	word = r.FormValue("word")
	entity.Words = append(entity.Words, word)

	// データストアへの書き込み
	key, err = datastore.Put(c, key, entity)
	Check(c, err)

	// 意味を付加する
	targetURL = fmt.Sprintf("https://api.datamarket.azure.com/Bing/MicrosoftTranslator/v1/Translate?Text=%%27%s%%27&To=%%27ja%%27", word)
	request, err = http.NewRequest("GET", targetURL, nil)
	Check(c, err)
	request.SetBasicAuth("", "BY/r96i694uamK+xuSv/6PrzIkfjraA1XFXIhzJ/4tE=")
	client = new(http.Client)
	resp, err = client.Do(request)
	Check(c, err)
	responseXML = make([]byte, 2048)
	_, err = resp.Body.Read(responseXML)
	Check(c, err)
	
	// xml解析
	result = Result{Text: "none"}
	err = xml.Unmarshal(responseXML, &result)
	Check(c, err)
	log.Printf("%s", result.Text)
	
	// 現在の単語数を返す
	fmt.Fprintf(w, "{\"wordnum\":%d, \"status\":\"%s\"}", len(entity.Words), resp.Status)
}

/**
 * データストアの単語をすべて削除する
 * ajaxから呼び出すためのAPI
 */
func clear(w http.ResponseWriter, r *http.Request) {
	var c appengine.Context
	var u *user.User
	var key *datastore.Key
	var entity *Entity
	var err error
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	
	// 空の単語リストを作成
	entity = new(Entity)
	entity.Words = []string{}
	
	// 上書き
	key = datastore.NewKey(c, "words", u.ID, 0, nil)
	_, err = datastore.Put(c, key, entity)
	Check(c, err)
}

/**
 * すべての単語リストを取得する
 */
func get(c appengine.Context, u *user.User) []string {
	var entity *Entity
	var key *datastore.Key
	
	entity = new(Entity)
	key = datastore.NewKey(c, "words", u.ID, 0, nil)
	datastore.Get(c, key, entity)
	
	return entity.Words
}

/**
 * 単語リストをHTML形式で返す
 * サーバから使用する
 */
func getWordsHTML(c appengine.Context, u *user.User) string {
	var result string
	var key *datastore.Key
	var entity *Entity
	
	// 単語リストの取得
	key = datastore.NewKey(c, "words", u.ID, 0, nil)
	entity = new(Entity)
	datastore.Get(c, key, entity)
	
	// 結合
	result = strings.Join(entity.Words, ", ")
	
	return result
}