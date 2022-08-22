package dto

type CreateHotelRequest struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Rate    int8   `json:"rate"`
}
