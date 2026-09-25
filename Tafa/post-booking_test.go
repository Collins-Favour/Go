package main

import (
	"encoding/json"
	"testing"
)

func TestPostBooking(t *testing.T) {
	bookingData := BookingRequest{
		UserName:    "Magnetto steward",
		PhoneNumber: "0786510813",
		UserTickets: 13,
	}
	jsonBytes, err := json.Marshal(bookingData)
	if err != nil {
		t.Errorf("failed to encode : %v", err)
	}

}
