package crawler



import (
"github.com/chrismclaughlin55/list"
"github.com/jmcvetta/neoism"
"net/url"
"fmt"


)



type Crawler struct {
	u	*url.URL
	db	*neoism.Database
	res	list.Lister

}

func CrawlerFactory () *Crawler {
/*
This func creates a new crawler struct and initializes the db connection and the list
*/
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		panic(fmt.Sprintf("shit went with connect db: %s\n",err))

	}
	list := new(list.LinkedList)
	list.Init()
	c := &Crawler{db: db, res: list}

	return c
}

func ( self *Crawler) nextURL() error {
/*
 This func queries the DB for a leaf node that needs to be parsed
*/


}

func ( self *Crawler) parse() {
/*
 This func Parses a webpage. To be used within the ProcessURL func
*/

}
