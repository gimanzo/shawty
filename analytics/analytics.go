package analytics

import (
	"github.com/jpillora/go-ogle-analytics"
	"os"
	"fmt"
)

func Log(category EventCategory, url string, hash string, status EventStatus) {
	trackingId := os.Getenv("GOOGLE_TRACKER_ID")
	if trackingId == "" {
		return
	}

	client, err := ga.NewClient(trackingId)
	if err != nil {
		fmt.Println(err)
		return
	}

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