package main

import (
	"encoding/json"
	"fmt"
	"github/bridger/interviews/TrueAccord/internal/api"
	"github/bridger/interviews/TrueAccord/internal/enricher"
	"log"
)

func main() {
	// Get all debts, plans, and payments
	debts, err := api.GetDebts()
	if err != nil {
		log.Fatal("unable to get debts")
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

	enrichedPlans := enricher.CreatePaymentPlans(plans, payments)
	enrichedDebts := enricher.EnrichDebts(debts, enrichedPlans)

	for _, debt := range enrichedDebts {
		rawDebt, err := json.Marshal(debt)
		if err != nil {
			log.Printf("unable to marshal debt %d", debt.ID)
		}
		fmt.Println(rawDebt)
	}
}
