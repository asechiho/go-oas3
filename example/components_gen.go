// This file is generated by github.com/mikekonan/go-oas3. DO NOT EDIT.

package example

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/google/uuid"
	countries "github.com/mikekonan/go-types/v2/country"
	currency "github.com/mikekonan/go-types/v2/currency"
	email "github.com/mikekonan/go-types/v2/email"
	url "github.com/mikekonan/go-types/v2/url"
	"regexp"
	"strings"
	"time"
)

var regexParamRegex = regexp.MustCompile("^[.?\\d]+$")

type Email = email.Email

type genericResponse struct {
	Result GenericResponseResultEnum `json:"result"`
}

type GenericResponse struct {
	Result GenericResponseResultEnum `json:"result"`
}

func (body *GenericResponse) UnmarshalJSON(data []byte) error {
	var value genericResponse
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	body.Result = value.Result

	return nil
}
func (body GenericResponse) Validate() error {
	return nil
}

type Time = time.Time

type URL = url.URL

type Boolean = bool

type createTransactionRequest struct {
	Amount        float64              `json:"amount"`
	AmountCents   int                  `json:"amountCents"`
	CallbackURL   url.URL              `json:"callbackURL"`
	Country       countries.Alpha2Code `json:"country"`
	Currency      currency.Code        `json:"currency"`
	Description   *string              `json:"description"`
	Details       *string              `json:"details,omitempty"`
	Email         email.Email          `json:"email"`
	RegexParam    string               `json:"regexParam"`
	Title         string               `json:"title"`
	TransactionID uuid.UUID            `json:"transactionID"`
}

type CreateTransactionRequest struct {
	Amount        float64              `json:"amount"`
	AmountCents   int                  `json:"amountCents"`
	CallbackURL   url.URL              `json:"callbackURL"`
	Country       countries.Alpha2Code `json:"country"`
	Currency      currency.Code        `json:"currency"`
	Description   string               `json:"description"`
	Details       *string              `json:"details,omitempty"`
	Email         email.Email          `json:"email"`
	RegexParam    string               `json:"regexParam"`
	Title         string               `json:"title"`
	TransactionID uuid.UUID            `json:"transactionID"`
}

func (body *CreateTransactionRequest) UnmarshalJSON(data []byte) error {
	var value createTransactionRequest
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	body.Currency = value.Currency
	body.Email = value.Email
	if !regexParamRegex.MatchString(body.RegexParam) {
		return fmt.Errorf("RegexParam not matched by the '^[.?\\d]+$' regex")
	}
	body.RegexParam = value.RegexParam
	body.Title = strings.TrimSpace(value.Title)
	body.TransactionID = value.TransactionID
	body.AmountCents = value.AmountCents
	body.CallbackURL = value.CallbackURL
	body.Details = value.Details
	body.Amount = value.Amount
	body.Country = value.Country

	if value.Description == nil {
		return fmt.Errorf("Description is required")
	}

	body.Description = strings.TrimSpace(*value.Description)

	return nil
}
func (body CreateTransactionRequest) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(&body.Currency, validation.Skip.When(body.Currency == ""), validation.RuneLength(3, 3)),
		validation.Field(&body.Title, validation.Skip.When(body.Title == ""), validation.RuneLength(8, 50)),
		validation.Field(&body.Amount, validation.Min(0.009).Exclusive()),
		validation.Field(&body.Country, validation.Skip.When(body.Country == ""), validation.RuneLength(2, 2)),
		validation.Field(&body.Description, validation.Required, validation.RuneLength(8, 100)))
}

type RawPayload = []byte

type WithEnum string

var WithEnumOne WithEnum = "one"
var WithEnumTwo WithEnum = "two"

func (enum WithEnum) Check() error {
	switch enum {
	case WithEnumOne, WithEnumTwo:

		return nil
	}

	return fmt.Errorf("invalid WithEnum enum value")
}

func (enum *WithEnum) UnmarshalJSON(data []byte) error {
	var strValue string
	if err := json.Unmarshal(data, &strValue); err != nil {

		return err
	}
	enumValue := WithEnum(strValue)
	if err := enumValue.Check(); err != nil {

		return err
	}
	*enum = enumValue

	return nil
}

type GenericResponseResultEnum string

var GenericResponseResultEnumSuccess GenericResponseResultEnum = "success"
var GenericResponseResultEnumFailed GenericResponseResultEnum = "failed"

func (enum GenericResponseResultEnum) Check() error {
	switch enum {
	case GenericResponseResultEnumSuccess, GenericResponseResultEnumFailed:

		return nil
	}

	return fmt.Errorf("invalid GenericResponseResultEnum enum value")
}

func (enum *GenericResponseResultEnum) UnmarshalJSON(data []byte) error {
	var strValue string
	if err := json.Unmarshal(data, &strValue); err != nil {

		return err
	}
	enumValue := GenericResponseResultEnum(strValue)
	if err := enumValue.Check(); err != nil {

		return err
	}
	*enum = enumValue

	return nil
}
