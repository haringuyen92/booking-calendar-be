package filters

type BookingFilter struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    uint   `json:"user_id" bson:"user_id,omitempty"`
	StoreID   uint   `json:"store_id" bson:"store_id,omitempty"`
	IsDeleted bool   `json:"is_deleted" bson:"is_deleted"`
}
