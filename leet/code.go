package leet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

const (
	url       = "https://leetcode-cn.com/graphql/"
	todayUrl  = "https://leetcode-cn.com/problems/"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
)

type leet interface {
	GetNewData() string
}

type Leet struct {
}

var _ leet = &Leet{}

func (l *Leet) GetNewData() string {
	b, _ := json.Marshal(struct {
		OperationName string `json:"operationName"`
		Variables     string `json:"variables"`
		Query         string `json:"query"`
	}{
		OperationName: "questionOfToday",
		Variables:     "{}",
		Query: `query questionOfToday {
  todayRecord {
    question {
      title
      titleCn: translatedTitle
      titleSlug
	}
  }
}`,
	})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
	b, _ = ioutil.ReadAll(resp.Body)
	title := gjson.GetBytes(b, "data.todayRecord.0.question.titleCn").String()
	url := todayUrl + gjson.GetBytes(b, "data.todayRecord.0.question.titleSlug").String()
	return fmt.Sprintf("[%s](%s)\n\n", title, url)
}
