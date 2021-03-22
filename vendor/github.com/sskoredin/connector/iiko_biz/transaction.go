package iiko_biz_client

import (
	entity "github.com/sskoredin/connector/domain/entity"
	"github.com/sskoredin/connector/query"
	"time"
)

func (s *Service) GetTransactionsByOrganizationAndPeriod(organizationID, customerID string, start, end time.Time) ([]entity.Transaction, error) {
	period := entity.ReportPeriod{
		Start: start.Format(entity.DateLayout),
		End:   end.Format(entity.DateLayout),
	}
	q, err := query.New(s.linkWithValue(organizationTransactionReport, organizationID), s.token).
		IikoBizTransactionPeriod(period).
		UserID(customerID).
		Build()

	if err != nil {
		return nil, err
	}

	var transactions []entity.Transaction
	if err := q.GetWithAttemptsInStruct(transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
