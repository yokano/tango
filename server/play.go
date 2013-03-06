package tango

import (
	"net/http"
	"text/template"
	"appengine"
	"appengine/user"
	"encoding/json"
)

/**
 * 単語学習ページの表示
 * play.html 学習ページ全体のテンプレート
 */
func play(w http.ResponseWriter, r *http.Request) {
	type Contents struct {
		Words string
	}
	
	var c appengine.Context
	var u *user.User
	var contents *Contents
	var err error
	var page *template.Template
	var entities []Entity
	var bytes []byte
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	
	// 単語一覧を取得
	entities = get(c, u)
	
	// JSONに変換
	bytes, err = json.Marshal(entities)
	Check(c, err)
	contents = new(Contents)
	contents.Words = string(bytes)
	
	// ページテンプレート取得
	page, err = template.ParseFiles("server/play.html")
	Check(c, err)
	
	// 表示
	page.Execute(w, contents)
}