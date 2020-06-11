package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	_ "martialscans"
)

func main() {
	chapters, err := martialscans.retriveChapters("41")
	if err != nil {
		logrus.Fatalln(err)
	}
	for _, element := range chapters {
		fmt.Println(element.Title)
		links, err := retriveImages(element.URL)
		if err != nil {
			logrus.Fatalln(err)
		}
		fmt.Println(links)
	}
}
