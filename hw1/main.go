/*
1. На основе `html/template` создать CV(резюме) **only 1 page**
2. CV для нескольких страниц(templates)
3. *Additional*: оформить его с помощью hugo


#### Подсказки
Создать новый сайт
```
hugo new site <site name>
```
Запустить сервер по адресу http://127.0.0.1:1313
```
hugo server
```
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/gorilla/mux"
)

type Essential struct {
	Name      string
	Surname   string
	Patronim  string
	Birthdate time.Time
}
type Additional struct {
	PhoneNumber           string
	Email                 string
	Education             string
	AdditionalInformation string
}
type Data struct {
	sync.Mutex
	*Essential
	*Additional
}

func New() *Data {
	d := &Data{sync.Mutex{}, &Essential{Name: "Ilya", Surname: "Druz", Patronim: "Ivanovich", Birthdate: time.Now()}, &Additional{PhoneNumber: "79852453347", Email: "ivandruz@mail.ru", Education: "MAI", AdditionalInformation: "Programmer"}}
	return d

}

func (d *Data) fillData() {
	d.Lock()
	defer d.Unlock()
	fmt.Printf("Enter name: ")
	fmt.Scanf("%s\n", &d.Name)

	for d.Name == "" || string(d.Name[0]) == strings.ToLower(string(d.Name[0])) {
		fmt.Printf("Error, please enter correct name: ")
		fmt.Scanf("%s\n", &d.Name)
		fmt.Println()
	}

	fmt.Printf("Enter surname: ")

	fmt.Scanf("%s\n", &d.Surname)

	for d.Surname == "" || string(d.Surname[0]) == strings.ToLower(string(d.Surname[0])) {
		fmt.Printf("Error, please enter correct surname: ")
		fmt.Scanf("%s\n", &d.Surname)
		fmt.Println()
	}

	fmt.Printf("Enter patronim: ")
	fmt.Scanf("%s\n", &d.Patronim)
	for d.Patronim == "" || string(d.Patronim[0]) == strings.ToLower(string(d.Patronim[0])) {
		fmt.Printf("Error, please enter correct patronim: ")
		fmt.Scanf("%s\n", &d.Patronim)
		fmt.Println()
	}

	fmt.Printf("Enter birthdate year: ")
	var y, m, day int
	fmt.Scanf("%d\n", &y)
	for y < 1900 || y == time.Now().Year() || y > time.Now().Year() {
		fmt.Printf("Error, enter birthdate year: ")
		fmt.Scanf("%d\n", &y)
		fmt.Println()
	}
	fmt.Printf("Enter birthdate month: ")
	fmt.Scanf("%d\n", &m)
	for m < 0 || m > 12 {
		fmt.Printf("Error, enter birthdate month: ")
		fmt.Scanf("%d\n", &m)
		fmt.Println()
	}
	fmt.Printf("Enter birthdate day: ")
	fmt.Scanf("%d\n", &day)
	for day < 0 || day > 31 {
		fmt.Printf("Error, enter birthdate day: ")
		fmt.Scanf("%d\n", &day)
		fmt.Println()
	}
	d.Birthdate = time.Date(y, time.Month(m), day, 0, 0, 0, 0, time.UTC)

	fmt.Printf("Enter phone number: ")
	fmt.Scanf("%s\n", &d.PhoneNumber)
	for !isValidPhoneNumber(d.PhoneNumber) {
		fmt.Printf("Error, enter valid phone number: ")
		fmt.Scanf("%s\n", &d.PhoneNumber)
		fmt.Println()
	}

	fmt.Printf("Enter email: ")
	fmt.Scanf("%s\n", &d.Email)
	for !isValidEmail(d.Email) {
		fmt.Printf("Error, enter valid email: ")
		fmt.Scanf("%s\n", &d.Email)
		fmt.Println()
	}

	fmt.Printf("Enter education: ")
	fmt.Scanf("%s\n", &d.Education)
	for d.Education == "" {
		fmt.Printf("Error, enter education: ")
		fmt.Scanf("%s\n", &d.Education)
		fmt.Println()
	}

	fmt.Printf("Enter additional information: ")
	fmt.Scanf("%s\n", &d.AdditionalInformation)
	for d.AdditionalInformation == "" {
		fmt.Printf("Error, enter additional information: ")
		fmt.Scanf("%s\n", &d.AdditionalInformation)
		fmt.Println()
	}

}

func isValidPhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) < 10 || len(phoneNumber) > 15 {
		return false
	}
	for _, c := range phoneNumber {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isValidEmail(email string) bool {
	if len(email) < 3 || !strings.Contains(email, "@") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	if !strings.Contains(parts[1], ".") {
		return false
	}
	return true
}

func (d *Data) showAllInformation() {
	fmt.Printf("Name: %s \nSurname: %s \nPatronim: %s\n", d.Name, d.Surname, d.Patronim)
	fmt.Printf("Birthdate: %v\n", d.Birthdate)
	fmt.Printf("Education: %s\n", d.Education)
	fmt.Printf("Email: %s\n", d.Email)
	fmt.Printf("Additional: %s\n", d.AdditionalInformation)
}

func CV(w http.ResponseWriter, r *http.Request) {
	templates["CV"].Execute(w, &Data{sync.Mutex{}, &Essential{Name: "Ilya", Surname: "Druz", Patronim: "Ivanovich", Birthdate: time.Now()}, &Additional{PhoneNumber: "79852453347", Email: "ilyadruz@mail.ru", Education: "MAI", AdditionalInformation: "Programmer"}})
}

var templates = make(map[string]*template.Template, 1)

func loadTemplates() {
	templateNames := [1]string{"CV"}
	for i, v := range templateNames {
		fmt.Println(i, v)
	}
	for index, name := range templateNames {
		t, err := template.ParseFiles("CV.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()

	fmt.Println("Starting main process")
	r := mux.NewRouter()
	//d := New()
	//d.fillData()
	//d.showAllInformation()
	r.HandleFunc("/", CV)
	log.Fatal(http.ListenAndServe("127.0.0.1:1313", r))
}
