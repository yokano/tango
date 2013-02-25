package tango

import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"strings"
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
	var c appengine.Context
	var u *user.User
	var key *datastore.Key
	var err error
	var entity *Entity
	var word string
	
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
 * 単語リストを取得
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