package query

import (
	"fmt"
	iiko_client "github.com/sskoredin/connector/period"
)

type queryParams struct {
	params []string
}

func New(token string) *queryParams {
	p := new(queryParams)
	p.key(token)
	return p
}

func (p queryParams) ReportType(v string) queryParams {
	p.params = append(p.params, fmt.Sprintf("report=%s", v))
	return p
}
func (p queryParams) Period(period iiko_client.ReportPeriod) queryParams {
	p.params = append(p.params, fmt.Sprintf("from=%s", period.Start))
	p.params = append(p.params, fmt.Sprintf("to=%s", period.End))
	return p
}
func (p queryParams) GroupRows(v ...string) queryParams {
	for _, s := range v {
		p.params = append(p.params, fmt.Sprintf("groupRow=%s", s))
	}
	return p
}
func (p queryParams) Args(v ...string) queryParams {
	for _, s := range v {
		p.params = append(p.params, fmt.Sprintf("arg=%s", s))
	}
	return p
}

func (p queryParams) key(v string) queryParams {
	p.params = append(p.params, fmt.Sprintf("key=%s", v))
	return p
}
