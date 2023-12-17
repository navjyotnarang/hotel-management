package api

type Room struct {
	ID       int    `json:"id"`
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}

// Other models can be added for guests, reservations, etc.
