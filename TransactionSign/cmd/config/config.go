/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "c或cf即可此处查看config相关信息",
	Long:    `这里是config相关的详细信息`,
	Aliases: []string{"c", "cf"},
	Example: `
	--getAll 查看所有存在的配置项
	--set [key] 设置key对应的配置项
	--delete [key] 删除key对应的配置项
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var GetConfigAllCmd = &cobra.Command{
	Use:     "getAll",
	Short:   "g或ga即可此处查看config所有信息",
	Long:    `这里是config相关的详细信息`,
	Aliases: []string{"g", "ga"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	fmt.Println("config called=========")

	ConfigCmd.AddCommand(GetConfigAllCmd)

}
