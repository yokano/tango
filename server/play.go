package tango

import (
	"net/http"
	"html/template"
	"appengine"
	"appengine/user"
)

func play(w http.ResponseWriter, r *http.Request) {
	var c appengine.Context
	var u *user.User
	type Contents struct {
		Words []string
	}
	var contents *Contents
	var t *template.Template
	var err error
	
	c = appengine.NewContext(r)
	u = user.Current(c)
	
	// 単語一覧を取得
	contents = new(Contents)
	contents.Words = get(c, u)
	
	t, err = template.ParseFiles("client/play.html")
	Check(c, err)
	t.Execute(w, contents)
}