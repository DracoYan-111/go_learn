/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd 表示在没有任何子命令的情况下调用的基本命令
var RootCmd = &cobra.Command{
	Use:   "transaction_sign",
	Short: "交易签名的命令行工具",
	Long:  `使用go语言的Cobra框架实现的交易签名的命令行工具`,
	//Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute 将所有子命令添加到根命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 RootCmd 发生一次。
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 您将在此定义标志和配置设置。
	// Cobra 支持持久标志，如果在此定义，
	// 对您的应用程序而言将是全局的。
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "任意命令带config即可查看config路径")
	// 对help的介绍修改为中文。
	RootCmd.Flags().BoolP("help", "h", false, "显示帮助信息，并查看所有可用命令的详细描述")

	// 禁用默认的Completion命令
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	// 禁用默认的Help命令
	RootCmd.SetHelpCommand(&cobra.Command{}) // 使用一个空命令替换 help 命令

}

// initConfig 读取配置文件和 ENV 变量（如果设置）。
func initConfig() {
	if cfgFile != "" {
		// 使用标志中的配置文件。
		viper.SetConfigFile(cfgFile)
	} else {

		// // 在主目录中搜索名为“.config”的配置。
		viper.AddConfigPath("./")
		viper.SetConfigType("env")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // 读取匹配的环境变量

	// 如果找到配置文件，则读取它。
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "配置文件加载成功")
	} else {
		for {
			fmt.Println("配置文件不存在，是否创建?(y/n)")

			var input string
			fmt.Scanln(&input)

			if input == "y" || input == "Y" {
				// 创建配置文件
				viper.WriteConfigAs("./.config.env")
				fmt.Println("配置文件创建成功")
				break
			} else if input == "n" || input == "N" {
				os.Exit(0)
				fmt.Println("取消创建")
			} else {
				fmt.Println("输入错误")
				continue
			}

		}
	}
}
