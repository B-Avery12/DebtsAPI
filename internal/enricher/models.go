package enricher

import (
	"encoding/json"
	"time"
)

const (
	// WeeklyFrequency - Represents weekly (7 days) payment frequency
	WeeklyFrequency = "WEEKLY"
	// BiWeeklyFrequency - Represents bi weekly (14 days) payment frequency
	BiWeeklyFrequency = "BI_WEEKLY"
)

// Debt - Money owed to a collector.
type Debt struct {
	ID                 int         `json:"id"`
	AmountOwed         float64     `json:"amount"`                          // As we need to represent how many cents are in the amount we need a float type
	RemainingAmount    json.Number `json:"remaining_amount"`                // Amount left to pay off debt
	NextPaymentDueDate string      `json:"next_payment_due_date,omitempty"` // ISO 8601 UTC date of next payment. Next installment date after lastest payment (installment dates are start date + installment frequency)
	IsInPaymentPlan    bool        `json:"is_in_payment_plan"`
}

// PaymentPlan - Payment plan for paying of a Debt
type PaymentPlan struct {
	ID                   int
	DebtID               int     // Debt this payment plan is linked. 1 to 1 relationship with debt
	AmountToPay          float64 // Amount (USD) needed to resolve this payment plan
	InstallmentFrequency string
	InstallmentAmount    float64   // Amount due each installment
	StartDate            time.Time // ISO 8601 date of when the first payment due
	LastPayment          *Payment
}

// Payment - payment towards a payment plan
type Payment struct {
	PaymentPlanID int       // Payment Plan this payment goes towards. Many to 1 relationship
	Amount        float64   // Amount
	Date          time.Time // ISO 8601 date of when payment occurred
}
