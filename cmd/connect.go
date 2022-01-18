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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/olekukonko/tablewriter"

	"github.com/sanoyo/mini-redash-go/config"
	"github.com/sanoyo/mini-redash-go/db"
	"github.com/sanoyo/mini-redash-go/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	maxconn     = 1
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
	// init zap logging
	log.InitLogger()

	config, err := config.InitConfig()
	if err != nil {
		errors.WithStack(err)
	}

	// init db setting
	db, err := db.InitDB(maxconn, maxLifetime, config.DB.CreateDSN())
	if err != nil {
		errors.WithStack(err)
	}
	log.Logger.Info("database connected")

	// TODO: -f オプションとかでファイルを渡せるようにする
	sql, err := readSQLFile("sample/sample.sql")
	if err != nil {
		errors.WithStack(err)
	}

	rows, err := db.Query(context.TODO(), sql)
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
				// ret += string(strconv.Itoa(v.(int))) + " "
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
	_, ids, err = delete_int(ids, 0)
	if err != nil {
		errors.WithStack(err)
	}

	data := make([][]string, len(tempStr))
	for i, v := range data {
		v = []string{fmt.Sprint(ids[i]), tempStr[i]}
		data = append(data, v)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func readSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	b := bytes.NewBuffer(content)

	return b.String(), nil
}

func delete_int(slice []int, s int) (int, []int, error) {
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
