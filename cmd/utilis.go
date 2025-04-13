package cmd

import (
	"fmt"

	"github.com/Mario-Valente/kiwify-webhoock/internal/models"
)

func formatAbandoned(p models.Abandoned, index int) string {
	return fmt.Sprintf(
		"%d. 🧾 Nome: %s\n📧 Email: %s\n💰 Phone: %s \n🌍 Pais: %s\n\n",
		index,
		p.Name,
		p.Email,
		p.Phone,
		p.Country,
	)
}

func formatPurchase(p models.Purchase, index int) string {
	return fmt.Sprintf(
		"%d. 🧾 Nome: %s\n📧 Email: %s\n💰 Phone: %s \n🌍 Estado: %s\n\n",
		index,
		p.Customer.FullName,
		p.Customer.Email,
		p.Customer.Mobile,
		p.Customer.State,
	)
}
