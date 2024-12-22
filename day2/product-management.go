package main

import (
	"errors"
	"fmt"
)

type Category string

const (
	Eletronicos Category = "Eletronicos"
	Livros      Category = "Livros"
	Vestuario   Category = "Vestuario"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Category Category
}

func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("o preço deve ser maior que zero")
	}
	return nil
}

func ValidateCategory(category Category) error {
	switch category {
	case Eletronicos, Livros, Vestuario:
		return nil
	default:
		return errors.New("categoria inválida")
	}
}

func AddProduct(products *[]Product, product Product) error {
	if err := ValidatePrice(product.Price); err != nil {
		return err
	}
	if err := ValidateCategory(product.Category); err != nil {
		return err
	}
	*products = append(*products, product)
	return nil
}

func GetProductsByCategory(products []Product, category Category) []Product {
	var result []Product
	for _, product := range products {
		if product.Category == category {
			result = append(result, product)
		}
	}
	return result
}

func ListProducts(products []Product) {
	for _, product := range products {
		fmt.Printf("ID: %s, Nome: %s, Preço: R$ %.2f, Categoria: %s\n", product.ID, product.Name, product.Price, product.Category)
	}
}

func main() {
	products := []Product{}

	err := AddProduct(&products, Product{
		ID:       "1",
		Name:     "Smartphone",
		Price:    1500.0,
		Category: Eletronicos,
	})
	if err != nil {
		fmt.Println("Erro ao adicionar produto:", err)
	}

	err = AddProduct(&products, Product{
		ID:       "2",
		Name:     "Livro Golang",
		Price:    40.0,
		Category: Livros,
	})
	if err != nil {
		fmt.Println("Erro ao adicionar produto:", err)
	}

	fmt.Println("Todos os produtos:")
	ListProducts(products)

	fmt.Println("\nProdutos na categoria Livros:")
	Livros := GetProductsByCategory(products, Livros)
	ListProducts(Livros)
}
