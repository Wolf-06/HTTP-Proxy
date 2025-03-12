package res

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Recieved req")
	setReq(req)
	targetURL, Url_err := url.Parse("http://www.testingmcafeesites.com/")
	if Url_err != nil {
		log.Fatalf("Error while url parsing: ", Url_err)
	}

	//create a request copy with absolute url
	newReq := cloneRequest(req, targetURL)

	// send the request and recieve the response
	client := &http.Client{}
	resp, resp_err := client.Do(newReq)
	if resp_err != nil {
		log.Fatalf("Error sending the request: ", resp_err)
	}
	copyResponse(w, resp)
	setResp(resp)
}

func cloneRequest(req *http.Request, targetUrl *url.URL) *http.Request {

	tempReq := req.Clone(req.Context())
	// converting the realative url to absolute url, un-comment if needed.
	//tempReq.URL = targetUrl
	//tempReq.Host = targetUrl.Host
	//tempReq.URL.Path = targetUrl.Path
	//tempReq.URL.RawQuery = targetUrl.RawQuery
	tempReq.RequestURI = "" //resets the uri
	return tempReq
}
