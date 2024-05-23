package httpReq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"report/dao"
)

func LoginItms() (loginCookie *http.Cookie, err error) {
	email := os.Getenv("itms_email")
	password := os.Getenv("itms_password")
	loginJsonStr := `{"email":"` + email + `","password":"` + password + `"}`
	loginResp, err := http.Post(local+loginUrl, "application/json", bytes.NewReader([]byte(loginJsonStr)))
	if err != nil {
		return nil, fmt.Errorf("login http.Post() err: %v", err)
	}
	for _, cookie := range loginResp.Cookies() {
		loginCookie = cookie
	}
	return loginCookie, err
}

func ListTodayJob(loginCookie *http.Cookie) (listJobReply dao.ListJobReply, err error) {
	listJobReq, err := http.NewRequest("PUT", local+listRephrasedJob, nil)
	if err != nil {
		log.Fatalf("http.NewRequest() err: %v", err)
	}
	listJobReq.AddCookie(loginCookie)
	client := &http.Client{}
	listJobResp, err := client.Do(listJobReq)
	if err != nil {
		log.Fatalf("client.Do(req) err: %v", err)
	}
	listJobRespByte, err := io.ReadAll(listJobResp.Body)
	if err != nil {
		log.Fatalf("io.ReadAll(listJobResp.Body) err: %v", err)
	}
	err = json.Unmarshal(listJobRespByte, &listJobReply)
	return listJobReply, err
}

func GetJob(loginCookie *http.Cookie, id string) (getJobReply dao.GetJobReply, err error) {
	getJobReq, err := http.NewRequest("GET", local+getRephrasedJobUrl+id, nil)
	if err != nil {
		log.Fatalf("http.NewRequest() err: %v", err)
	}
	getJobReq.AddCookie(loginCookie)
	client := &http.Client{}
	getJobResp, err := client.Do(getJobReq)
	if err != nil {
		log.Fatalf("client.Do(req) err: %v", err)
	}
	getJobRespByte, err := io.ReadAll(getJobResp.Body)
	if err != nil {
		log.Fatalf("io.ReadAll(getJobResp.Body) err: %v", err)
	}
	err = json.Unmarshal(getJobRespByte, &getJobReply)
	if err != nil {
		log.Fatalf("json.Unmarshal(getJobRespByte, &getJobReply) err: %v", err)
	}
	return getJobReply, err
}

func ListRephrasedJob(loginCookie *http.Cookie, startDate, endDate string) (listJobReply dao.ListJobReply, err error) {
	request := `{"offset":0,"limit":99999,"filter":{"startDate":"` + startDate + `","endDate":"` + endDate + `"}}`
	fmt.Println("request:", request)
	listJobReq, err := http.NewRequest("PUT", local+listRephrasedJob, bytes.NewReader([]byte(request)))
	if err != nil {
		log.Fatalf("http.NewRequest() err: %v", err)
	}
	listJobReq.AddCookie(loginCookie)
	client := &http.Client{}
	listJobResp, err := client.Do(listJobReq)
	if err != nil {
		log.Fatalf("client.Do(req) err: %v", err)
	}
	listJobRespByte, err := io.ReadAll(listJobResp.Body)
	if err != nil {
		log.Fatalf("io.ReadAll(listJobResp.Body) err: %v", err)
	}
	err = json.Unmarshal(listJobRespByte, &listJobReply)
	return listJobReply, err
}
