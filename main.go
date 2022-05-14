package main

import (
	"flag"
	"github.com/Ignatella/RegisterAghPanel/html"
	"github.com/Ignatella/RegisterAghPanel/request"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	PanelUrlC           = "https://panel.dsnet.agh.edu.pl"
	LoginUrlC           = PanelUrlC + "/login_check"
	ReservationPageUrlC = PanelUrlC + "/reserv/rezerwuj/"
)

var (
	PanelUrl           string
	LoginUrl           string
	ReservationPageUrl string
)

var (
	TargetReservationPageId string
	TargetReservationDate   string
	TargetReservationTime   string
)

var (
	Username string
	Password string
)

func init() {
	flag.StringVar(&PanelUrl, "panelUrl", PanelUrlC, "panel url")
	flag.StringVar(&LoginUrl, "loginUrl", LoginUrlC, "login request url")
	flag.StringVar(&ReservationPageUrl, "reservationPageUrl", ReservationPageUrlC, "login request url")

	flag.StringVar(&Username, "username", "", "user password")
	flag.StringVar(&Password, "password", "", "user password")

	flag.StringVar(&TargetReservationPageId, "reservationPageId", "", "desired reservation page id. Ex. 2750")
	flag.StringVar(&TargetReservationDate, "date", "", "desired reservation date")
	flag.StringVar(&TargetReservationTime, "time", "", "desired reservation time")

	flag.Parse()
}

func main() {
	start := time.Now()

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = request.Login(client, LoginUrl, Username, Password)
	if err != nil {
		log.Fatal(err)
		return
	}

	htmlString, err := request.GetBookingPage(client, ReservationPageUrl+TargetReservationPageId)
	if err != nil {
		log.Println(err)
		return
	}

	reservations, err := html.FindAvailableReservations(htmlString)
	if err != nil {
		log.Fatal(err)
		return
	}

	targetReservationPath := reservations[TargetReservationDate][TargetReservationTime]

	if targetReservationPath == "" {
		log.Printf(
			"Desired date/time(" + TargetReservationDate + "/" + TargetReservationTime + ") is not available.",
		)
		return
	}

	err = request.Reserve(client, PanelUrl+targetReservationPath)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Success! Took: %s", elapsed)
}
