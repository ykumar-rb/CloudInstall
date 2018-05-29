package handlers

import (
	"github.com/CloudInstall/libhttp"
	"html/template"
	"net/http"
        "fmt"
)

func GetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/create/create.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, "")
}

func GetEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/create/edit.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, "")
}

func ProcessCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/create/submit.html")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, "")
	PrintSubmissionReport(w,r)
//	r.ParseForm()
//	fmt.Fprintln(w, r.Form)

}

func PrintSubmissionReport(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w,"\n#################################################################\n")
	fmt.Fprintln(w,"                    CREATE SUBMISSION REPORT                      \n\n")
	fmt.Fprintf(w,"Created Installation Env:")
	fmt.Fprintln(w,r.Form["installname"])
	fmt.Fprintln(w,"\n\n#################################################################")
}