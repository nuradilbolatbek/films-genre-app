package repository

import (
	"FilmsProject"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type GenreFilmsPostgres struct {
	db *sqlx.DB
}

func NewGenreFilmsPostgres(db *sqlx.DB) *GenreFilmsPostgres {
	return &GenreFilmsPostgres{db: db}
}

func (r *GenreFilmsPostgres) Create(listId int, item FilmsProject.GenreFilms) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (tittle, description) values ($1, $2) RETURNING id", filmsListsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", genreFilmsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *GenreFilmsPostgres) GetAll(userId, listId int) ([]FilmsProject.GenreFilms, error) {
	var items []FilmsProject.GenreFilms
	query := fmt.Sprintf(`SELECT ti.id, ti.tittle, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		filmsListsTable, genreFilmsTable, userListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *GenreFilmsPostgres) GetById(userId int, itemId int) (FilmsProject.GenreFilms, error) {
	var item FilmsProject.GenreFilms
	query := fmt.Sprintf(`SELECT ti.id, ti.tittle, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		filmsListsTable, genreFilmsTable, userListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *GenreFilmsPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		filmsListsTable, genreFilmsTable, userListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *GenreFilmsPostgres) Update(userId, itemId int, input FilmsProject.UpdateItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		filmsListsTable, setQuery, genreFilmsTable, userListsTable, argId, argId+1)

	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
