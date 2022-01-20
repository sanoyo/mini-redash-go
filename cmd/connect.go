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
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"

	"github.com/sanoyo/mini-redash-go/config"
	"github.com/sanoyo/mini-redash-go/db"
	"github.com/sanoyo/mini-redash-go/log"
	"github.com/sanoyo/mini-redash-go/query"
	"github.com/sanoyo/mini-redash-go/view"

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
		flag, err := cmd.PersistentFlags().GetString("file")
		if err != nil {
			log.Logger.Error("file is empty")
		}
		execute(flag)
	},
}

func init() {
	// init logging
	log.InitLogger()

	// init config
	config.InitConfig()

	// init db
	err := db.InitDB(maxconn, maxLifetime, config.Config.DB.CreateDSN())
	if err != nil {
		os.Exit(0)
	}
	log.Logger.Info("database connected")

	connectCmd.PersistentFlags().String("file", "", "file option")
	rootCmd.AddCommand(connectCmd)
}

func execute(flag string) {
	// ファイルをもとに解析する
	header, data, err := query.Analyse(flag)
	if err != nil {
		errors.WithStack(err)
	}

	// 結果を表示させる
	view.Show(header, data)
}
