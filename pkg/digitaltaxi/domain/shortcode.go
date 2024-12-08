package domain

// IncomingSMSPayload represents the payload sent to Africa's Talking server when an incoming message is received
type IncomingSMSPayload struct {
	ID          string `json:"id"`
	LinkID      string `json:"linkId"`
	Date        string `json:"date"`
	From        string `json:"from"`
	Text        string `json:"text"`
	To          string `json:"to"`
	NetworkCode string `json:"networkCode"`
}

// USSDPayload represents the payload sent to Africa's Talking server when an incoming USSD request is received
type USSDPayload struct {
	SessionID        string                `json:"sessionId"`
	PhoneNumber      string                `json:"phoneNumber"`
	NetworkCode      string                `json:"networkCode"`
	ServiceCode      string                `json:"serviceCode"`
	Text             string                `json:"text"`
	Level            int                   `json:"level"`
	ImmunizationData *ImmunizationResponse `json:"immunizationData,omitempty"`
	EmergencyData    *EmergencyResponse    `json:"emergencyData,omitempty"`
}

// ImmunizationResponse represents the data saved in the db when the user makes an immunization request
type ImmunizationResponse struct {
	Disease   string `json:"disease,omitempty"`
	County    string `json:"county,omitempty"`
	SubCounty string `json:"subCounty,omitempty"`
	Ward      string `json:"ward,omitempty"`
	Facility  string `json:"facility,omitempty"`
}

// EmergencyResponse holds the data saved in the db when the user makes an emergency request
type EmergencyResponse struct {
}
