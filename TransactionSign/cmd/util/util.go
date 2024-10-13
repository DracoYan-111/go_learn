/*
right © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// utilCmd represents the util command
var UtilCmd = &cobra.Command{
	Use:   "util",
	Short: "这里是工具相关命令",
	Long: figure.NewFigure("Util", "larry3d", true).String() +
		`
=====这里有常用的区块链工具=====
`,
}
var unitMultipliers map[string]string

var ethereumConverter = &cobra.Command{
	Use:   "ethereum",
	Short: "ethereum 单位转换器",
	Long: figure.NewFigure("Converter", "larry3d", true).String() +
		`
=====这是将单位转换为各个以太单位的工具=====
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ethereum 单位转换器")
		errors := checkInput(args)
		return errors
	},
}

func init() {
	// 定义单位精度
	unitMultipliers = map[string]string{
		"wei":    "1",
		"kwei":   "1000",
		"mwei":   "1000000",
		"gwei":   "1000000000",
		"szabo":  "1000000000000",
		"finney": "1000000000000000",
		"ether":  "1000000000000000000",
		"kether": "1000000000000000000000",
		"mether": "1000000000000000000000000",
		"gether": "1000000000000000000000000000",
		"tether": "1000000000000000000000000000000",
	}
	UtilCmd.AddCommand(ethereumConverter)
	ethereumConverter.Flags().BoolP("number", "n", false, "转换数量")
	ethereumConverter.Flags().BoolP("uint", "u", false, "数量单位")
	ethereumConverter.Flags().BoolP("help", "h", false, "显示帮助信息")
}

// ========== 单位转换 ===============
// 输入检查逻辑
func checkInput(args []string) error {

	if len(args) != 2 {
		return errors.New("请输入数量和单位")
	}

	if isNumber := func(s string) bool {
		_, err := new(big.Int).SetString(s, 10)
		return err
	}; !isNumber(args[0]) {
		return errors.New("请输入正确的数量")
	}

	if isUint := func(s string) bool {
		_, ok := unitMultipliers[s]
		return ok
	}; !isUint(args[1]) {
		return errors.New("请输入正确的单位")
	}

	ethNumberConverter(args[0], args[1])
	return nil
}

// 单位转换逻辑
func ethNumberConverter(number, uints string) {
	// 定义单位列表
	uintsOrder := []string{"wei", "kwei", "mwei", "gwei", "szabo", "finney", "ether", "kether", "mether", "gether", "tether"}

	// 判断number是否合法
	inputNubmerBigInt, _ := new(big.Int).SetString(number, 10)

	// 判断uints是否合法
	inPutUnitsValueBigInt, _ := new(big.Int).SetString(unitMultipliers[uints], 10)

	// 将输入number和单位精度相乘并转为bigint
	numberInWei := new(big.Int).Mul(inputNubmerBigInt, inPutUnitsValueBigInt)
	// 将bigint转为rat准备精准运算
	numberInWeiRat := new(big.Rat).SetInt(numberInWei)

	for _, uintName := range uintsOrder {
		// 获取每个单位的精度
		uintNameValueBigInt, _ := new(big.Int).SetString(unitMultipliers[uintName], 10)
		// 将每个单位的精度相除
		result := new(big.Rat).Quo(numberInWeiRat, new(big.Rat).SetInt(uintNameValueBigInt)).FloatString(30)

		//  检查result是否包含".""
		if strings.Contains(result, ".") {
			result = strings.TrimRight(result, "0")
			result = strings.TrimRight(result, ".")
		}

		fmt.Printf("%-7s: %s\n", uintName, result)
	}
}

// =================================
