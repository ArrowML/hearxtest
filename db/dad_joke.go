package db

import (
	"context"
	"database/sql"
	"fmt"
	"hearxtest/model"
	"strconv"
)

type PostgresDadJokeRepository struct {
	DB        *sql.DB
	TableName string
}

func (pg PostgresDadJokeRepository) Save(ctx context.Context, js *[]model.DadJoke) (int64, error) {

	sqlStr := fmt.Sprintf("INSERT INTO %s (joke, punchline, rating) VALUES ", pg.TableName)

	values := []interface{}{}
	for i, joke := range *js {
		values = append(values, joke.Joke, joke.Punchline, joke.Rating)
		numFields := 3
		n := i * numFields

		sqlStr += `(`
		for j := 0; j < numFields; j++ {
			sqlStr += `$` + strconv.Itoa(n+j+1) + `,`
		}
		sqlStr = sqlStr[:len(sqlStr)-1] + `),`
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	res, err := pg.DB.Exec(sqlStr, values...)
	if err != nil {
		return -1, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (pg PostgresDadJokeRepository) FetchPage(ctx context.Context, page, limit int) (*[]model.DadJoke, error) {

	offset := (page - 1) * limit
	sqlStr := fmt.Sprintf("SELECT joke, punchline, rating FROM %s ORDER BY id ASC LIMIT %d OFFSET %d", pg.TableName, limit, offset)

	rows, err := pg.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var js []model.DadJoke
	for rows.Next() {
		var j model.DadJoke
		err = rows.Scan(&j.Joke, &j.Punchline, &j.Rating)
		if err != nil {
			return nil, err
		}
		js = append(js, j)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	if len(js) == 0 {
		return nil, sql.ErrNoRows
	}

	return &js, nil
}

func (pg PostgresDadJokeRepository) FetchJoke(ctx context.Context, id int) (model.DadJoke, error) {

	sqlStr := fmt.Sprintf("SELECT joke, punchline, rating FROM %s WHERE id = %d", pg.TableName, id)

	var j model.DadJoke
	row := pg.DB.QueryRow(sqlStr)
	err := row.Scan(&j.Joke, &j.Punchline, &j.Rating)
	if err != nil {
		return j, err
	}
	return j, nil
}

func (pg PostgresDadJokeRepository) FetchAllIDs(ctx context.Context) ([]int, error) {

	sqlStr := fmt.Sprintf("SELECT id FROM %s", pg.TableName)

	rows, err := pg.DB.Query(sqlStr)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	err = rows.Err()
	if err != nil {
		return ids, err
	}
	if len(ids) == 0 {
		return ids, sql.ErrNoRows
	}
	return ids, nil
}
