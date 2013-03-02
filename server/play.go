package tango

import (
	"net/http"
	"text/template"
	"appengine"
	"appengine/user"
	"strings"
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
	var words []string

	c = appengine.NewContext(r)
	u = user.Current(c)
	
	// 単語一覧を取得
	contents = new(Contents)
	words = get(c, u)
	
	// ページテンプレート取得
	page, err = template.ParseFiles("server/play.html")
	Check(c, err)
	
	// 単語一覧をHTML内に書き込む
	for i := 0; i < len(words); i++ {
		words[i] = strings.Join([]string{"\"", words[i], "\""}, "")
	}
	contents.Words = strings.Join(words, ",")
	
	// 表示
	page.Execute(w, contents)
}