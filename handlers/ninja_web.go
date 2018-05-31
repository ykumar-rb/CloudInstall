package handlers

import (
	"github.com/CloudInstall/libhttp"
	"html/template"
	"net/http"
	"fmt"
	"os"
	"io"
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
	//PrintSubmissionReport(w,r)


	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	filePath := "./ZTPFiles/"+handler.Filename
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)


	UploadedFileName := handler.Filename
	fmt.Printf("Uploaded file:%s filePath:%s", UploadedFileName, filePath)

	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Printf("********************")
	fmt.Printf(r.Form.Get("maclist"))

}

func PrintSubmissionReport(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, "\n#################################################################\n")
	fmt.Fprintln(w, "                    CREATE SUBMISSION REPORT                      \n\n")
	fmt.Fprintf(w, "Created Installation Env:")
	fmt.Fprintln(w, r.Form["installname"])
	fmt.Fprintln(w, "\n\n#################################################################")

	r.Form.Get("")

	fmt.Printf(r.Form.Get("maclist"))
}
