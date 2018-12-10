package logic

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"
)

var urls = []string{
	"http://api.weibo.com/2/short_url/shorten.json?source=31641035&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=1905839263&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=783190658&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2702428363&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=82966982&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=3105114937&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=1905839263&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=569452181&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2702428363&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=31024382&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=783190658&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2735371158&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=1965726745&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2027761570&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2323547071&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2135576995&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2612767607&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=915345515&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=4229079448&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=1078446352&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=1262673699&url_long=%s",
	"http://api.weibo.com/2/short_url/shorten.json?source=2849184197&url_long=%s",
}

type RelObj struct {
	Urls []Rel `json:"urls"`
}
type Rel struct {
	Result     bool   `json:"result"`
	UrlShort   string `json:"url_short"`
	UrlLong    string `json:"url_long"`
	ObjectType string `json:"object_type"`
	Type       int    `json:"type"`
	ObjectId   string `json:"object_id"`
}

func Short(longUrl string) string { // get三次release失败说明代理阵亡
	addr := urls[rand.Intn(len(urls))]
	addr = fmt.Sprintf(addr, url.QueryEscape(longUrl))
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(1 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*1)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	resp, err := c.Get(addr)
	if err == nil { // 提交异常,返回错误
		body, _ := ioutil.ReadAll(resp.Body)
		//glog.Info(string(body), err)
		rel := &RelObj{}
		err = json.Unmarshal(body, rel)
		//glog.Info(rel, err)
		if len(rel.Urls) > 0{
			longUrl =rel.Urls[0].UrlShort
		}

	} else {
		glog.Errorf("[error] short err is %v", err)
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return longUrl
}

func Urlshort(c echo.Context) error { // 开启心跳
	longUrl := c.QueryParam("url")
	//glog.Info(longUrl)
	return c.String(http.StatusOK, Short(longUrl))
}
