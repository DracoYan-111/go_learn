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
	"errors"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration file related operations",
	Long: `
Here are configuration file related operations, including:
`,
}

var getConfigAll = &cobra.Command{
	Use:   "getAll",
	Short: "查看所有配置文件内容",
	Long:  figure.NewFigure("Nothing", "smslant", false).String(),
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range viper.AllSettings() {
			fmt.Printf("%s=%v\n", k, v)
		}
	},
}

var getConfig = &cobra.Command{
	Use:   "get",
	Short: "查看指定的配置内容",
	Long: `
	get string 查询指定配置项,多个使用空格分隔
	`,
	Example: `get KEY KEY...:多个使用空格分隔`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//检查是否设置了传入参数
		if len(args) > 0 {
			for i, v := range args {
				if viper.Get(v) != nil {
					fmt.Printf("%s=%v\n", args[i], viper.GetString(v))
				} else {
					return errors.New("该配置项不存在")
				}
			}
		} else {
			return errors.New("请传入至少一个需要查询的配置项,多个使用空格分隔")
		}
		return nil
	},
}

var setConfig = &cobra.Command{
	Use:   "set",
	Short: "修改指定的配置内容",
	Long: `
	set string string 设置指定配置项
	`,
	Example: `set KEY VALUE:设置指定配置项,必须出现仅一组KEY VALUE两个参数`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//检查是否设置了传入参数
		fmt.Println(args)
		if len(args) == 2 {
			if viper.Get(args[0]) != nil {
				// 设置配置文件
				viper.Set(args[0], args[1])
				// 写入配置文件
				viper.WriteConfig()
				fmt.Println("修改配置文件成功")

			} else {
				return errors.New("该配置项不存在")
			}
		} else {
			return errors.New("请传入至少一组KEY VALUE")
		}
		return nil
	},
}

var addConfig = &cobra.Command{
	Use:   "add",
	Short: "添加配置项",
	Long: `
	add -key -value 添加配置项
	`,
	Example: `
	add -key -value :添加配置项,必须出现仅一组KEY VALUE`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 检查 --key 和 --value 是否被使用，并且必须提供参数
		keyFlag, _ := cmd.Flags().GetBool("key")
		valueFlag, _ := cmd.Flags().GetBool("value")
		if keyFlag && valueFlag {
			viper.Set(args[0], args[1])
			viper.WriteConfig()
			fmt.Println("添加配置文件成功")
		} else {
			return errors.New("至少包含key和value两个参数")
		}
		return nil
	},
}

var delConfig = &cobra.Command{
	Use:   "del",
	Short: "删除指定的配置内容",
	Long: `
	del string 删除指定配置项
	`,
	Example: `del KEY:删除指定配置项`,
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFlag, _ := cmd.Flags().GetBool("key")
		if keyFlag {
			if viper.Get(args[0]) != nil {
				deleteConfigCmd(args[0])
				fmt.Println("删除配置文件成功")
			} else {
				return errors.New("该配置项不存在")
			}
		} else {
			return errors.New("请传入至少一个需要删除的配置项")
		}

		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(getConfigAll)
	ConfigCmd.AddCommand(getConfig)
	ConfigCmd.AddCommand(setConfig)
	ConfigCmd.AddCommand(addConfig)
	ConfigCmd.AddCommand(delConfig)

	// 添加标志来显示中文帮助。
	getConfig.Flags().BoolP("help", "h", false, "show help information and view detailed descriptions of all available commands")

	addConfig.Flags().BoolP("key", "k", false, "config配置项KEY")
	addConfig.Flags().BoolP("value", "v", false, "config配置项VALUE")

	delConfig.Flags().BoolP("key", "k", false, "config配置项KEY")
}

// deleteConfigCmd represents the deleteConfig command
func deleteConfigCmd(key string) error {

	newAllConfig := make(map[string]any, len(viper.AllSettings()))
	for k, v := range viper.AllSettings() {
		newAllConfig[k] = v
	}

	delete(newAllConfig, key)

	files, err := os.Create(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	defer files.Close()
	viper.Reset()

	for k, v := range newAllConfig {
		if _, ok := v.(string); !ok {
			return fmt.Errorf("config value for key %q is not a string", k)
		}

		viper.Set(k, v)
		if _, err := files.WriteString(k + "=" + v.(string) + "\n"); err != nil {
			return err
		}
	}

	return viper.WriteConfig()
}
