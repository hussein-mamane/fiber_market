package models

type Offer struct {
	OfferId     int    `json:"offer_id"`
	ProductName string `json:"product_name"`
	ImagePath   string `json:"image_path"`
	//ProposerId int `json:proposer_id`,
}
