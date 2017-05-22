package main

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "encoding/json"
    // "net/url"
)

type Page struct {
    Name string
    DBStatus bool
}

type SearchResult struct {
    Title string
    Author string
    Year string
    ID string
}

func main() {
    // Must absorbes the error from parsed files and halt execution of program
    templates := template.Must(template.ParseFiles("templates/index.html"))

    db, _ := sql.Open("sqlite3", "dev.db")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        p := Page{Name: "Gopher"}

        if name := r.FormValue("name"); name != "" {
            p.Name = name
        }
        p.DBStatus = db.Ping() == nil

        if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
        results := []SearchResult{
            SearchResult{"The Beach", "Alex Garland", "1998", "234233"},
            SearchResult{"1984", "George Orwell", "1948", "213236"},
            SearchResult{"Trainspotting", "Irvine Welsh", "1994", "234662"},
            SearchResult{"Quiet", "Susan Cain", "2009", "231453"},
        }

        encoder := json.NewEncoder(w)
        if err := encoder.Encode(results); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })
    // db.Close()
    fmt.Println(http.ListenAndServe(":8080", nil))
}

// type ClassifySearchResponse struct {
//     Results []SearchResult `xml:"works>work"`
// }

// func search(query string) ([]SearchResult, error) {
//     var resp *http.Response
//     var err error
//     baseUrl := "http://classify.oclc.org/classify2/Classify?&summary=true&title="

//     if resp, err = http.Get(baseUrl + url.QueryEscape(query)); err != nil {
//         return []SearchResult{}, err
//     }

//     defer resp.Body.Close()
//     var body []byte
//     if body, err = ioutil.ReadAll(resp.Body); err != nil {
//         return []SearchResult{}, err
//     }
// }
