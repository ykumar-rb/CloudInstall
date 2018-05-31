package handlers

import (
	"github.com/CloudInstall/libhttp"
	"html/template"
	"net/http"
	"fmt"
	"os"
	"io"
	"encoding/json"
	"strings"
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
	macList := r.Form.Get("maclist")
	envName := r.Form.Get("installname")
	autoUpdate := r.Form.Get("Enable Auto Updates")

	mapD := map[string]string{"EnvironmentName": envName, "Mac": macList,"InstructionFileName":filePath,"AutoUpdate":autoUpdate}
	mapB, err := json.Marshal(mapD)
	if err != nil {
		err = fmt.Errorf("error in marshalling the request to json for applying token : %s", err)
		return
	}

	body := strings.NewReader(string(mapB))
	url := "http://restapi3.apiary.io/notes"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		err = fmt.Errorf("error in forming the request: %s ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//sendHTTPPostReqToPNPServer(payloadStr)

}


/*
func sendHTTPPostReqToPNPServer(jsonStr []byte){

	url := "http://restapi3.apiary.io/notes"
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
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
*/