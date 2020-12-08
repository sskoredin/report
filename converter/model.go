package converter

type ReportValue struct {
	Date                       string
	JurName                    string
	DishGroupSecondParent      string
	DishCode                   string
	DishName                   string
	PayTypes                   string
	DiscountType               string
	DishAmountInt              float64
	CloseTime                  string
	OrderNum                   int
	DishSumInt                 float64
	DiscountSum                float64
	DishDiscountSumInt         float64
	ProductCostBaseProductCost float64
}

type Report []ReportValue
