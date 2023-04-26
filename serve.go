package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type Credit struct {
	sum     int
	age     int
	percent int
}

func raschetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/valid.html")
}

func itogHandler(w http.ResponseWriter, r *http.Request) {
	Sum := r.FormValue("sum")
	Age := r.FormValue("age")
	Percent := r.FormValue("percent")
	sum, _ := strconv.Atoi(Sum)
	age, _ := strconv.Atoi(Age)
	percent, _ := strconv.Atoi(Percent)
	month := float64(sum) * (float64(percent/100/12) + float64(percent/100/12)/math.Pow(float64(1+(percent/100/12)), float64(age)) - 1.0)
	result := month*float64(age) - float64(sum)
	temp, _ := template.New("result").Parse("<h1>{{ .}}</h1>")
	temp.Execute(w, result)
	fmt.Fprintln(w, "ВАША ПЕРЕПЛАТА ПО КРЕДИТУ")
}

// Переплата = Ежемесячный платеж X Количество месяцев — Сумма основного долга.
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/user.html")
	})
	http.HandleFunc("/postform", raschetHandler)
	http.HandleFunc("/result", itogHandler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
