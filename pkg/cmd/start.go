package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/ziczhu/fibonacci_rest_api/pkg/api"
	"github.com/ziczhu/fibonacci_rest_api/pkg/config"
)

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "start fibonacci service",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()

		conf := config.InitConfig()

		if conf.Env == "prod" {
			gin.SetMode(gin.ReleaseMode)
		}

		router.GET("/api/v1/fibonacci/:n", api.GetFibonacciSequence(conf))

		router.Run(fmt.Sprintf(":%d", conf.Port))
	},
}
