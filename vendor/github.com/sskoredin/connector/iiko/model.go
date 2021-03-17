package iiko_client

import "encoding/xml"

type ResponseGoodsToFTPData struct {
	XMLName xml.Name `xml:"report"`
	Text    string   `xml:",chardata"`
	R       []struct {
		Text               string `xml:",chardata"`
		ProductNum         string `xml:"Product.Num"`
		FinalBalanceAmount string `xml:"FinalBalance.Amount"`
	} `xml:"r"`
}

type ResponseIikoOlAPReportData struct {
	XMLName xml.Name `xml:"report"`
	Text    string   `xml:",chardata"`
	R       []struct {
		Text                       string  `xml:",chardata"`
		DiscountSum                float64 `xml:"DiscountSum"`
		DishAmountInt              float64 `xml:"DishAmountInt"`
		DishDiscountSumInt         float64 `xml:"DishDiscountSumInt"`
		JurName                    string  `xml:"JurName"`
		DishGroupSecondParent      string  `xml:"DishGroup.SecondParent"`
		DishCode                   string  `xml:"DishCode"`
		OrderNum                   int     `xml:"OrderNum"`
		CloseTime                  string  `xml:"CloseTime"`
		OrderDiscountType          string  `xml:"OrderDiscount.Type"`
		ProductCostBaseProductCost float64 `xml:"ProductCostBase.ProductCost"`
		DishSumInt                 float64 `xml:"DishSumInt"`
		PayTypes                   string  `xml:"PayTypes"`
		DishName                   string  `xml:"DishName"`
	} `xml:"r"`
}

type ResponseIikoAmountReportData struct {
	XMLName xml.Name `xml:"report"`
	Text    string   `xml:",chardata"`
	R       []struct {
		Text                string `xml:",chardata"`
		StartBalanceAmount  string `xml:"StartBalance.Amount"`
		AmountIn            string `xml:"Amount.In"`
		FinalBalanceAmount  string `xml:"FinalBalance.Amount"`
		ProductNum          string `xml:"Product.Num"`
		ProductName         string `xml:"Product.Name"`
		AmountOut           string `xml:"Amount.Out"`
		ProductSecondParent string `xml:"Product.SecondParent"`
	} `xml:"r"`
}
