package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	/**
	not found route untuk path undefined
	*/
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to home")) // plain text

	// render html page
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	// dummy data map
	data := map[string]string{
		"title":   "Golang web basic",
		"content": "I am learning golang web with agung setiawan",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar Golang web"))
}

func SuhadaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("suhada, this is awesome"))
}

// handler untuk query
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	/**
	validasi query id
	*/
	idNumb, err := strconv.Atoi(id) // convert string ke integer
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("product page"))
	// fmt.Fprint(w, "Product Page : ", idNumb)

	// dummy data interface
	// data := map[string]interface{}{
	// 	"content": idNumb,
	// }

	// dummy 1 data struct
	// data := entity.Product{
	// 	ID:    1,
	// 	Name:  "Mobilio",
	// 	Price: 220000000,
	// 	Stock: 3,
	// }

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 270000000, Stock: 2},
		{ID: 3, Name: "Fortuner", Price: 550000000, Stock: 1},
	}

	tmp, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
		return
	}

	log.Println(err)

}

// hendle http method
func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method // GET  POST
	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is hapenning, keep calming", http.StatusBadRequest)
	}
}

// Form handler
func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Error is hapenning, keep calm", http.StatusBadRequest)
}

func ProcessPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}

		// extract form
		name := r.Form.Get("name")
		message := r.Form.Get("message")
		data := map[string]interface{}{
			"Name":    name,
			"Message": message,
		}

		log.Println(name, message)
		// w.Write([]byte(name))

		tmp, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmp.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is hapenning, keep calm", http.StatusInternalServerError)
			return
		}

		log.Println(err)

		return

	}
	http.Error(w, "Error is hapenning, keep calm", http.StatusBadRequest)

}
