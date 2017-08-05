package data

import (
	"errors"
	"net/http"

	"github.com/rms1000watt/degeneres-test/helpers"
	log "github.com/sirupsen/logrus"
)

var (
	ErrFailedDecodingInput = errors.New("Failed decoding input")
)

type ManagementInP struct {
	Power []*bool `json:"power,omitempty" validate:"" transform:""`
}

type AttendanceP struct {
	Date *string `json:"date,omitempty" validate:"" transform:""`
}

type AddressP struct {
	Address1 *string  `json:"address_1,omitempty" validate:"" transform:""`
	Address2 *string  `json:"address_2,omitempty" validate:"" transform:""`
	City     *string  `json:"city,omitempty" validate:"" transform:""`
	State    *string  `json:"state,omitempty" validate:"" transform:""`
	Postal   *PostalP `json:"postal,omitempty" validate:"" transform:""`
}

type PostalP struct {
	Code      *string `json:"code,omitempty" validate:"" transform:""`
	Region    *string `json:"region,omitempty" validate:"" transform:""`
	Utilities *string `json:"utilities,omitempty" validate:"" transform:""`
}

type PersonInP struct {
	Id         *int64         `json:"id,omitempty" validate:"" transform:""`
	FirstName  *string        `json:"first_name,omitempty" validate:"maxLength=100" transform:"truncate=50"`
	LastName   *string        `json:"last_name,omitempty" validate:"maxLength=1000,required,minLength=1" transform:"truncate=50,hash"`
	Address    *AddressP      `json:"address,omitempty" validate:"" transform:""`
	Attendance []*AttendanceP `json:"attendance,omitempty" validate:"" transform:""`
}

type TicketInP struct {
	Id *string `json:"id,omitempty" validate:"" transform:""`
}

func GetManagementIn(r *http.Request) (managementIn ManagementIn, err error) {
	inputP := &ManagementInP{}
	if err := helpers.Unmarshal(r, inputP); err != nil {
		log.Error("Failed decoding input:", err)
		return managementIn, ErrFailedDecodingInput
	}

	msg, err := helpers.Validate(inputP)
	if err != nil {
		return managementIn, err
	}

	if msg != "" {
		return managementIn, errors.New(msg)
	}

	if err := helpers.Transform(inputP); err != nil {
		return managementIn, err
	}

	managementIn = ConvertManagementInP(inputP)

	return
}

func GetPersonIn(r *http.Request) (personIn PersonIn, err error) {
	inputP := &PersonInP{}
	if err := helpers.Unmarshal(r, inputP); err != nil {
		log.Error("Failed decoding input:", err)
		return personIn, ErrFailedDecodingInput
	}

	msg, err := helpers.Validate(inputP)
	if err != nil {
		return personIn, err
	}

	if msg != "" {
		return personIn, errors.New(msg)
	}

	if err := helpers.Transform(inputP); err != nil {
		return personIn, err
	}

	personIn = ConvertPersonInP(inputP)

	return
}

func GetTicketIn(r *http.Request) (ticketIn TicketIn, err error) {
	inputP := &TicketInP{}
	if err := helpers.Unmarshal(r, inputP); err != nil {
		log.Error("Failed decoding input:", err)
		return ticketIn, ErrFailedDecodingInput
	}

	msg, err := helpers.Validate(inputP)
	if err != nil {
		return ticketIn, err
	}

	if msg != "" {
		return ticketIn, errors.New(msg)
	}

	if err := helpers.Transform(inputP); err != nil {
		return ticketIn, err
	}

	ticketIn = ConvertTicketInP(inputP)

	return
}

func ConvertManagementInP(managementInP *ManagementInP) (managementIn ManagementIn) {

	power := []bool{}
	for _, field := range managementInP.Power {
		power = append(power, *field)
	}
	managementIn.Power = power

	return
}

func ConvertAttendanceP(attendanceP *AttendanceP) (attendance Attendance) {

	if attendanceP.Date != nil {
		attendance.Date = *attendanceP.Date
	}

	return
}

func ConvertAddressP(addressP *AddressP) (address Address) {

	if addressP.Address1 != nil {
		address.Address1 = *addressP.Address1
	}

	if addressP.Address2 != nil {
		address.Address2 = *addressP.Address2
	}

	if addressP.City != nil {
		address.City = *addressP.City
	}

	if addressP.State != nil {
		address.State = *addressP.State
	}

	if addressP.Postal != nil {
		address.Postal = ConvertPostalP(addressP.Postal)
	}

	return
}

func ConvertPostalP(postalP *PostalP) (postal Postal) {

	if postalP.Code != nil {
		postal.Code = *postalP.Code
	}

	if postalP.Region != nil {
		postal.Region = *postalP.Region
	}

	if postalP.Utilities != nil {
		postal.Utilities = *postalP.Utilities
	}

	return
}

func ConvertPersonInP(personInP *PersonInP) (personIn PersonIn) {

	if personInP.Id != nil {
		personIn.Id = *personInP.Id
	}

	if personInP.FirstName != nil {
		personIn.FirstName = *personInP.FirstName
	}

	if personInP.LastName != nil {
		personIn.LastName = *personInP.LastName
	}

	if personInP.Address != nil {
		personIn.Address = ConvertAddressP(personInP.Address)
	}

	attendance := []Attendance{}
	for _, field := range personInP.Attendance {
		attendance = append(attendance, ConvertAttendanceP(field))
	}
	personIn.Attendance = attendance

	return
}

func ConvertTicketInP(ticketInP *TicketInP) (ticketIn TicketIn) {

	if ticketInP.Id != nil {
		ticketIn.Id = *ticketInP.Id
	}

	return
}
