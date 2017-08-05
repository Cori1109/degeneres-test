package data

type PersonIn struct {
	Id         int64        `json:"id,omitempty"`
	FirstName  string       `json:"first_name,omitempty"`
	LastName   string       `json:"last_name,omitempty"`
	Address    Address      `json:"address,omitempty"`
	Attendance []Attendance `json:"attendance,omitempty"`
}

type PersonOut struct {
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Address   Address `json:"address,omitempty"`
	Profile   Profile `json:"profile,omitempty"`
}

type TicketIn struct {
	Id string `json:"id,omitempty"`
}

type TicketOut struct {
	Row  string `json:"row,omitempty"`
	Seat string `json:"seat,omitempty"`
}

type ManagementIn struct {
	Power []bool `json:"power,omitempty"`
}

type ManagementOut struct {
	Success []bool `json:"success,omitempty"`
}

type Profile struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type Attendance struct {
	Date string `json:"date,omitempty"`
}

type Address struct {
	Address1 string `json:"address_1,omitempty"`
	Address2 string `json:"address_2,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Postal   Postal `json:"postal,omitempty"`
}

type Postal struct {
	Code      string `json:"code,omitempty"`
	Region    string `json:"region,omitempty"`
	Utilities string `json:"utilities,omitempty"`
}
