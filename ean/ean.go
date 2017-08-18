package ean

import (
	"encoding/json"
	"fmt"
	"github.com/ctrlaltreboot/igor/helper"
	"net/http"
)

type EanResponse struct {
	HotelListResponse HotelListResponse `json:"HotelListResponse"`
}

type HotelListResponse struct {
	HotelList HotelList `json:"HotelList"`
}

type HotelList struct {
	HotelSummary []HotelSummary `json:"HotelSummary"`
}

type HotelSummary struct {
	HotelId             string          `json:"hotelId"`
	RoomRateDetailsList RoomRateDetails `json:"RoomRateDetailsList"`
}

type RoomRateDetails struct {
	RoomTypeCode    string    `json:"roomTypeCode"`
	RoomDescription string    `json:"roomDescription"`
	RateInfos       RateInfos `json:"RateInfos`
}

type RateInfos struct {
	RateInfo      RateInfo `json:"RateInfo"`
	NonRefundable string   `json:"nonRefundable"`
}

type RateInfo struct {
	ChargeableRateInfo ChargeableRateInfo `json:"ChargeableRateInfo"`
}

type ChargeableRateInfo struct {
	Total string `json:"@total"`
}

type CheapestOffer struct {
	PropertyName  string `json:"property_name"`
	RoomTypeName  string `json:"room_type_name"`
	OfferName     string `json:"offer_name"`
	Amount        string `json:"amount"`
	NonRefundable string `json:"is_non_refundable"`
}

func List(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch("http://127.0.0.1:5092/ean")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}

func Cheapest(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch("http://127.0.0.1:5092/ean")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	var e EanResponse
	err = json.Unmarshal(b, &e)
	fmt.Fprintf(w, "%q", e)
	for _, deets := range e.HotelListResponse.HotelList.HotelSummary {
		fmt.Fprintf(w, "%q", deets.RoomRateDetailsList.RateInfos.RateInfo.ChargeableRateInfo.Total)
	}
}
