package tango

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"appengine/urlfetch"
	"net/http"
	"strings"
	"fmt"
	"encoding/xml"
)

// データストアのデータ型
type Entity struct {
	UserID string `json:"-"`
	Word string
	Meaning string
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
	var entities []Entity
	var meaning string
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	word = r.FormValue("word")
	meaning = r.FormValue("meaning")
	
	// 単語一覧の読み込み
	key = datastore.NewIncompleteKey(c, "words", nil)
	entity = new(Entity)
	datastore.Get(c, key, entity)

	// 意味が空白なら自動翻訳
	if meaning == "" {
		targetURL = fmt.Sprintf("https://api.datamarket.azure.com/Bing/MicrosoftTranslator/v1/Translate?Text=%%27%s%%27&To=%%27ja%%27", word)
		request, err = http.NewRequest("GET", targetURL, nil)
		Check(c, err)
		request.SetBasicAuth("", "BY/r96i694uamK+xuSv/6PrzIkfjraA1XFXIhzJ/4tE=")
		client = urlfetch.Client(c)
		resp, err = client.Do(request)
		Check(c, err)
		responseXML = make([]byte, 2048)
		_, err = resp.Body.Read(responseXML)
		Check(c, err)
		
		// xml解析
		result = Result{Text: "none"}
		err = xml.Unmarshal(responseXML, &result)
		Check(c, err)
		meaning = result.Text
	}

	// 単語を追加
	entity.UserID = u.ID
	entity.Word = word
	entity.Meaning = meaning
		
	// データストアへの書き込み
	key, err = datastore.Put(c, key, entity)
	Check(c, err)
	
	// 現在の単語数を返す
	entities = get(c, u)
	fmt.Fprintf(w, "{\"wordnum\":%d}", len(entities))
}

/**
 * データストアの単語をすべて削除する
 * ajaxから呼び出すためのAPI
 */
func clear(w http.ResponseWriter, r *http.Request) {
	var c appengine.Context
	var u *user.User
	var keys []*datastore.Key
	var err error
	var query *datastore.Query
	var count int
	var entities []Entity
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	query = datastore.NewQuery("words").Filter("UserID =", u.ID)
	count, err = query.Count(c)
	Check(c, err)
	entities = make([]Entity, count)
	keys, err = query.GetAll(c, &entities)
	Check(c, err)
	err = datastore.DeleteMulti(c, keys)
	Check(c, err)
}

/**
 * データストアから指定された単語を削除する
 * API
 */
func delete(w http.ResponseWriter, r *http.Request) {
	var c appengine.Context
	var u *user.User
	var query *datastore.Query
	var word string
	var iterator *datastore.Iterator
	var key *datastore.Key
	var err error
	var entity Entity
	
	word = r.FormValue("word")
	if(word == "") {
		return
	}

	c = appengine.NewContext(r)
	u = user.Current(c)
	query = datastore.NewQuery("words").Filter("UserID =", u.ID).Filter("Word =", word)
	iterator = query.Run(c)
	
	for ; ; {
		key, err = iterator.Next(&entity)
		if err != nil {
			break
		}
		err = datastore.Delete(c, key)
		Check(c, err)
	}
}

/**
 * すべての単語リストを取得する
 */
func get(c appengine.Context, u *user.User) []Entity {
	var result []Entity
	var query *datastore.Query
	var err error
	var iterator *datastore.Iterator
	var entity *Entity
	
	query = datastore.NewQuery("words").Filter("UserID =", u.ID)
	Check(c, err)
	
	err = nil
	result = make([]Entity, 0)
	entity = new(Entity)
	for iterator = query.Run(c); ; {
		_, err = iterator.Next(entity)
		if err != nil {
			break
		}
		result = append(result, *entity)
	}
	
	return result
}

/**
 * 単語リストをHTML形式で返す
 * サーバから使用する
 */
func getWordsHTML(c appengine.Context, u *user.User) string {
	var result string
	var entities []Entity
	var li string
	
	result = ""
	entities = get(c, u)
	for i := 0; i < len(entities); i++ {
		li = fmt.Sprintf("<li>%s</li>", entities[i].Word)
		result = strings.Join([]string{result, li}, "")
	}
	
	return result
}