/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sanoyo/mini-redash-go/config"
	"github.com/sanoyo/mini-redash-go/db"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	maxconn     = 25
	maxLifetime = 5 * time.Minute
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect your data source",
	Long:  `Cobra is a CLI library for Go that empowers applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func Run() {
	config, err := config.InitConfig()
	if err != nil {
		errors.WithStack(err)
	}

	// init db setting
	db, err := db.InitDB(maxconn, maxLifetime, config.DB.CreateDSN())
	if err != nil {
		errors.WithStack(err)
	}
	// TODO: zap 使う
	fmt.Println("database connected")

	sql, err := ReadSQLFile("sample/sample.sql")
	if err != nil {
		errors.WithStack(err)
	}

	fmt.Println("sql", sql)

	// 1件の結果を返す
	mybook := MyBook{}
	rows, err := db.Query(sql)
	if err != nil {
		errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&mybook.ID, &mybook.Name)
		if err != nil {
			errors.WithStack(err)
		}
	}
	err = rows.Err()
	if err != nil {
		errors.WithStack(err)
	}

	fmt.Println("mybook", mybook)

	// TODO: FieldDescrptions使う
	// https://pkg.go.dev/github.com/jackc/pgx#Rows.FieldDescriptions
}

func ReadSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	b := bytes.NewBuffer(content)

	return b.String(), nil
}

type MyBook struct {
	ID   int
	Name string
}
