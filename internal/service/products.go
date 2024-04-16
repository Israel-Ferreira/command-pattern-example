package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/models"
)

func GenerateRandomReleaseDate() *time.Time {
	now := time.Now()

	random_qty_hours := int64(rand.Uint32()) * 24

	now = now.Add(time.Duration(random_qty_hours) * time.Hour)

	return &now
}

func LoadProducts() []models.Product {
	products := []models.Product{
		{Sku: "12334", ModelName: "Barbeador Phillips Smart", ReleaseDate: GenerateRandomReleaseDate()},
		{Sku: "929292", ModelName: "Aspirador de Pó Robo Eletrolux", ReleaseDate: GenerateRandomReleaseDate()},
		{Sku: "020020", ModelName: "TV OLED LG 81 Polegadas Smart C/ WebOS", ReleaseDate: GenerateRandomReleaseDate()},
	}

	return products
}

func FindProductBySKU(sku string) (models.Product, error) {
	products := LoadProducts()

	var foundProduct models.Product

	for _, product := range products {
		if product.Sku == sku {
			foundProduct = product
			break
		}
	}

	if foundProduct.Sku == "" {
		return foundProduct, errors.New("produto não encontrado")
	}

	return foundProduct, nil

}
