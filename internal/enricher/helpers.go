package enricher

import (
	"encoding/json"
	"github/bridger/interviews/TrueAccord/internal/api"
	"strconv"
	"time"
)

const (
	customFormat = "2006-01-02"
)

// EnrichDebts - convert from api model to enricher model
func EnrichDebts(apiDebts api.Debts, paymentPlansByDebtID map[int]PaymentPlan) map[int]Debt {
	debts := make(map[int]Debt, len(apiDebts))
	for index, debt := range apiDebts {
		newDebt := Debt{
			ID:              debt.ID,
			AmountOwed:      debt.AmountOwed,
			RemainingAmount: json.Number(strconv.FormatFloat(debt.AmountOwed, 'e', 2, 64)),
			IsInPaymentPlan: false,
		}
		if plan, ok := paymentPlansByDebtID[debt.ID]; ok {
			newDebt.IsInPaymentPlan = true
			newDebt.RemainingAmount = json.Number(strconv.FormatFloat(plan.AmountToPay, 'e', 2, 64))
			// Now we just need the NexPaymentDueDate - Its the next installment date after the last payment. Start at the start date and add installmen frequency until past the last payment then quit
			if plan.AmountToPay == 0 {
				newDebt.NextPaymentDueDate = ""
				newDebt.IsInPaymentPlan = false // Since the debt is paid mark as not being in a payment plan
			}
		} else {
			newDebt.NextPaymentDueDate = ""
		}

		debts[index] = newDebt
	}

	return debts
}

// CreatePaymentPlans - Converts from api model to enricher model and enriches with last payment, and update amount to pay based on payments total
func CreatePaymentPlans(apiPlans api.PaymentPlans, apiPayments api.Payments) map[int]PaymentPlan {
	paymentsByPlanID := createPayments(apiPayments)
	plans := make(map[int]PaymentPlan, len(apiPlans))
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
			Payments:             paymentsByPlanID[plan.ID],
		}
		if len(newPlan.Payments) > 0 {
			lastPayment := newPlan.Payments[0]
			for _, payment := range newPlan.Payments {
				if payment.Date.After(lastPayment.Date) {
					lastPayment = payment
				}
			}

			newPlan.LastPayment = &lastPayment
		} else {
			newPlan.LastPayment = nil
		}
		for _, payment := range newPlan.Payments {
			newPlan.AmountToPay = newPlan.AmountToPay - payment.Amount
		}
		plans[index] = newPlan
	}

	return plans
}

// Converts api payments to a enrich  models and creates a map not a slice
func createPayments(apiPayments api.Payments) map[int][]Payment {
	payments := make(map[int][]Payment)
	for _, apiPayment := range apiPayments {
		paymentDate, err := time.Parse(customFormat, apiPayment.Date)
		if err != nil {
			// Set a custom date at year 1
			paymentDate, _ = time.Parse(customFormat, "0001-12-12")
		}

		payment := Payment{
			PaymentPlanID: apiPayment.PaymentPlanID,
			Amount:        apiPayment.Amount,
			Date:          paymentDate,
		}

		if _, ok := payments[payment.PaymentPlanID]; !ok {
			payments[payment.PaymentPlanID] = make([]Payment, 0)
		}
		payments[payment.PaymentPlanID] = append(payments[payment.PaymentPlanID], payment)
	}

	return payments
}
