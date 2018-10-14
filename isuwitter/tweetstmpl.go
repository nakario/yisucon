package main
import (
    "net/http"
)
func tweetsTmpl(rw http.ResponseWriter,tweets []*Tweet) {
    w, ok := rw.(http.ResponseWriter);
    if !ok { return }
    w.WriteHeader(http.StatusOK)
    w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
    for _ , tweet := range tweets  {
        w.Write([]byte(` <div class="tweet" data-time="`))
        w.Write([]byte(tweet.Time))
        w.Write([]byte(`"><p><a href="/`))
        w.Write([]byte(tweet.UserName))
        w.Write([]byte(`" class="tweet-user-name">`))
        w.Write([]byte(tweet.UserName))
        w.Write([]byte(`</a></p><p>`))
        w.Write([]byte(tweet.HTML))
        w.Write([]byte(`</p><p class="time">`))
        w.Write([]byte(tweet.Time))
        w.Write([]byte(`</p></div> `))
    }
}

