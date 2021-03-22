package entity

type Nomenclature struct {
	Groups            []Group       `json:"groups"`
	ProductCategories []interface{} `json:"productCategories"`
	Products          []Product     `json:"products"`
	Revision          int           `json:"revision"`
	UploadDate        string        `json:"uploadDate"`
}

type Group struct {
	AdditionalInfo   string        `json:"additionalInfo"`
	Code             interface{}   `json:"code"`
	Description      interface{}   `json:"description"`
	ID               string        `json:"id"`
	IsDeleted        bool          `json:"isDeleted"`
	Name             string        `json:"name"`
	SeoDescription   interface{}   `json:"seoDescription"`
	SeoKeywords      interface{}   `json:"seoKeywords"`
	SeoText          interface{}   `json:"seoText"`
	SeoTitle         interface{}   `json:"seoTitle"`
	Tags             interface{}   `json:"tags"`
	Images           []interface{} `json:"images"`
	IsIncludedInMenu bool          `json:"isIncludedInMenu"`
	Order            int           `json:"order"`
	ParentGroup      interface{}   `json:"parentGroup"`
}

type Product struct {
	AdditionalInfo         interface{}   `json:"additionalInfo"`
	Code                   string        `json:"code"`
	Description            string        `json:"description"`
	ID                     string        `json:"id"`
	IsDeleted              bool          `json:"isDeleted"`
	Name                   string        `json:"name"`
	SeoDescription         interface{}   `json:"seoDescription"`
	SeoKeywords            interface{}   `json:"seoKeywords"`
	SeoText                interface{}   `json:"seoText"`
	SeoTitle               interface{}   `json:"seoTitle"`
	Tags                   interface{}   `json:"tags"`
	CarbohydrateAmount     int           `json:"carbohydrateAmount"`
	CarbohydrateFullAmount int           `json:"carbohydrateFullAmount"`
	DifferentPricesOn      interface{}   `json:"differentPricesOn"`
	DoNotPrintInCheque     bool          `json:"doNotPrintInCheque"`
	EnergyAmount           int           `json:"energyAmount"`
	EnergyFullAmount       int           `json:"energyFullAmount"`
	FatAmount              int           `json:"fatAmount"`
	FatFullAmount          int           `json:"fatFullAmount"`
	FiberAmount            int           `json:"fiberAmount"`
	FiberFullAmount        int           `json:"fiberFullAmount"`
	GroupID                interface{}   `json:"groupId"`
	GroupModifiers         []interface{} `json:"groupModifiers"`
	MeasureUnit            string        `json:"measureUnit"`
	Modifiers              []interface{} `json:"modifiers"`
	Price                  int           `json:"price"`
	ProductCategoryID      interface{}   `json:"productCategoryId"`
	ProhibitedToSaleOn     interface{}   `json:"prohibitedToSaleOn"`
	Type                   string        `json:"type"`
	UseBalanceForSell      bool          `json:"useBalanceForSell"`
	Weight                 int           `json:"weight"`
	IsIncludedInMenu       bool          `json:"isIncludedInMenu"`
	Order                  int           `json:"order"`
	ParentGroup            string        `json:"parentGroup"`
	WarningType            int           `json:"warningType"`
}
