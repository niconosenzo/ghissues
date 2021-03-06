package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niconosenzo/ghissues"
)

func main() {
	result, err := ghissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		//fmt.Printf("#%-5d %9.9v %9.9s %.55s %.55s\n",
		fmt.Printf("#%d %s %s %s %s\n",
			item.Number, item.CreatedAt.Format("Mon Jan 2 2006"), item.User.Login, item.Title, item.HTMLURL)
	}
}
