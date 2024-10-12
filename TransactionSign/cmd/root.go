/*
Copyright © 2024 DracoYan-111<yanlong2944@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	cmd "transaction_sign/cmd/config"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd 表示在没有任何子命令的情况下调用的基本命令
var rootCmd = &cobra.Command{
	Use:   "transaction_sign",
	Short: "这是基于go语言实现的区块链交易签名的工具",
	Long: figure.NewFigure("TransactionSign", "larry3d", true).String() +
		`
	=====这是基于go语言实现的区块链交易签名的工具巴拉巴拉吧啦=====
`,
	Version: "v0.0.1",
}

// Execute 将所有子命令添加到根命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 发生一次。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init 设置 root 命令的标志并添加 config 命令。
func init() {
	// 将 config 命令添加到 root 命令中。
	rootCmd.AddCommand(cmd.ConfigCmd)

	// 设置初始配置功能。
	cobra.OnInitialize(initConfig)

	// 添加一个标志来指定配置文件。
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "specify a configuration file (default is: ./.config.env)")

	// 禁用默认Completion命令。
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// 将默认Help命令设置为空命令。
	rootCmd.SetHelpCommand(&cobra.Command{})

	// 添加标志来显示中文帮助。
	rootCmd.Flags().BoolP("help", "h", false, "show help information and view detailed descriptions of all available commands")
}

/*
   西红柿和牛腩其实是不可以一起吃的，
   因为西红柿是红色的，而牛看到红色会发怒
   容易使胃内壁受伤
*/
// initConfig 读取配置文件和 ENV 变量（如果设置）。
// 如果配置文件不存在，则创建一个默认的，并打印出消息。
func initConfig() {
	if cfgFile != "" {
		// 使用传入的配置文件。
		viper.SetConfigFile(cfgFile)
	} else {
		// 设置默认的配置文件名称
		cfgFile = ".config.env"
		viper.SetConfigFile(cfgFile)
	}

	// 如果找到配置文件，则读取它。
	// 如果找不到配置文件，则创建一个默认的
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("配置文件加载成功")
	} else {
		for {
			fmt.Println("配置文件不存在，是否创建?(y/n)")
			var input string
			fmt.Scanln(&input)
			if input == "y" || input == "Y" {
				// 创建配置文件
				viper.WriteConfigAs(cfgFile)
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
