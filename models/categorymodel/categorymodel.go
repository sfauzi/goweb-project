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

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUE(?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInserId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInserId > 0

}

func Detail(id int) entities.Category {

	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}
