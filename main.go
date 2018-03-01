package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
	"io/ioutil"
	
	"github.com/url-shortener/models"
	"github.com/url-shortener/database"
	str "github.com/url-shortener/stringhelper"
	_ "github.com/go-sql-driver/mysql"
)

func short(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var entry models.Entry

		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &entry)

		saved, err := entry.Verify()
		// We already have that url persisted, just return it!
		if err == nil {
			json.NewEncoder(w).Encode(saved)
			return
		}

		// Persist a new entry
		insertedID, err := entry.Save()
		if err != nil {
			log.Fatal(err)
		}

		entry.ID = insertedID
		entry.Encoded = str.Encode(insertedID)

		_, err = entry.SaveEncoded()
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(entry)
	} else {
		e := map[string]string{"Message": "Method Not Allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(e)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		code := r.URL.Path[len("/"):]
		var entry models.Entry 

		// no parameters in url, render main page
		if code == "" {
			var templates = template.Must(template.ParseFiles(
				"templates/index.html",
			))
			templates.ExecuteTemplate(w, "index.html", nil)

			return
		}
		
		// Search for the given hash
		entry, err := entry.FindByHash(code)
		
		// There is no hash? Just send a message
		if err != nil {
			e := map[string]string{"Message": "Link não encontrado"}
			json.NewEncoder(w).Encode(e)
			return
		}

		// Found the url by the given hash? Redirect the user
		http.Redirect(w, r, entry.URL, http.StatusFound)

	} else {
		e := map[string]string{"Message": "Method Not Allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(e)
	}
}

func main() {
	db := database.Connect()

	defer db.Close()

	models.Init(db)

	fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/short/", short)
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":3000", nil))
}