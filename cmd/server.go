package main

import (
        "fmt"
        "html/template"
        "net/http"
)

type WebData struct {
        Title string
        Text  []string
}
var Disp WebData;
func serv(w http.ResponseWriter, r *http.Request) {
        tmpl, _ := template.ParseFiles("index.html")
        tmpl.Execute(w, "")
}
func asc(w http.ResponseWriter, r *http.Request) {
        Disp.Text=Conv(r.FormValue("input"),r.FormValue("font"))
        tmpl1, _ := template.ParseFiles("page1.html")
        tmpl1.Execute(w, Disp)
}
func main() {
        
        //Disp.Title = "ASCII-ART WEB MY FIRST WEB APP"
        //Disp.Text = ;
        //var fontName string
        http.HandleFunc("/", serv)
        http.HandleFunc("/ascii-art", asc)
        fmt.Printf("%s", Disp.Text)
        fmt.Println("Server is listening...")
        http.ListenAndServe(":8181", nil)
}