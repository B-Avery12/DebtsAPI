package api

const (
	// WeeklyFrequency - Represents weekly (7 days) payment frequency
	WeeklyFrequency = "WEEKLY"
	// BiWeeklyFrequency - Represents bi weekly (14 days) payment frequency
	BiWeeklyFrequency = "BI_WEEKLY"
)

// Debt - Money owed to a collector.
type Debt struct {
	ID         int     `json:"id"`
	AmountOwed float64 `json:"amount"` // As we need to represent how many cents are in the amount we need a float type
}

// Debts -
type Debts []Debt

// PaymentPlan - Payment plan for paying of a Debt
type PaymentPlan struct {
	ID                   int     `json:"id"`
	DebtID               int     `json:"debt_id"`       // Debt this payment plan is linked. 1 to 1 relationship with debt
	AmountToPay          float64 `json:"amount_to_pay"` // Amount (USD) needed to resolve this payment plan
	InstallmentFrequency string  `json:"installment_frequency,omitempty"`
	InstallmentAmount    float64 `json:"installment_amount"` // Amount due each installment
	StartDate            string  `json:"start_date"`         // ISO 8601 date of when the first payment due
}

// PaymentPlans -
type PaymentPlans []PaymentPlan

// Payment - Payment towards a payment plan
type Payment struct {
	PaymentPlanID int     `json:"payment_plan_id"` // Payment Plan this payment goes towards. Many to 1 relationship
	Amount        float64 `json:"amount"`          // Amount
	Date          string  `json:"date"`            // ISO 8601 date of when payment occurred
}

// Payments -
type Payments []Payment
