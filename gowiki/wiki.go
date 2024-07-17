package gowiki

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func main() {
// p1 := &Page{Title: "Il Cantico delle Balene",
// 	Body: []byte("Il canto delle balene consiste in una serie di suoni emessi dai cetacei per poter comunicare.")}

// p1.save()

// p2, _ := loadPage(p1.Title)
// fmt.Println(string(p2.Body))

// }
