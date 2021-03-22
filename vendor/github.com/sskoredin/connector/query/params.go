package query

import (
	"crypto/sha1"
	"fmt"
	"github.com/sskoredin/connector/domain/entity"
	"net/url"
	"strings"
)

type Params struct {
	params []string
	url    string
}

func New(link, token string) Params {
	var p Params
	if token != "" {
		if strings.Contains(link, "9900") || strings.Contains(link, "biz") || strings.Contains(link, "9901") {
			p = p.token(token)
		} else {
			p = p.key(token)
		}
	}
	p.url = link
	return p
}

func (p Params) Build() (Query, error) {
	u, err := url.Parse(p.url)
	if err != nil {
		return "", err
	}
	u.Scheme = "http"
	q := u.Query()
	for _, s := range p.params {
		v := strings.Split(s, "=")
		q.Add(v[0], v[1])
	}

	u.RawQuery = q.Encode()
	res := u.String()

	return Query(res), nil
}

func (p Params) ReportType(v string) Params {
	p.params = append(p.params, fmt.Sprintf("report=%s", v))
	return p
}

func (p Params) Period(period entity.ReportPeriod) Params {
	p.params = append(p.params, fmt.Sprintf("from=%s", period.Start))
	p.params = append(p.params, fmt.Sprintf("to=%s", period.End))
	return p
}

func (p Params) IikoBizPeriod(period entity.ReportPeriod) Params {
	p.params = append(p.params, fmt.Sprintf("dateFrom=%s", period.Start))
	p.params = append(p.params, fmt.Sprintf("dateTo=%s", period.End))
	return p
}

func (p Params) IikoBizTransactionPeriod(period entity.ReportPeriod) Params {
	p.params = append(p.params, fmt.Sprintf("date_from=%s", period.Start))
	p.params = append(p.params, fmt.Sprintf("date_to=%s", period.End))
	return p
}

func (p Params) GroupRows(v ...string) Params {
	for _, s := range v {
		p.params = append(p.params, fmt.Sprintf("groupRow=%s", s))
	}
	return p
}

func (p Params) Args(v ...string) Params {
	for _, s := range v {
		p.params = append(p.params, fmt.Sprintf("arg=%s", s))
	}
	return p
}

func (p Params) Organization(s string) Params {
	p.params = append(p.params, fmt.Sprintf("organization=%s", s))
	return p
}

func (p Params) OrganizationID(s string) Params {
	p.params = append(p.params, fmt.Sprintf("organizationID=%s", s))
	return p
}

func (p Params) CustomerId(s string) Params {
	p.params = append(p.params, fmt.Sprintf("customerId=%s", s))
	return p
}

func (p Params) CategoryId(s string) Params {
	p.params = append(p.params, fmt.Sprintf("categoryId=%s", s))
	return p
}

func (p Params) WalletId(s string) Params {
	p.params = append(p.params, fmt.Sprintf("walletId=%s", s))
	return p
}

func (p Params) Sum(s string) Params {
	p.params = append(p.params, fmt.Sprintf("sum=%s", s))
	return p
}

func (p Params) ID(s string) Params {
	p.params = append(p.params, fmt.Sprintf("id=%s", s))
	return p
}

func (p Params) Login(s string) Params {
	p.params = append(p.params, fmt.Sprintf("login=%s", s))
	return p
}

func (p Params) Password(s string) Params {
	s = passHash(s)
	p.params = append(p.params, fmt.Sprintf("password=%s", s))
	return p
}

func (p Params) Timeout(s string) Params {
	p.params = append(p.params, fmt.Sprintf("request_timeout=%s", s))
	return p
}

func (p Params) RequestTimeout(s string) Params {
	p.params = append(p.params, fmt.Sprintf("requestTimeout=%s", s))
	return p
}

func (p Params) UserID(s string) Params {
	p.params = append(p.params, fmt.Sprintf("user_id=%s", s))
	return p
}

func (p Params) UserSecret(s string) Params {
	p.params = append(p.params, fmt.Sprintf("user_secret=%s", s))
	return p
}

func (p Params) key(v string) Params {
	p.params = append(p.params, fmt.Sprintf("key=%s", v))
	return p
}

func (p Params) token(v string) Params {
	p.params = append(p.params, fmt.Sprintf("access_token=%s", v))
	return p
}

func passHash(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
