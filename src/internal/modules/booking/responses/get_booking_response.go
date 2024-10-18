package booking_responses

import "time"

type GetBookingResponse struct {
	ID        string    `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
