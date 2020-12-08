package client

import (
	"crypto/sha1"
	"encoding/xml"
	_ "encoding/xml"
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (s Service) Collect() (*ResponseData, error) {
	err := s.readConfig()
	if err != nil {
		return nil, err
	}

	token, err := s.auth()
	if err != nil {
		return nil, err
	}

	return s.report(token, start(), end())
}

func (s Service) auth() (string, error) {
	params := make(map[string]string, 2)
	params["login"] = s.config.User
	params["pass"] = passHash(s.config.Password)

	resp, err := get(s.queryBuilder("/resto/api/auth", params))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(string(resp))
	if len(result) > 45 {
		return "", errors.New("Not authorized ")
	}
	return fmt.Sprintf(string(resp)), nil
}

func (s Service) queryBuilder(link string, params interface{}) string {
	host := fmt.Sprintf("%s%s", s.config.API, link)
	u, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "http"
	q := u.Query()
	if pp, ok := params.([]string); ok {
		for _, s := range pp {
			v := strings.Split(s, "=")
			q.Add(v[0], v[1])
		}
	}
	if pp, ok := params.(map[string]string); ok {
		for k, v := range pp {
			q.Set(k, v)
		}
	}

	u.RawQuery = q.Encode()
	query := u.String()
	s.logger.Println(query)
	fmt.Println(query)
	return query
}

func passHash(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func get(query string) ([]byte, error) {
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (s Service) report(token string, start, end string) (*ResponseData, error) {

	params := make([]string, 0)
	params = append(params, fmt.Sprintf("key=%s", token))
	params = append(params, "report=SALES")
	params = append(params, fmt.Sprintf("from=%s", start))
	params = append(params, fmt.Sprintf("to=%s", end))
	params = append(params, "groupRow=JurName")
	params = append(params, "groupRow=DishGroup.SecondParent")
	params = append(params, "groupRow=DishCode")
	params = append(params, "groupRow=DishName")
	params = append(params, "groupRow=PayTypes")
	params = append(params, "groupRow=OrderDiscount.Type")
	params = append(params, "agr=DishAmountInt")
	params = append(params, "groupRow=CloseTime")
	params = append(params, "groupRow=OrderNum")
	params = append(params, "agr=DishSumInt")
	params = append(params, "agr=DiscountSum")
	params = append(params, "agr=DishDiscountSumInt")
	params = append(params, "agr=ProductCostBase.ProductCost")

	d, err := get(s.queryBuilder("/resto/api/reports/olap", params))
	if err != nil {
		return nil, err
	}

	result := new(ResponseData)

	if err := xml.Unmarshal(d, result); err != nil {
		return nil, err
	}

	return result, nil
}

func start() string {
	return now.BeginningOfMonth().Format("02.01.2006")
}

func end() string {
	return now.BeginningOfDay().Format("02.01.2006")
}
