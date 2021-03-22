package entity

type IikoDiscount struct {
	DepartmentCrmIds         []string      `json:"departmentCrmIds"`
	ID                       string        `json:"id"`
	Name                     string        `json:"name"`
	Percent                  int           `json:"percent"`
	IsCategorisedDiscount    bool          `json:"isCategorisedDiscount"`
	ProductCategoryDiscounts []interface{} `json:"productCategoryDiscounts"`
	Comment                  string        `json:"comment"`
	CanBeAppliedSelectively  bool          `json:"canBeAppliedSelectively"`
	MinOrderSum              int           `json:"minOrderSum"`
	Mode                     string        `json:"mode"`
	Sum                      int           `json:"sum"`
	IsManual                 bool          `json:"isManual"`
	IsCard                   bool          `json:"isCard"`
	CanApplyByCardNumber     bool          `json:"canApplyByCardNumber"`
	IsAutomatic              bool          `json:"isAutomatic"`
}
