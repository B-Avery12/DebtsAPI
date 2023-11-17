package main

import (
	"encoding/json"
	"fmt"
	"github/bridger/interviews/DebtsAPI/internal/api"
	"github/bridger/interviews/DebtsAPI/internal/enricher"
	"log"
)

func main() {
	// Get all debts, plans, and payments
	debts, err := api.GetDebts()
	if err != nil {
		log.Fatal("unable to get debts")
		return
	}
	if len(debts) <= 0 {
		log.Println("no debts, exiting")
		return
	}
	plans, err := api.GetPaymentPlans()
	if err != nil {
		log.Fatal("unable to get payment plans")
	}
	payments, err := api.GetPayments()
	if err != nil {
		log.Fatal("unable to get payment")
	}

	convertedPlans := enricher.ConvertPaymentPlans(plans)
	convertedPayments := enricher.ConvertPayments(payments)
	convertedDebts := enricher.ConvertDebts(debts)

	paymentsByPlan := enricher.SortPaymentsByPlan(convertedPayments)
	enrichedplans := enricher.EnrichPaymentPlans(convertedPlans, paymentsByPlan)
	sortedPlans := enricher.SortPlansByDebtID(enrichedplans)

	enrichedDebts := enricher.EnrichDebts(convertedDebts, sortedPlans)

	for _, debt := range enrichedDebts {
		rawDebt, err := json.Marshal(debt)
		if err != nil {
			log.Printf("unable to marshal debt %d", debt.ID)
		}
		fmt.Println(string(rawDebt))
	}
}
