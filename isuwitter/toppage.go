package main

import (
	"net/http"
	"github.com/unrolled/render"
)


func burnOutTopPage(render *render.Render,w http.ResponseWriter, name  string, flush string, Tweets []*Tweet) {
	hw, ok := w.(http.ResponseWriter);
	if !ok {
		return
	}
	hw.WriteHeader(http.StatusOK)
	hw.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
	hw.Write([]byte("<!DOCTYPE html><html>からのぺーじだにょ</html>"))
	// WriteString(hw , "<!DOCTYPE html><html></html>")
	// if !hw.Written(){
	// 	hw.size = 0
	// 	hw.WriteHeader(http.StatusOK)
	// }

	// hw.WriteHeaderNow()
	// h := HTML{
	// 	Head: Head{
	// 		ContentType: render.opt.HTMLContentType + render.compiledCharset,
	// 		Status:      http.StatusOK,
	// 	},
	// 	Name: "index",
	// 	Templates: r.templates,
	// }
	// // e.Render(w, data)
	// return r.Render(w, h, binding)
	// render.HTML(w, , "index", struct {
	// 	Name   string
	// 	Tweets []*Tweet
	// }{
	// 	name, tweets,
	// })


// `<!DOCTYPE html><html><head><title>Isuwitter</title><link rel="stylesheet" href="/css/style.css" /></head><body><header class="header"><a class="title" href="/">Isuwitter</a>
// {{ if .Name }}
//   <form class="logout" action="/logout" method="post"><button type="submit">ログアウト</button></form>
//   <span class="name">こんにちは {{ .Name }}さん</span>
// {{ else }}
//   <span class="name">こんにちは ゲストさん</span>
// {{ end }}
//   <form class="search" action="/search" method="get"><input type="text" name="q" placeholder="search" /></form></header><div class="container">
// {{ if .Name }}
//   <div class="post"><form action="/" method="post"><textarea name="text" cols="50" rows="5"></textarea><button type="submit">投稿</button></form></div><div class="timeline">
// 	{{ range .Tweets }}
//   <div class="tweet" data-time="{{ .Time }}">
//     <p><a href="/{{ .UserName }}" class="tweet-user-name">{{ .UserName }}</a></p>
//     <p>{{ raw .HTML }}</p>
//     <p class="time">{{ .Time }}</p>
//   </div>
//   {{ end }}
// 	</div><button class="readmore">さらに読み込む</button>
// {{ else }}
//   {{ if .Flush }}
//      <p class="flush">{{ .Flush }}</p>
//   {{ end }}
//   <form class="login" action="/login" method="post"><input type="text" name="name"><input type="password" name="password"><button type="submit">ログイン</button></form>
// {{ end }}
// </div><script src="/js/script.js"></script></body></html>`
}