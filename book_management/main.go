package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "strconv"


    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tmpl *template.Template


func init() {
    var err error
    tmpl = template.Must(template.ParseGlob("templates/*.html"))
    db, err = sql.Open("mysql", "root:Djm@1211@tcp(localhost:3306)/booksDB")
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/add-book", addBookHandler)
    http.HandleFunc("/delete-book", deleteBookHandler)

    log.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    books, err := getBooks()
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.ExecuteTemplate(w, "index.html", books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    title := r.FormValue("title")
    author := r.FormValue("author")
    publishedDate := r.FormValue("published_date")

    _, err := db.Exec("INSERT INTO books (title, author, published_date) VALUES (?, ?, ?)", title, author, publishedDate)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    id, err := strconv.Atoi(r.FormValue("id"))
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    _, err = db.Exec("DELETE FROM books WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}


func getBooks() ([]Book, error) {
    rows, err := db.Query("SELECT id, title, author, published_date FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []Book
    for rows.Next() {
        var b Book
        err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate)
        if err != nil {
            log.Println(err)
            continue
        }
        books = append(books, b)
    }
    return books, nil
}
