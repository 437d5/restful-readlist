// Package db contain CRUD operations to be called from handlers
package db

import (
	"context"
	"fmt"
	"github.com/437d5/restful-readlist/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	Client *pgxpool.Pool
}

func (c *Controller) Create(ctx context.Context, data models.Book) error {
	query := `INSERT INTO books (done, author, title, year_published, rating)
	VALUES (@done, @author, @title, @yearPublished, @rating);`
	args := pgx.NamedArgs{
		"done":          data.Done,
		"author":        data.Author,
		"title":         data.Title,
		"yearPublished": data.YearPublished,
		"rating":        data.Rating,
	}
	_, err := c.Client.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (c *Controller) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM books WHERE id = $1;`
	_, err := c.Client.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}
	return nil
}

func (c *Controller) Update(ctx context.Context, id int, data models.Book) error {
	query := `UPDATE books
				SET done = @done, author = @author, title = @title, 
				    year_published = @year_published, rating = @rating
				WHERE id = @id;`

	args := pgx.NamedArgs{
		"done":          data.Done,
		"author":        data.Author,
		"title":         data.Title,
		"yearPublished": data.YearPublished,
		"rating":        data.Rating,
		"id":            id,
	}

	_, err := c.Client.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func (c *Controller) GetByID(ctx context.Context, id int) (rawData models.Book, err error) {
	query := `SELECT FROM books WHERE id = @id;`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := c.Client.Query(ctx, query, args)
	if err != nil {
		return models.Book{}, fmt.Errorf("unable to get row: %w", err)
	}
	defer rows.Close()

	rawDataArr, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Book, error) {
		var (
			pk             int
			done           bool
			author         string
			title          string
			year_published int
			rating         int
		)
		err := row.Scan(&pk, &done, &author, &title, &year_published, &rating)
		if err != nil {
			return models.Book{}, err
		}
		return models.Book{
			Done:          done,
			Author:        author,
			Title:         title,
			YearPublished: year_published,
			Rating:        rating,
		}, nil
	})
	rawData = rawDataArr[0]
	return rawData, nil
}

func (c *Controller) Get(ctx context.Context) (rawData []models.Book, err error) {
	query := `SELECT * FROM books;`

	rows, err := c.Client.Query(ctx, query)
	if err != nil {
		return []models.Book{{}}, fmt.Errorf("unable to get row: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		book := models.Book{}
		var pk int

		err := rows.Scan(&pk, &book.Done, &book.Author, &book.Title, &book.YearPublished, &book.Rating)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		rawData = append(rawData, book)
	}

	return rawData, nil
}
