package tango

import (
	"net/http"
	"appengine"
)

func init() {
	http.HandleFunc("/", top)
	http.HandleFunc("/add", add)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/clear", clear)
	http.HandleFunc("/play", play)
}

/*
	関数 Check(c, err)
	- エラーが発生していたらコンソールへ出力する
	
	引数
	- c : コンソール出力用コンテキスト
	- err : error型　他の関数から渡されるエラー変数
	
	戻り値
	- なし
*/
func Check(c appengine.Context, err error) {
	if err != nil {
		c.Errorf(err.Error())
	}
}