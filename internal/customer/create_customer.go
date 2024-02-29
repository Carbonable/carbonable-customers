package customer

import (
	"errors"
	"net/http"
	"slices"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	ErrWalletRequired  = errors.New("wallet is required")
	ErrWalletInvalid   = errors.New("wallet is invalid")
	ErrCountryRequired = errors.New("country is required")
	ErrNetworkRequired = errors.New("network is required")
	ErrInvalidNetwork  = errors.New("network is invalid")
)

type customerDTO struct {
	Wallet  string `json:"wallet"`
	Country string `json:"country"`
	Network string `json:"network"`
}

func (c *customerDTO) Validate() []error {
	var errs []error
	if c.Wallet == "" {
		errs = append(errs, ErrWalletRequired)
	}
	var f felt.Felt
	err := f.UnmarshalJSON([]byte(c.Wallet))
	if err != nil {
		errs = append(errs, ErrWalletInvalid)
	}

	if c.Country == "" {
		errs = append(errs, ErrCountryRequired)
	}
	if c.Network == "" {
		errs = append(errs, ErrNetworkRequired)
	}
	if !slices.Contains([]string{"mainnet", "goerli", "sepolia"}, c.Network) {
		errs = append(errs, ErrInvalidNetwork)
	}

	return errs
}

func CreateCustomer(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var customerDTO customerDTO
		if err := c.Bind(&customerDTO); err != nil {
			return err
		}

		err := customerDTO.Validate()
		if len(err) > 0 {
			var errStrings []string
			for _, e := range err {
				errStrings = append(errStrings, e.Error())
			}
			return c.JSON(http.StatusBadRequest, errStrings)
		}

		customer := FromHttpRequest(customerDTO.Wallet, customerDTO.Country, customerDTO.Network)
		db.Create(&customer)
		return c.String(http.StatusOK, "Hello, World!")
	}
}
