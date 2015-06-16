package crawler

import (
	"strings"
	"fmt"
	"github.com/chrismclaughlin55/list"
	"github.com/chrismclaughlin55/crawler/neoQuery"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/jmcvetta/neoism"
)

type Crwler struct {
	u   *url.URL
	queue *list.ArrayList
	db *neoism.Database
}

func (c *Crwler) Init(u string) {
	var err error
	c.db,err = neoism.Connect("http://localhost:7474/db/data")	
	if err != nil {
		panic(fmt.Sprintf("shit went bad: %s\n",err))
	}
	c.u,_ = url.Parse(u)
	c.queue = new(list.ArrayList)
	c.queue.Init()
}

func (c *Crwler) Start() {
	c.Shell()
	for {
	 	_ = c.nextURL()
		c.Parse()
	}
}

func (c *Crwler) Shell() {
	fmt.Println("Shell Started")

}

func (c *Crwler) Parse() {
	reader, err := GetHTML(c.u.String())
	if err != nil {
		fmt.Printf("error getting html: %s", err)
		return
	}
	tokenizer := html.NewTokenizer(reader)
	for token:=tokenizer.Next(); token != html.ErrorToken; token =tokenizer.Next(){
		if token == html.StartTagToken{
			tag, val, _ := tokenizer.TagAttr()
			if string(tag) == "href" {
				fmt.Printf("%s\n", val)
				u := string(val)
				if strings.Contains(string(val), "http") {
					link,_ := url.Parse(u)
					c.queueURL(link)
				}
			}
		}
	}
}

func (c *Crwler) nextURL()  error {
	data, err := c.queue.Pop()
	if err != nil {
		return  err
	}
	c.u,_ = data.(*url.URL)
	return  nil

}

func (c *Crwler) queueURL(url *url.URL) {
	if c.queue.Length() <= 9 {
		c.queue.Push(url)
	}
	err := neoQuery.InsertUrl(c.db, c.u.String(), url.String())
	if err != nil {
		fmt.Printf("error inserting into neo4j:%s\n", err)
	}
}

func GetHTML(url string) (*strings.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR in getHTML: %v\n", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, err
	}
	r := strings.NewReader(string(body))
	return r, nil
}
