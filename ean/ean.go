package ean

import (
	"fmt"
	"github.com/ctrlaltreboot/igor/helper"
	"net/http"
)

type EanResponse struct {
	CustomerSessionID      string `json:"customerSessionId"`
	NumberOfRoomsRequested int32  `json:"numberOfRoomsRequested"`
	MoreResultsAvailable   bool   `json:"moreResultsAvailable"`
	CacheKey               string `json:"cacheKey"`
	CacheLocation          string `json:"cacheLocation"`
	CachedSupplierResponse map[string]interface{}
	HotelList              HotelList `json:"HotelList"`
}

type HotelList struct {
	Size                string         `json:"@size"`
	ActivePropertyCount string         `json:"activePropertyCount"`
	HotelSummary        []HotelSummary `json:"HotelSummary"`
}

type HotelSummary struct {
	Order                  string              `json:"@order"`
	UbsScore               string              `json:"@ubsScore"`
	HotelId                string              `json:"hotelId"`
	Name                   string              `json:"name"`
	Address1               string              `json:"address1"`
	City                   string              `json:"city"`
	StateProvinceCode      string              `json:"stateProvinceCode"`
	PostalCode             string              `json:"postalCode"`
	CountryCode            string              `json:"countryCode"`
	AirportCode            string              `json:"airportCode"`
	SupplierType           string              `json:"supplierType"`
	PropertyCategory       string              `json:"propertyCategory"`
	HotelRating            string              `json:"hotelRating"`
	HotelRatingDisplay     string              `json:"hotelRatingDisplay"`
	ConfidenceRating       string              `json:"confidenceRating"`
	AmenityMask            string              `json:"amenityMask"`
	TripAdvisorRating      string              `json:"tripAdvisorRating"`
	TripAdvisorReviewCount string              `json:"tripAdvisorReviewCount"`
	TripAdvisorRatingUrl   string              `json:"tripAdvisorRatingUrl"`
	LocationDescription    string              `json:"locationDescription"`
	ShortDescription       string              `json:"shortDescription"`
	HighRate               string              `json:"highRate"`
	LowRate                string              `json:"lowRate"`
	RateCurrencyCode       string              `json:"rateCurrencyCode"`
	Latitude               string              `json:"latitude"`
	Longitude              string              `json:"longitude"`
	ProximityDistance      string              `json:"proximityDistance"`
	ProximityUnit          string              `json:"proximityUnit"`
	HotelInDestination     string              `json:"hotelInDestination"`
	ThumbNailUrl           string              `json:"thumbNailUrl"`
	DeepLink               string              `json:"deepLink"`
	RoomRateDetailsList    RoomRateDetailsList `json:"RoomRateDetailsList"`
}

type RoomRateDetailsList struct {
	RoomRateDetails RoomRateDetails `json:"RoomRateDetails"`
}

type RoomRateDetails struct {
	RoomTypeCode        string    `json:"roomTypeCode"`
	RateCode            string    `json:"rateCode"`
	MaxRoomOccupancy    string    `json:"maxRoomOccupancy"`
	QuotedRoomOccupancy string    `json:"quotedRoomOccupancy"`
	MinGuestAge         string    `json:"minGuestAge"`
	RoomDescription     string    `json:"roomDescription"`
	PropertyAvailable   bool      `json:"propertyAvailable"`
	PropertyRestricted  bool      `json:"propertyRestricted"`
	ExpediaPropertyId   string    `json:"expediaPropertyId"`
	RateInfos           RateInfos `json:"RateInfos`
}

type RateInfos struct {
	Size             string   `json:"@size"`
	RateInfo         RateInfo `json:"RateInfo"`
	NonRefundable    string   `json:"nonRefundable"`
	RateType         string   `json:"rateType"`
	PromoId          string   `json:"promoId"`
	PromoDescription string   `json:"promoDescription"`
	PromoType        string   `json:"promoType"`
	CurrentAllotment string   `json:"currentAllotment"`
}

type RateInfo struct {
	PriceBreakdown     string             `json:"@priceBreakdown"`
	Promo              string             `json:"@promo"`
	RateChange         string             `json:"@rateChange"`
	RoomGroup          Room               `json:"RoomGroup"`
	ChargeableRateInfo ChargeableRateInfo `json:"ChargeableRateInfo"`
}

type Room struct {
	numberOfAdults
	numberOfChildren
	rateKey
	ChargeableNightlyRates []ChargeableNightlyRate `json:"ChargeableNightlyRates"`
}

type ChargeableNightlyRate struct {
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

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
