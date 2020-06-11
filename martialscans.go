package martialscans

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/anaskhan96/soup"
)

type Link struct {
	Title string
	URL   string
}

// links, err := retriveChapters("41")
// if err != nil {
// 	logrus.Fatalln(err)
// }
// for _, element := range links {
// 	fmt.Println(element.Title)
// }

func retriveChapters(manga string) ([]Link, error) {
	var links []Link

	url := "https://martialscans.com/wp-admin/admin-ajax.php"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("action", "manga_get_chapters")
	_ = writer.WriteField("manga", manga)
	err := writer.Close()

	if err != nil {
		return links, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return links, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return links, err
	}
	defer res.Body.Close()

	document := soup.HTMLParse(string(body))

	rc := document.Find("div", "class", "page-content-listing").FindAll("a")
	for _, element := range rc {
		link := Link{
			Title: element.Text(),
			URL:   element.Attrs()["href"],
		}
		links = append(links, link)
	}

	return links, nil
}

// links, err := retriveImages("martial-peak", "574")
// if err != nil {
// 	logrus.Fatalln(err)
// }
// fmt.Println(links)
func retriveImages(chapter string) ([]string, error) {
	var links []string

	//"https://martialscans.com/manhua/martial-peak/chapter-574/?style=list"
	// manhuaURL := "https://martialscans.com/manhua/" + manhua + "/chapter-" + chapter + "/?style=list"
	manhuaURL := chapter + "?style=list"
	response, err := soup.Get(manhuaURL)

	if err != nil {
		return links, err
	}

	document := soup.HTMLParse(response)
	rc := document.Find("div", "class", "reading-content").FindAll("img")
	for _, element := range rc {
		links = append(links, element.Attrs()["data-src"])
	}

	return links, nil
}
