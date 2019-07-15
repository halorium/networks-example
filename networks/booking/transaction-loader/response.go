package booking

type transaction struct {
	AccomodationType string `csv:"accomodation type" json:"accomodation-type,omitempty"`
	AffiliateID      string `csv:"affiliate id"      json:"affiliate-id,omitempty"`
	Arrival          string `csv:"arrival"           json:"arrival,omitempty"`
	BookNumber       string `csv:"book nr."          json:"book-number,omitempty"`
	Booked           string `csv:"booked"            json:"booked,omitempty"`
	BookerCountry    string `csv:"booker country"    json:"booker-country,omitempty"`
	BookerLanguage   string `csv:"booker language"   json:"booker-language,omitempty"`
	Commission       string `csv:"comission ( EUR )" json:"comission,omitempty"`
	Departure        string `csv:"departure"         json:"departure,omitempty"`
	Fee              string `csv:"fee ( EUR )"       json:"fee,omitempty"`
	HotelCity        string `csv:"hotel city"        json:"hotel-city,omitempty"`
	HotelCountry     string `csv:"hotel country"     json:"hotel-country,omitempty"`
	HotelName        string `csv:"hotel name"        json:"hotel-name,omitempty"`
	HotelUfi         string `csv:"hotel ufi"         json:"hotel-ufi,omitempty"`
	Label            string `csv:"label"             json:"label,omitempty"`
	Percent          string `csv:"perc"              json:"percent,omitempty"`
	SlipNumber       string `csv:"slip nr."          json:"slip-number,omitempty"`
}
