package tango

import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"appengine/user"
)

/**
 * データストアに単語を追加する
 */
func add(w http.ResponseWriter, r *http.Request) {
	type Entity struct {
		Words []string
	}
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