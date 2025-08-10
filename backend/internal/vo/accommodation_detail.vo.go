package vo

type Beds struct {
	SingleBed           uint `json:"single_bed" validate:"gte=0,lte=10"`
	DoubleBed           uint `json:"double_bed" validate:"gte=0,lte=10"`
	LargeDoubleBed      uint `json:"large_double_bed" validate:"gte=0,lte=10"`
	ExtraLargeDoubleBed uint `json:"extra_large_double_bed" validate:"gte=0,lte=10"`
}
type GetAccommodationDetailsInput struct {
	AccommodationID string `uri:"id"`
	CheckIn         string `form:"check_in"`
	CheckOut        string `form:"check_out"`
}

type GetAccommodationDetailsByManagerInput struct {
	AccommodationID string `uri:"id"`
}

type CreateAccommodationDetailInput struct {
	AccommodationID string   `json:"accommodation_id" validate:"required"`
	Name            string   `json:"name" validate:"required,min=1,max=255"`
	Guests          uint8    `json:"guests" validate:"gte=1,lte=50"`
	Beds            Beds     `json:"beds" validate:"required"`
	Facilities      []string `json:"facilities"`
	Price           string   `json:"price" validate:"required"`
	DiscountID      string   `json:"discount_id"`
}

type FacilityDetailOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateAccommodationDetailOutput struct {
	ID              string                 `json:"id"`
	AccommodationID string                 `json:"accommodation_id"`
	Name            string                 `json:"name"`
	Guests          uint8                  `json:"guests"`
	Beds            Beds                   `json:"beds"`
	Facilities      []FacilityDetailOutput `json:"facilities"`
	Price           string                 `json:"price"`
	DiscountID      string                 `json:"discount_id"`
	Images          []string               `json:"images"`
}

type GetAccommodationDetailsOutput struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Guests         uint8                  `json:"guests"`
	Beds           Beds                   `json:"beds"`
	Facilities     []FacilityDetailOutput `json:"facilities"`
	AvailableRooms uint8                  `json:"available_rooms"`
	Price          string                 `json:"price"`
	DiscountID     string                 `json:"discount_id"`
	Images         []string               `json:"images"`
}

type GetAccommodationDetailsByManagerOutput struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Guests         uint8                  `json:"guests"`
	Beds           Beds                   `json:"beds"`
	Facilities     []FacilityDetailOutput `json:"facilities"`
	AvailableRooms uint8                  `json:"available_rooms"`
	Price          string                 `json:"price"`
	DiscountID     string                 `json:"discount_id"`
}

type UpdateAccommodationDetailInput struct {
	ID              string   `json:"id" validate:"required"`
	AccommodationID string   `json:"accommodation_id"`
	Name            string   `json:"name"`
	Guests          uint8    `json:"guests"`
	Beds            Beds     `json:"beds"`
	Facilities      []string `json:"facilities"`
	Price           string   `json:"price"`
	DiscountID      string   `json:"discount_id"`
}

type UpdateAccommodationDetailOutput struct {
	ID              string                 `json:"id"`
	AccommodationID string                 `json:"accommodation_id"`
	Name            string                 `json:"name"`
	Guests          uint8                  `json:"guests"`
	Beds            Beds                   `json:"beds"`
	Facilities      []FacilityDetailOutput `json:"facilities"`
	Price           string                 `json:"price"`
	DiscountID      string                 `json:"discount_id"`
	Images          []string               `json:"images"`
}

type DeleteAccommodationDetailInput struct {
	ID string `json:"id" validate:"required"`
}
