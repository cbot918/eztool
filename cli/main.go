package main

import (
	"cli/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "et",
	Short: "A brief description of your application",
	Long:  `eztool cli`,
}

// login
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登入",
	Long:  `login to eztool cloud`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		api.Login()
	},
}

// // 命令一
// var mockMsgCmd = &cobra.Command{
// 	Use:   "mockMsg",
// 	Short: "批次傳送測試文字訊息",
// 	Long:  ``,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("mockMsg called")
// 	},
// }

// // 命令二
// var exportCmd = &cobra.Command{
// 	Use:   "export",
// 	Short: "匯出資料",
// 	Long:  ``,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("export called")
// 	},
// }

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// rootCmd.AddCommand(mockMsgCmd)
	// rootCmd.AddCommand(exportCmd)
	// exportCmd.Flags().StringP("out", "k", "./backup", "匯出路徑")

	rootCmd.AddCommand(loginCmd)

}

func main() {
	Execute()
}
