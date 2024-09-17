package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dpcamargo/fullcycle-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Name",
		Description: sql.NullString{String: "Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Before Update ---")
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          categories[0].ID,
		Name:        "New Name",
		Description: sql.NullString{String: "New Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- After Update ---")
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, categories[0].ID)
	if err != nil {
		panic(err)
	}

	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- After Delete ---")
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
}
