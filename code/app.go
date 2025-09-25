package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) MakeFile() {
	filename := p.Title + ".txt"
	file, _ := os.Create(filename)
	file.Write(p.Body)
}

func Read(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/Read/")]
	file, _ := os.ReadFile(string(filename) + ".txt")
	p := &Page{Title: string(filename), Body: file}
	temp, _ := template.ParseFiles("Read.html")
	temp.Execute(w, p)
}
func save(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/Save/")]
	p1 := &Page{Title: string(filename), Body: []byte("complete")}
	p1.MakeFile()
	temp, _ := template.ParseFiles("Save.html")
	temp.Execute(w, p1)
}
func main() {
	log.Fatal(http.ListenAndServe(":8000", nil))
	http.HandleFunc("/Save/", save)
	http.HandleFunc("/Read/", Read)
}
