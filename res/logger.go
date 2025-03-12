package res

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type LogEntry struct {
	Id       int
	Request  RequestLog
	Response ResposneLog
}

type pair struct {
	request  *http.Request
	response *http.Response
}

type RequestLog struct {
	Method  string              `json:"Method"`
	Url     url.URL             `json:"Url"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type ResposneLog struct {
	StatusCode int                 `json:"Status Code"`
	Headers    map[string][]string `json:"headers`
	Body       string              `json:"body"`
}

var logs []LogEntry
var shared pair

func OpenLogFile() *os.File {
	v, err := os.OpenFile("Log.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error in Opening the task file: ")
	}
	return v
}

func ReadLogFile(LogFile *os.File, logs []LogEntry) {
	decoder := json.NewDecoder(LogFile)
	err := decoder.Decode(&logs)
	if err != nil {
		fmt.Println("Error in Reading the log file: ", err)
	}
	fmt.Println(logs)
}

func WriteFile(LogFile *os.File) { // function to add the task to the list
	jsondata, errr := json.MarshalIndent(logs, "", "")
	fmt.Println(jsondata)
	if errr != nil {
		fmt.Println("Error in Marshal: ", errr)
	}

	err := os.WriteFile("Loged.json", jsondata, 0644) // overwrites the file again with all the changes
	if err != nil {
		fmt.Println("Error while append: ", err)
		return
	}
}

func setReq(req *http.Request) {
	shared.request = req
	fmt.Println("Shared memory changed")
	return
}

func setResp(resp *http.Response) {
	shared.response = resp
	return
}

func makeLogRequest(req *http.Request) RequestLog {
	body := ""
	if req.Body != nil {
		defer req.Body.Close()
		buf := make([]byte, 1024)
		n, _ := req.Body.Read(buf)
		body = string(buf[:n])
	}
	return RequestLog{req.Method, *req.URL, req.Header, body}
}

func makeResponseLog(resp *http.Response) ResposneLog {
	body := ""
	if resp.Body != nil {
		defer resp.Body.Close()
		buf := make([]byte, 1024)
		n, _ := resp.Body.Read(buf)
		body = string(buf[:n])
	}
	return ResposneLog{resp.StatusCode, resp.Header, body}
}

func Logrequest(req *http.Request) {
	fmt.Println("Logging")
	n := len(logs)
	tempEntry := LogEntry{n + 1, makeLogRequest(req), ResposneLog{0, nil, ""}}
	logs = append(logs, tempEntry)
	fmt.Println(logs)
	fmt.Println("Wrote in the file")
	fmt.Println("Logged the request")
	log.Println("request: ")
	log.Printf("Method; %s\n", req.Method)
	log.Printf("URL: %s\n", req.URL)
	log.Println("Header: ")
	for key, value := range req.Header {
		log.Printf("%s : %v", key, value)
	}
	setReq(nil)
}
func Logresponse(resp *http.Response, LogFile *os.File) {
	n := len(logs)
	logs[n-1].Response = makeResponseLog(resp)
	WriteFile(LogFile)
	setResp(nil)
}

func Logger() {
	logFile := OpenLogFile()
	ReadLogFile(logFile, logs)
	go func() {
		for {
			if shared.request != nil && shared.response == nil {
				Logrequest(shared.request)
				fmt.Println("___")
				fmt.Println(len(logs))
			}

			if shared.response != nil && shared.request == nil {
				Logresponse(shared.response, logFile)
			}
		}

	}()
}
