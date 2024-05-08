package models

type Reservation struct {
	ID                 string              `json:"id"`
	UserID             string              `json:"user_id"`
	ProductID          string              `json:"product_id"`
	TransactionType    string              `json:"transaction_type"`
	ReservationDate    string              `json:"reservation_date"`
	PickupDate         string              `json:"pickup_date"`
	ReturnDate         string              `json:"return_date"`
	PaymentStatus      string              `json:"payment_status"`
	ReservationStatus  string              `json:"reservation_status"`
	TotalAmount        string              `json:"total_amount"`
	PaymentMethod      string              `json:"payment_method"`
	DeliveryAddress    DeliveryAddress     `json:"delivery_address"`
	ConfirmationDate   string              `json:"confirmation_date"`
	CancellationDate   string              `json:"cancellation_date"`
	Rating             string              `json:"rating"`
	Comment            string              `json:"comment"`
	Duration           string              `json:"duration"`
	AdditionalServices []AdditionalService `json:"additional_services"`
}

type AdditionalService struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type DeliveryAddress struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
}
