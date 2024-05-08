package models

type Product struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	Brand            string   `json:"brand"`
	Material         string   `json:"material"`
	Size             string   `json:"size"`
	Color            string   `json:"color"`
	RentalPrice      string   `json:"rental_price"`
	SalePrice        string   `json:"sale_price"`
	StockQuantity    string   `json:"stock_quantity"`
	AvailableForRent string   `json:"available_for_rent"`
	AvailableForSale string   `json:"available_for_sale"`
	RegistrationDate string   `json:"registration_date"`
	LastUpdate       string   `json:"last_update"`
	Images           []string `json:"images"`
	Category         string   `json:"category"`
	Tags             []string `json:"tags"`
	Rating           string   `json:"rating"`
	Reviews          []Review `json:"reviews"`
}

type Review struct {
	UserID    string `json:"user_id"`
	Rating    string `json:"rating"`
	Comment   string `json:"comment"`
	Timestamp string `json:"timestamp"`
}
