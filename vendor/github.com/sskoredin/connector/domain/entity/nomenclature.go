package entity

type Nomenclature struct {
	Groups            []Group           `json:"groups"`
	ProductCategories []ProductCategory `json:"productCategories"`
	Products          []Product         `json:"products"`
	Revision          int               `json:"revision"`
	UploadDate        string            `json:"uploadDate"`
}

type ProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
	Code              string  `json:"code"`
	ID                string  `json:"id,omitempty"`
	IsDeleted         bool    `json:"isDeleted"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	ProductCategoryID string  `json:"productCategoryId"`
	Type              string  `json:"type"`
}
