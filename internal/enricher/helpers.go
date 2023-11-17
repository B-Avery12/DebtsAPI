package enricher

import (
	"encoding/json"
	"errors"
	"github/bridger/interviews/DebtsAPI/internal/api"
	"log"
	"math"
	"strconv"
	"time"
)

const (
	customFormat = "2006-01-02"
)

// EnrichDebts - enrich debts based of enriched payment plans
func EnrichDebts(debts []Debt, paymentPlansByDebtID map[int]PaymentPlan) []Debt {
	enrichedDebts := make([]Debt, len(debts))
	for index, debt := range debts {
		if plan, ok := paymentPlansByDebtID[debt.ID]; ok {
			if plan.AmountToPay == 0 {
				debt.NextPaymentDueDate = ""
				debt.IsInPaymentPlan = false // Since the debt is paid mark as not being in a payment plan
			} else {
				debt.IsInPaymentPlan = true
				// Now we just need the NexPaymentDueDate - Its the next installment date after the last payment. Start at the start date and add installmen frequency until past the last payment then quit
				nextPaymentDate, err := getNextPayment(plan.StartDate, plan.InstallmentFrequency, plan.LastPayment.Date)
				if err != nil {
					log.Printf("unable to get nextPaymentDate with err: %s for debt: %d and plan: %d\n", err.Error(), debt.ID, plan.ID)
					continue
				}
				debt.NextPaymentDueDate = nextPaymentDate.Format(customFormat)
			}
			debt.RemainingAmount = json.Number(strconv.FormatFloat(plan.AmountToPay, 'e', 2, 64))
		} else {
			debt.NextPaymentDueDate = ""
			debt.RemainingAmount = json.Number(strconv.FormatFloat(debt.AmountOwed, 'e', 2, 64))
		}

		enrichedDebts[index] = debt
	}

	return enrichedDebts
}

// ConvertDebts - Convert from api model to enrich model
func ConvertDebts(apiDebts api.Debts) []Debt {
	debts := make([]Debt, len(apiDebts))
	for index, debt := range apiDebts {
		debts[index] = Debt{
			ID:              debt.ID,
			AmountOwed:      debt.AmountOwed,
			RemainingAmount: json.Number(strconv.FormatFloat(debt.AmountOwed, 'e', 2, 64)),
			IsInPaymentPlan: false,
		}
	}

	return debts
}

// SortPlansByDebtID -
func SortPlansByDebtID(plans []PaymentPlan) map[int]PaymentPlan {
	sortedPlans := make(map[int]PaymentPlan)
	for _, plan := range plans {
		sortedPlans[plan.DebtID] = plan
	}

	return sortedPlans
}

// EnrichPaymentPlans - enrich payment plans based on payments paid towards them
func EnrichPaymentPlans(plans []PaymentPlan, payments map[int][]Payment) []PaymentPlan {
	enrichedPlans := make([]PaymentPlan, len(plans))
	for index, plan := range plans {
		planPayments := payments[plan.ID]
		if len(planPayments) > 0 {
			lastPayment := planPayments[0]
			for _, payment := range planPayments {
				if payment.Date.After(lastPayment.Date) {
					lastPayment = payment
				}
				plan.AmountToPay -= payment.Amount
			}

			plan.LastPayment = &lastPayment
		} else {
			plan.LastPayment = nil
		}
		plan.AmountToPay = math.Round(plan.AmountToPay*100) / 100
		enrichedPlans[index] = plan
	}

	return enrichedPlans
}

// ConvertPaymentPlans - convert from api paymentplans to enrich paymentplan
func ConvertPaymentPlans(apiPlans api.PaymentPlans) []PaymentPlan {
	plans := make([]PaymentPlan, len(apiPlans))
	for index, plan := range apiPlans {
		startDate, err := time.Parse(customFormat, plan.StartDate)
		if err != nil {
			// As no start date is given default to year one
			startDate, err = time.Parse(customFormat, "0001-11-12")
		}
		newPlan := PaymentPlan{
			ID:                   plan.ID,
			DebtID:               plan.DebtID,
			AmountToPay:          plan.AmountToPay,
			InstallmentFrequency: plan.InstallmentFrequency,
			InstallmentAmount:    plan.InstallmentAmount,
			StartDate:            startDate,
		}
		plans[index] = newPlan
	}

	return plans
}

// ConvertPayments -
func ConvertPayments(apiPayments api.Payments) []Payment {
	payments := make([]Payment, len(apiPayments))
	for index, apiPayment := range apiPayments {
		paymentDate, err := time.Parse(customFormat, apiPayment.Date)
		if err != nil {
			// Set a custom date at year 1
			paymentDate, _ = time.Parse(customFormat, "0001-12-12")
		}

		payments[index] = Payment{
			PaymentPlanID: apiPayment.PaymentPlanID,
			Amount:        apiPayment.Amount,
			Date:          paymentDate,
		}
	}

	return payments
}

// SortPaymentsByPlan - sorts payments by plan they are paid towards
func SortPaymentsByPlan(payments []Payment) map[int][]Payment {
	sortedPayments := make(map[int][]Payment)
	for _, payment := range payments {
		if _, ok := sortedPayments[payment.PaymentPlanID]; !ok {
			sortedPayments[payment.PaymentPlanID] = make([]Payment, 0)
		}
		sortedPayments[payment.PaymentPlanID] = append(sortedPayments[payment.PaymentPlanID], payment)
	}

	return sortedPayments
}

func getNextPayment(startDate time.Time, installmentFrequency string, lastPayment time.Time) (time.Time, error) {
	var installmentDuration time.Duration
	switch installmentFrequency {
	case api.WeeklyFrequency:
		installmentDuration = time.Duration(time.Hour * 24 * 7)
	case api.BiWeeklyFrequency:
		installmentDuration = time.Duration(time.Hour * 24 * 14)
	default:
		return time.Now(), errors.New("Invalid installment frequency")
	}
	for !startDate.After(lastPayment) {
		startDate = startDate.Add(installmentDuration)
	}
	return startDate, nil
}
