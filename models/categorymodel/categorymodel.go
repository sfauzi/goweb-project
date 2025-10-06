package categorymodel

import (
	"goweb-project/config"
	"goweb-project/entities"
	"log"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT id, name, created_at, updated_at FROM categories`)
	if err != nil {
		log.Fatalf("Error querying categories: %v", err)
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		// ✅ Gunakan pointer & untuk Scan
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			log.Fatalf("Error scanning category: %v", err)
		}
		categories = append(categories, category)
	}

	return categories
}
