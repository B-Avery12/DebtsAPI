package enricher

import (
	"encoding/json"
	"github/bridger/interviews/DebtsAPI/internal/api"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertDebts(t *testing.T) {
	apiDebts := api.Debts{
		{
			ID:         0,
			AmountOwed: 1.23,
		},
		{
			ID:         1,
			AmountOwed: 2.23,
		},
		{
			ID:         2,
			AmountOwed: 3.23,
		},
		{
			ID:         3,
			AmountOwed: 4.23,
		},
	}

	actualDebts := ConvertDebts(apiDebts)

	assert.Equal(t, apiDebts[0].AmountOwed, actualDebts[0].AmountOwed)
	assert.Equal(t, apiDebts[0].ID, actualDebts[0].ID)

	assert.Equal(t, apiDebts[1].AmountOwed, actualDebts[1].AmountOwed)
	assert.Equal(t, apiDebts[1].ID, actualDebts[1].ID)

	assert.Equal(t, apiDebts[2].AmountOwed, actualDebts[2].AmountOwed)
	assert.Equal(t, apiDebts[2].ID, actualDebts[2].ID)

	assert.Equal(t, apiDebts[3].AmountOwed, actualDebts[3].AmountOwed)
	assert.Equal(t, apiDebts[3].ID, actualDebts[3].ID)
}

func TestConvertPlans(t *testing.T) {
	apiPlans := api.PaymentPlans{
		{
			ID:                   1,
			DebtID:               1,
			AmountToPay:          1.23,
			InstallmentFrequency: "BI_WEEKLY",
			InstallmentAmount:    .1,
			StartDate:            "2010-01-02",
		},
		{
			ID:                   2,
			DebtID:               2,
			AmountToPay:          1.23,
			InstallmentFrequency: "WEEKLY",
			InstallmentAmount:    .4,
			StartDate:            "2011-11-14",
		},
		{
			ID:                   3,
			DebtID:               3,
			AmountToPay:          1.23,
			InstallmentFrequency: "BI_WEEKLY",
			InstallmentAmount:    .15,
			StartDate:            "2012-06-16",
		},
	}

	time.Date(2010, time.January, 2, 0, 0, 0, 0, time.UTC)

	actualPlans := ConvertPaymentPlans(apiPlans)

	assert.Equal(t, apiPlans[0].ID, actualPlans[0].ID)
	assert.Equal(t, apiPlans[0].DebtID, actualPlans[0].DebtID)
	assert.Equal(t, apiPlans[0].AmountToPay, actualPlans[0].AmountToPay)
	assert.Equal(t, apiPlans[0].InstallmentFrequency, actualPlans[0].InstallmentFrequency)
	assert.Equal(t, apiPlans[0].InstallmentAmount, actualPlans[0].InstallmentAmount)
	assert.Equal(t, time.Date(2010, time.January, 2, 0, 0, 0, 0, time.UTC), actualPlans[0].StartDate)

	assert.Equal(t, apiPlans[1].ID, actualPlans[1].ID)
	assert.Equal(t, apiPlans[1].DebtID, actualPlans[1].DebtID)
	assert.Equal(t, apiPlans[1].AmountToPay, actualPlans[1].AmountToPay)
	assert.Equal(t, apiPlans[1].InstallmentFrequency, actualPlans[1].InstallmentFrequency)
	assert.Equal(t, apiPlans[1].InstallmentAmount, actualPlans[1].InstallmentAmount)
	assert.Equal(t, time.Date(2011, time.November, 14, 0, 0, 0, 0, time.UTC), actualPlans[1].StartDate)

	assert.Equal(t, apiPlans[2].ID, actualPlans[2].ID)
	assert.Equal(t, apiPlans[2].DebtID, actualPlans[2].DebtID)
	assert.Equal(t, apiPlans[2].AmountToPay, actualPlans[2].AmountToPay)
	assert.Equal(t, apiPlans[2].InstallmentFrequency, actualPlans[2].InstallmentFrequency)
	assert.Equal(t, apiPlans[2].InstallmentAmount, actualPlans[2].InstallmentAmount)
	assert.Equal(t, time.Date(2012, time.June, 16, 0, 0, 0, 0, time.UTC), actualPlans[2].StartDate)
}

func TestConvertPayments(t *testing.T) {
	apiPayments := api.Payments{
		{
			PaymentPlanID: 1,
			Amount:        1.12,
			Date:          "2020-02-23",
		},
		{
			PaymentPlanID: 2,
			Amount:        2.12,
			Date:          "2021-01-23",
		},
		{
			PaymentPlanID: 3,
			Amount:        3.12,
			Date:          "2019-05-23",
		},
	}

	actualPayments := ConvertPayments(apiPayments)
	assert.Equal(t, apiPayments[0].Amount, actualPayments[0].Amount)
	assert.Equal(t, apiPayments[0].PaymentPlanID, actualPayments[0].PaymentPlanID)
	assert.Equal(t, time.Date(2020, time.February, 23, 0, 0, 0, 0, time.UTC), actualPayments[0].Date)

	assert.Equal(t, apiPayments[1].Amount, actualPayments[1].Amount)
	assert.Equal(t, apiPayments[1].PaymentPlanID, actualPayments[1].PaymentPlanID)
	assert.Equal(t, time.Date(2021, time.January, 23, 0, 0, 0, 0, time.UTC), actualPayments[1].Date)

	assert.Equal(t, apiPayments[2].Amount, actualPayments[2].Amount)
	assert.Equal(t, apiPayments[2].PaymentPlanID, actualPayments[2].PaymentPlanID)
	assert.Equal(t, time.Date(2019, time.May, 23, 0, 0, 0, 0, time.UTC), actualPayments[2].Date)
}

func TestSortPlansByDebtID(t *testing.T) {
	lastPayment1 := Payment{
		PaymentPlanID: 1,
		Amount:        1.10,
		Date:          time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
	}
	unsortedPlans := []PaymentPlan{
		{
			ID:                   1,
			DebtID:               1,
			AmountToPay:          20.20,
			InstallmentFrequency: "WEEKLY",
			InstallmentAmount:    1.10,
			StartDate:            time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			LastPayment:          &lastPayment1,
		},
		{
			ID:                   2,
			DebtID:               2,
			AmountToPay:          40.40,
			InstallmentFrequency: "BI_WEEKLY",
			InstallmentAmount:    2.20,
			StartDate:            time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC),
			LastPayment:          nil,
		},
	}

	plansByDebitID := SortPlansByDebtID(unsortedPlans)

	assert.Equal(t, unsortedPlans[0].ID, plansByDebitID[unsortedPlans[0].DebtID].ID)
	assert.Equal(t, unsortedPlans[0].DebtID, plansByDebitID[unsortedPlans[0].DebtID].DebtID)
	assert.Equal(t, unsortedPlans[0].AmountToPay, plansByDebitID[unsortedPlans[0].DebtID].AmountToPay)
	assert.Equal(t, unsortedPlans[0].InstallmentFrequency, plansByDebitID[unsortedPlans[0].DebtID].InstallmentFrequency)
	assert.Equal(t, unsortedPlans[0].InstallmentAmount, plansByDebitID[unsortedPlans[0].DebtID].InstallmentAmount)
	assert.Equal(t, unsortedPlans[0].StartDate, plansByDebitID[unsortedPlans[0].DebtID].StartDate)
	assert.Equal(t, unsortedPlans[0].LastPayment, plansByDebitID[unsortedPlans[0].DebtID].LastPayment)

	assert.Equal(t, unsortedPlans[1].ID, plansByDebitID[unsortedPlans[1].DebtID].ID)
	assert.Equal(t, unsortedPlans[1].DebtID, plansByDebitID[unsortedPlans[1].DebtID].DebtID)
	assert.Equal(t, unsortedPlans[1].AmountToPay, plansByDebitID[unsortedPlans[1].DebtID].AmountToPay)
	assert.Equal(t, unsortedPlans[1].InstallmentFrequency, plansByDebitID[unsortedPlans[1].DebtID].InstallmentFrequency)
	assert.Equal(t, unsortedPlans[1].InstallmentAmount, plansByDebitID[unsortedPlans[1].DebtID].InstallmentAmount)
	assert.Equal(t, unsortedPlans[1].StartDate, plansByDebitID[unsortedPlans[1].DebtID].StartDate)
	assert.Nil(t, plansByDebitID[unsortedPlans[1].DebtID].LastPayment)
}

func TestSortPaymentsByPlanID(t *testing.T) {
	unsortedPayments := []Payment{
		{
			PaymentPlanID: 1,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			PaymentPlanID: 1,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			PaymentPlanID: 2,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			PaymentPlanID: 2,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			PaymentPlanID: 2,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			PaymentPlanID: 3,
			Amount:        1.1,
			Date:          time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
	}

	sortedPayments := SortPaymentsByPlan(unsortedPayments)

	assert.Equal(t, 3, len(sortedPayments))
	assert.Equal(t, 2, len(sortedPayments[1]))
	assert.Equal(t, 3, len(sortedPayments[2]))
	assert.Equal(t, 1, len(sortedPayments[3]))
}

func TestGetNextPayment(t *testing.T) {
	startDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	lastPayment := time.Date(2020, time.January, 7, 0, 0, 0, 0, time.UTC)

	nextPayment, err := getNextPayment(startDate, "WEEKLY", lastPayment)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC), nextPayment)

	lastPayment = time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC)
	nextPayment, err = getNextPayment(startDate, "WEEKLY", lastPayment)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC), nextPayment)

	lastPayment = time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC)
	nextPayment, err = getNextPayment(startDate, "WEEKLY", lastPayment)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC), nextPayment)

	lastPayment = time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC)
	nextPayment, err = getNextPayment(startDate, "BI_WEEKLY", lastPayment)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC), nextPayment)

	lastPayment = time.Date(2020, time.January, 16, 0, 0, 0, 0, time.UTC)
	nextPayment, err = getNextPayment(startDate, "BI_WEEKLY", lastPayment)
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2020, time.January, 29, 0, 0, 0, 0, time.UTC), nextPayment)

	lastPayment = time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC)
	nextPayment, err = getNextPayment(startDate, "wrong", lastPayment)
	assert.NotNil(t, err)
}

func TestEnrichPaymentPlans(t *testing.T) {
	unEnrichedPlans := []PaymentPlan{
		{
			ID:                   1,
			DebtID:               1,
			AmountToPay:          20.20,
			InstallmentFrequency: "WEEKLY",
			InstallmentAmount:    1.10,
			StartDate:            time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:                   2,
			DebtID:               2,
			AmountToPay:          40.40,
			InstallmentFrequency: "BI_WEEKLY",
			InstallmentAmount:    2.20,
			StartDate:            time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC),
		},
	}

	paymentsByPlanID := map[int][]Payment{
		1: {
			{
				PaymentPlanID: 1,
				Amount:        10.20,
				Date:          time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
			},
			{
				PaymentPlanID: 1,
				Amount:        10.00,
				Date:          time.Date(2020, time.January, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		2: {
			{
				PaymentPlanID: 2,
				Amount:        40.00,
				Date:          time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	enrichedPlans := EnrichPaymentPlans(unEnrichedPlans, paymentsByPlanID)

	assert.Equal(t, len(unEnrichedPlans), len(enrichedPlans))
	assert.Equal(t, 0.0, enrichedPlans[0].AmountToPay)
	assert.Equal(t, 0.4, enrichedPlans[1].AmountToPay)
	assert.Equal(t, time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC), enrichedPlans[0].LastPayment.Date)
	assert.Equal(t, time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC), enrichedPlans[1].LastPayment.Date)
}

func TestEnrichDebts(t *testing.T) {
	lastPayment1 := Payment{
		PaymentPlanID: 1,
		Amount:        1.10,
		Date:          time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
	}
	lastPayment2 := Payment{
		PaymentPlanID: 2,
		Amount:        1.10,
		Date:          time.Date(2020, time.January, 16, 0, 0, 0, 0, time.UTC),
	}
	plans := map[int]PaymentPlan{
		1: {
			ID:                   1,
			DebtID:               1,
			InstallmentFrequency: "WEEKLY",
			StartDate:            time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			LastPayment:          &lastPayment1,
			AmountToPay:          1.4,
		},
		2: {
			ID:                   2,
			DebtID:               2,
			InstallmentFrequency: "BI_WEEKLY",
			StartDate:            time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			LastPayment:          &lastPayment2,
			AmountToPay:          4.4,
		},
		3: {
			ID:                   3,
			DebtID:               3,
			InstallmentFrequency: "BIWEEKLY",
			StartDate:            time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			LastPayment:          &lastPayment1,
			AmountToPay:          0,
		},
	}
	debts := []Debt{
		{
			ID:         1,
			AmountOwed: 20.5,
		},
		{
			ID:         2,
			AmountOwed: 20.5,
		},
		{
			ID:         3,
			AmountOwed: 5.5,
		},
		{
			ID:         4,
			AmountOwed: 10.5,
		},
	}

	enrichedDebts := EnrichDebts(debts, plans)

	assert.Equal(t, true, enrichedDebts[0].IsInPaymentPlan)
	assert.Equal(t, "2020-01-15", enrichedDebts[0].NextPaymentDueDate)
	assert.Equal(t, json.Number(strconv.FormatFloat(1.4, 'e', 2, 64)), enrichedDebts[0].RemainingAmount)

	assert.Equal(t, true, enrichedDebts[1].IsInPaymentPlan)
	assert.Equal(t, "2020-01-29", enrichedDebts[1].NextPaymentDueDate)
	assert.Equal(t, json.Number(strconv.FormatFloat(4.4, 'e', 2, 64)), enrichedDebts[1].RemainingAmount)

	assert.Equal(t, false, enrichedDebts[2].IsInPaymentPlan)
	assert.Equal(t, "", enrichedDebts[2].NextPaymentDueDate)
	assert.Equal(t, json.Number(strconv.FormatFloat(0.0, 'e', 2, 64)), enrichedDebts[2].RemainingAmount)

	assert.Equal(t, false, enrichedDebts[3].IsInPaymentPlan)
	assert.Equal(t, "", enrichedDebts[3].NextPaymentDueDate)
	assert.Equal(t, json.Number(strconv.FormatFloat(debts[3].AmountOwed, 'e', 2, 64)), enrichedDebts[3].RemainingAmount)
}
