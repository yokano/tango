package tango

import (
	"net/http"
	"appengine"
	"appengine/user"
	"text/template"
)

/**
 * トップページの表示
 *　ログインしている場合としていない場合で異なるファイルを出力
 */
func top(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	
	// ログインしていない
	if u == nil {
		contents := make(map[string]string)
		contents["LOGIN_URL"], _ = user.LoginURL(c, "")
		t, _ := template.ParseFiles("server/login.html")
		t.Execute(w, contents)
	} else {
		// ログインしている
		contents := make(map[string]string)
		contents["ID"] = u.ID
		contents["NAME"] = u.Email
		contents["LOGOUT_URL"], _ = user.LogoutURL(c, "")
		contents["Words"] = getWordsHTML(c, u)
		
		t, _ := template.ParseFiles("server/home.html")
		t.Execute(w, contents)
	}
}
