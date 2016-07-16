package analytics

import (
	"github.com/gimanzo/go-ogle-analytics"
	"os"
	"fmt"
	"net/http"
)

// Parameter Reference:
//https://developers.google.com/analytics/devguides/collection/protocol/v1/devguide
//https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters

func Log(category EventCategory, url string, hash string, status EventStatus, request *http.Request) {
	trackingId := os.Getenv("GOOGLE_TRACKER_ID")
	if trackingId == "" {
		return
	}

	client, err := ga.NewClient(trackingId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if(os.Getenv("DEBUG") != "") {
		fmt.Println("Headers", request)
	}

	client.ApplicationName("kwk.co")

	client.UserAgentOverride(request.UserAgent())
	client.UserLanguage(request.Header.Get("Accept-Language"))
	client.IPOverride(request.RemoteAddr)

	client.DocumentReferrer(request.Referer())
	client.DocumentHostName(request.Host)
	//client.DocumentEncoding(request.TransferEncoding)

	client.ProductCategory("gim.nz")

	go func() {
		fmt.Println("Sending event:", category, url, hash,  status, trackingId)
		client.Send(ga.NewEvent(string(category), url + "," + hash + "," + string(status)))
		if err != nil {
			fmt.Println(err)
		}
	}()
}

type EventCategory string

const  (
	CategoryEncode EventCategory = "Encode"
	CategoryRedirect = "Redirect"
	CategoryDecode = "Decode"
)

type EventStatus string

const  (
	StatusHit EventStatus = "Hit"
	StatusMiss = "Miss"
	StatusSuccess = "Success"
)