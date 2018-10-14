package main
import (
    "net/http"
)
func indexTmpl(rw http.ResponseWriter,name string ,tweets []*Tweet,flush string) {
    w, ok := rw.(http.ResponseWriter);
    if !ok { return }
    w.WriteHeader(http.StatusOK)
    w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
    w.Write([]byte(`<!DOCTYPE html><html><head><title>Isuwitter</title><link rel="stylesheet" href="/css/style.css" /></head><body><header class="header"><a class="title" href="/">Isuwitter</a>`))
    if name != "" {
        w.Write([]byte(`<form class="logout" action="/logout" method="post"><button type="submit">ログアウト</button></form><span class="name">こんにちは`))
        w.Write([]byte(name))
        w.Write([]byte(`さん</span>`))
    } else {
        w.Write([]byte(`<span class="name">こんにちは ゲストさん</span>`))
    }
    w.Write([]byte(`<form class="search" action="/search" method="get"><input type="text" name="q" placeholder="search" /></form></header><div class="container">`))
    if name != "" {
        w.Write([]byte(`<div class="post"><form action="/" method="post"><textarea name="text" cols="50" rows="5"></textarea><button type="submit">投稿</button></form></div>`))
        w.Write([]byte(`<div class="timeline">`))
        for _,tweet := range tweets  {
            w.Write([]byte(`<div class="tweet" data-time="`))
            w.Write([]byte(tweet.Time))
            w.Write([]byte(`"><p><a href="/`))
            w.Write([]byte(tweet.UserName))
            w.Write([]byte(`" class="tweet-user-name">`))
            w.Write([]byte(tweet.UserName))
            w.Write([]byte(`</a></p><p>`))
            w.Write([]byte(tweet.HTML))
            w.Write([]byte(`</p><p class="time">`))
            w.Write([]byte(tweet.Time))
            w.Write([]byte(`</p></div>`))
        }
        w.Write([]byte(`</div><button class="readmore">さらに読み込む</button>`))
    } else {
        if flush != "" {
            w.Write([]byte(`<p class="flush">`))
            w.Write([]byte(flush))
            w.Write([]byte(`</p>`))
        }
        w.Write([]byte(`<form class="login" action="/login" method="post"><input type="text" name="name"><input type="password" name="password"><button type="submit">ログイン</button></form>`))
    }
    w.Write([]byte(`</div><script src="/js/script.js"></script></body></html>`))

}

