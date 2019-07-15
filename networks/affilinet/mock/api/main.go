package main

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/halorium/httprouter"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
)

func main() {
	router := httprouter.New()

	router.POST("/V2.0/Logon.svc", postLogon)

	router.POST("/V2.0/PublisherStatistics.svc", postPublisherStatistics)

	panic(
		http.ListenAndServe(":80", router),
	)
}

func postLogon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !validHeaders(w, r) {
		http.Error(w, "Missing required SOAP headers", 500)
		return
	}

	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(200)

	res, err := ioutil.ReadFile("response-authentication.xml")

	if err != nil {
		logger.Panic("networks-affilinet-mock-api-post", flaw.From(err))
	}

	_, err = w.Write(res)

	if err != nil {
		logger.Panic("networks-affilinet-mock-api-post", flaw.From(err))
	}
}

func postPublisherStatistics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !validHeaders(w, r) {
		http.Error(w, "Missing required SOAP headers", 500)
		return
	}

	var err error

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logger.Panic("networks-affilinet-mock-api-post-publisher-statistics", flaw.From(err))
	}

	bodyString := string(bodyBytes)

	regExp := regexp.MustCompile(`CurrentPage>(\d)`)

	matches := regExp.FindStringSubmatch(bodyString)

	var responseBytes []byte

	if matches == nil {
		logger.Panic("networks-affilinet-mock-api-post-publisher-statistics", nil)
	} else {
		pageNumber := matches[1]
		responseBytes, err = ioutil.ReadFile("response-page-" + pageNumber + ".xml")

		if err != nil {
			logger.Panic("networks-affilinet-mock-api-post", flaw.From(err))
		}
	}

	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(200)

	_, err = w.Write(responseBytes)

	if err != nil {
		logger.Panic("networks-affilinet-mock-api-post", flaw.From(err))
	}
}

func validHeaders(_ http.ResponseWriter, r *http.Request) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return false
	}

	matched, err := regexp.MatchString("text/xml.*", contentType)
	if err != nil || !matched {
		return false
	}

	return r.Header.Get("SOAPAction") != ""
}
