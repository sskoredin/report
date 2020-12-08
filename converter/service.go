package converter

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"mail/client"
	"time"
)

func Convert(responseData *client.ResponseData) Report {
	if responseData == nil {
		return nil
	}
	report := make([]ReportValue, len(responseData.R))
	for i, v := range responseData.R {
		closeTime, _ := time.Parse(time.UnixDate, v.CloseTime)

		value := ReportValue{
			Date:                       closeTime.Format("02.01.2006"),
			JurName:                    v.JurName,
			DishGroupSecondParent:      v.DishGroupSecondParent,
			DishCode:                   v.DishCode,
			DishName:                   v.DishName,
			PayTypes:                   v.PayTypes,
			DiscountType:               v.OrderDiscountType,
			DishAmountInt:              v.DishAmountInt,
			CloseTime:                  closeTime.Format("02.01.2006 15:04:05"),
			OrderNum:                   v.OrderNum,
			DishSumInt:                 v.DishSumInt,
			DiscountSum:                v.DiscountSum,
			DishDiscountSumInt:         v.DishDiscountSumInt,
			ProductCostBaseProductCost: v.ProductCostBaseProductCost,
		}
		report[i] = value
	}
	return report
}

func ReportName() string {

	now := time.Now()
	year, month, _ := now.Date()
	return fmt.Sprintf("./%d.%d.xlsx", month, year)
}

func (rep Report) ToXlsx() error {
	if rep == nil {
		return nil
	}
	f := xlsx.NewFile()
	s, err := f.AddSheet("report")
	if err != nil {
		return err
	}
	//add column names
	headers := []string{
		"Учетный день",
		"Юридическое лицо",
		"Группа блюда 2-го уровня",
		"Код блюда",
		"Блюдо",
		"Тип оплаты",
		"Тип скидки",
		"Количество блюд",
		"Время закрытия",
		"Номер чека",
		"Сумма без скидки, р.",
		"Сумма скидки, р.",
		"Сумма со скидкой, р.",
		"Себестоимость, р.",
	}
	headersRows := s.AddRow()
	for i := range headers {
		setValue(headersRows, headers[i])
	}
	for i, _ := range s.Cols {
		s.Cols[i].SetStyle(&xlsx.Style{
			Border: xlsx.Border{
				Left:        "1",
				LeftColor:   "black",
				Right:       "1",
				RightColor:  "black",
				Top:         "1",
				TopColor:    "black",
				Bottom:      "1",
				BottomColor: "black",
			},
			Font: xlsx.Font{
				Size: 6,
			},
			ApplyAlignment: true,
			Alignment: xlsx.Alignment{
				WrapText: true,
			},
		})
		if in([]int{0, 1}, i) {
			s.Cols[i].Width = 15
		}
		if in([]int{3, 7, 9, 10, 11, 12, 13, 14, 15}, i) {
			s.Cols[i].Width = 15
		}
	}
	//add values
	for _, v := range rep {
		r := s.AddRow()
		setValue(r, v.Date)
		setValue(r, v.JurName)
		setValue(r, v.DishGroupSecondParent)
		setValue(r, v.DishCode)
		setValue(r, v.DishName)
		setValue(r, v.PayTypes)
		setValue(r, v.DiscountType)
		setValue(r, v.DishAmountInt)
		setValue(r, v.CloseTime)
		setValue(r, v.OrderNum)
		setValue(r, v.DishSumInt)
		setValue(r, v.DiscountSum)
		setValue(r, v.DishDiscountSumInt)
		setValue(r, v.ProductCostBaseProductCost)
	}

	return f.Save(ReportName())
}

func in(array []int, index int) bool {
	for _, i := range array {
		if i == index {
			return true
		}
	}

	return false
}

func setValue(r *xlsx.Row, value interface{}) {
	c := r.AddCell()

	if v, ok := value.(float64); ok {
		c.SetFloatWithFormat(v, ".2f")
		return
	}

	c.SetValue(value)
}
