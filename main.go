package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"fmt"
)

func createURL() string{
	u,err := url.Parse("/params")
	if err != nil{
		panic(err)
	}

	u.Host = "localhost:3000"
	u.Scheme = "http"

	query := u.Query() //Nos regresa un map
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()

	return u.String()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Nombre", "Valor del header")
		//fmt.Fprintf(w, "Hola mundo")
		http.Redirect(w,r, "/dos", http.StatusMovedPermanently)
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Nombre", "Valor del header")
		fmt.Fprintf(w, "Hola mundo, dos")
		//http.NotFound(w, r)
		http.Error(w, "Este es un error.", http.StatusConflict)
	})

	http.HandleFunc("/switch", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("El metodo es +" + r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hola Mundo con el metodo get")
		case "POST":
			fmt.Fprintf(w, "Hola mundo con el metodo Post")
		case "PUT":
			fmt.Fprintf(w, "Hola mundo con el metodo Post")
		case "DELETE":
			fmt.Fprintf(w, "Hola mundo con el metodo Post")
		default:
			http.Error(w, "Metodo no valido", http.StatusConflict)
		}
		
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){
		//fmt.Println(r.URL.RawQuery)
		fmt.Println(r.URL.Query()) //Map

		name := r.URL.Query().Get("name")
		if len(name) != 0{
			fmt.Println(name)
		}

		parametro := r.URL.Query().Get("parametro")
		if len(parametro) != 0{
			fmt.Println(parametro)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//r.URL.RawQuery
		//r.URL.Query()

		fmt.Println(r.URL)
		values := r.URL.Query()

		values.Del("otro")
		values.Add("name", "Eduardo")
		values.Add("course", "Go web")
		values.Add("Job", "CodigoFacilito")

		r.URL.RawQuery = values.Encode()
		fmt.Println(r.URL)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//fmt.Println( r.Header )

		accessToken := r.Header.Get("access_token")
		if len(accessToken) != 0 {
			fmt.Println(accessToken)
		}

		r.Header.Set("nombre", "valor")
	})

	// CREATEURL FUNCTION
	url := createURL()
	fmt.Println("LA url final es: " + url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		panic(errr)
	}

	client := &http.Client{}
	client.Do(request)

	if err != nil{
		panic(err)
	}
	request.Header.Set("user","Cody")

	fmt.Println("El header es ", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println("El body es ", string(body))
	fmt.Println("El status es ", response.Status)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}