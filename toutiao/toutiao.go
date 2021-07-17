package toutiao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

const (
	url = "https://toutiao.io"
)

var log *zap.SugaredLogger
var result string

type toutiao interface {
	GetNewData() string
}

type Toutiao struct {
}

var _ toutiao = &Toutiao{}

func (t *Toutiao) GetNewData() string {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Mobile Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Infof("%v", err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Infof("%v", err)
	}
	doc.Find(".daily .posts .post").Each(func(i int, s *goquery.Selection) {
		a := s.Find(".content .title a")
		title, _ := a.Attr("title")
		hreftmp, _ := a.Attr("href")
		href := url + hreftmp
		result += fmt.Sprintf("[%s](%s)\n\n", title, href)
	})
	return result
}

func Daily() {
	t := &Toutiao{}
	local, _ := time.LoadLocation("Asia/Shanghai")
	fileName := time.Now().In(local).Format("2006年01月02日")
	content := `# %s
## %s
%s
`
	_ = ioutil.WriteFile(fmt.Sprintf("./docs/%s.md", fileName), []byte(fmt.Sprintf(content, fileName, "头条", t.GetNewData())), 0644)
}

func init() {
	logger, err := zap.NewProduction()
	log = logger.Sugar()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.Sync()

	}()
}
