package constant

type ServiceID string

const (
	Gateway ServiceID = "gateway"
	User    ServiceID = "users"
	Store   ServiceID = "stores"
	Booking ServiceID = "bookings"
)

func (id ServiceID) String() string {
	switch id {
	case Gateway:
		return "gateway"
	case User:
		return "user"
	case Store:
		return "store"
	case Booking:
		return "booking"
	default:
		return ""
	}
}

type Service struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (s Service) IsValid() bool {
	if s.Host == "" {
		return false
	}
	return true
}

func NewServiceWithID(id ServiceID) ServiceID {
	return ServiceID(id)
}

var ServicesProperties = map[ServiceID]Service{
	Store: {
		Host: Store.String(),
		Port: 8090,
	},
	Booking: {
		Host: Booking.String(),
		Port: 8091,
	},
	User: {
		Host: User.String(),
		Port: 8092,
	},
}
