package converter

import (
	"fmt"
	"github.com/sskoredin/iiko_report/client"
	"github.com/tealeg/xlsx"
	"sort"
	"time"
)

func Convert(responseData *client.ResponseData) Report {
	if responseData == nil {
		return nil
	}
	values := make([]ReportValue, len(responseData.R))
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
			CloseTime:                  closeTime,
			OrderNum:                   v.OrderNum,
			DishSumInt:                 v.DishSumInt,
			DiscountSum:                v.DiscountSum,
			DishDiscountSumInt:         v.DishDiscountSumInt,
			ProductCostBaseProductCost: v.ProductCostBaseProductCost,
		}
		values[i] = value
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].CloseTime.Before(values[j].CloseTime)
	})

	return values
}

func ReportName(start, end string) string {
	now := time.Now()
	year, month, _ := now.Date()
	if len(end) > 0 {
		e, _ := time.Parse("02.01.2006", end)
		month = e.Month()
		year = e.Year()
	}
	return fmt.Sprintf("./%d.%d.xlsx", month, year)
}

func (rep Report) ToXlsx(start, end string) error {
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
	for i := range s.Cols {
		if in([]int{2, 4, 6, 8}, i) {
			s.Cols[i].Width = 25
		}
		if in([]int{0, 1, 3, 7, 9, 10, 11, 12, 13, 14, 15}, i) {
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
		setValue(r, v.CloseTime.Format("02.01.2006 15:04:05"))
		setValue(r, v.OrderNum)
		setValue(r, v.DishSumInt)
		setValue(r, v.DiscountSum)
		setValue(r, v.DishDiscountSumInt)
		setValue(r, v.ProductCostBaseProductCost)
	}

	return f.Save(ReportName(start, end))
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

	c.SetValue(value)
}
