package neoQuery

import( 
	"github.com/jmcvetta/neoism"
	"fmt"
)
type myDb interface{ 
	GetOrCreateNode(label string, key string, p neoism.Props) (*neoism.Node, bool, error)
}

func InsertUrl(db myDb, parentURL string, url string) error {
	fmt.Println("wtf")
	parProps := neoism.Props{ "url":parentURL}
	parNode, created, err := db.GetOrCreateNode("Website", "url", parProps)
	if err != nil {
		fmt.Printf("shit went bad: %s\n",err)
		return err 
	}	
	if created {
		fmt.Println("parent node created")
	} else {
		fmt.Println("parent node not created")
	}
	childProps := neoism.Props{ "url":url}
	childNode, created, err := db.GetOrCreateNode("Website", "url", childProps)
	if err != nil {
		fmt.Printf("shit went bad: %s\n",err)
		return err 
	}	
	if created {
		fmt.Println("child node created")
	} else {
		fmt.Println("child node not created")
	}
	rels,_ := parNode.Outgoing()
	for _,rel := range rels {
		checkNode,err := rel.End()
		if err != nil {
			fmt.Println("error getting relationship end")
			
		}
		if equalNodes(checkNode, childNode) {
			return nil
		} 
}

	rel,err := parNode.Relate("LINKED", childNode.Id(),nil)
fmt.Println(rel)	
return nil
	
}

func equalNodes ( node1,node2 *neoism.Node) bool {
	val1,ok1 := node1.Data["url"]
	val2,ok2 := node2.Data["url"]
	if val1 == val2 && ok1 && ok2 {
		return true
	}
	return false
}

