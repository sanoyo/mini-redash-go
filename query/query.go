package query

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/pkg/errors"
	"github.com/sanoyo/mini-redash-go/db"
)

func Analyse(file string) ([]string, [][]string, error) {
	sql, err := readSQLFile(file)
	if err != nil {
		errors.WithStack(err)
	}

	rows, err := db.DB.Query(context.TODO(), sql)
	if err != nil {
		errors.WithStack(err)
	}
	defer rows.Close()

	var header []string
	var tempStr []string
	once := sync.Once{}
	for rows.Next() {
		once.Do(
			func() {
				// https://pkg.go.dev/github.com/jackc/pgx#Rows.FieldDescriptions
				fds := rows.FieldDescriptions()

				for _, fd := range fds {
					column := string(fd.Name)
					header = append(header, column)
				}
			},
		)

		values, err := rows.Values()
		if err != nil {
			errors.WithStack(err)
		}

		for _, v := range values {
			switch v := v.(type) {
			case string:
				tempStr = append(tempStr, v)
			case int:
			}
		}
	}

	err = rows.Err()
	if err != nil {
		errors.WithStack(err)
	}

	ids := make([]int, len(tempStr))
	for i := range ids {
		ids = append(ids, i+1)
	}
	_, ids, err = deleteDuplicateInt(ids, 0)
	if err != nil {
		errors.WithStack(err)
	}

	data := make([][]string, len(tempStr))
	for i, v := range data {
		v = []string{fmt.Sprint(ids[i]), tempStr[i]}
		data = append(data, v)
	}

	return header, data, nil
}

func readSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	b := bytes.NewBuffer(content)

	return b.String(), nil
}

func deleteDuplicateInt(slice []int, s int) (int, []int, error) {
	ret := make([]int, len(slice))
	i := 0
	for _, x := range slice {
		if s != x {
			ret[i] = x
			i++
		}
	}
	if len(ret[:i]) == len(slice) {
		return 0, slice, fmt.Errorf("Couldn't find")
	}
	return s, ret[:i], nil
}
