package entity

type Organization struct {
	Address       string  `json:"address"`
	AverageCheque *string `json:"averageCheque"`
	Contact       struct {
		Email    string  `json:"email"`
		Location string  `json:"location"`
		Phone    *string `json:"phone"`
	} `json:"contact"`
	CurrencyIsoName  string  `json:"currencyIsoName"`
	Description      string  `json:"description"`
	FullName         string  `json:"fullName"`
	HomePage         *string `json:"homePage"`
	ID               string  `json:"id"`
	IsActive         bool    `json:"isActive"`
	Logo             *string `json:"logo"`
	LogoImage        *string `json:"logoImage"`
	MaxBonus         int     `json:"maxBonus"`
	MinBonus         int     `json:"minBonus"`
	Name             string  `json:"name"`
	NetworkID        *string `json:"networkId"`
	OrganizationType int     `json:"organizationType"`
	Phone            *string `json:"phone"`
	Timezone         string  `json:"timezone"`
	Website          *string `json:"website"`
	WorkTime         *string `json:"workTime"`
}
