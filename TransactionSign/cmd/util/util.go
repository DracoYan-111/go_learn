/*
right © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"hash/fnv"
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

// 单位转换
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

// 地址上色
var colorAddress = &cobra.Command{
	Use:   "color",
	Short: "给地址添加唯一的颜色",
	Long: figure.NewFigure("Color", "larry3d", true).String() +
		`
=====给地址上色，可以看到地址的唯一颜色表示，也可对比两个地址有个不同
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		checkAddressColor("0xa40c00E83a70243Cb8F2A7B0ce907d619F7f9ea3")
		return nil
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
	UtilCmd.AddCommand(colorAddress)

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
		_, err := new(big.Rat).SetString(s)
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

	// 将number转换为Rat
	inputNubmerArt, _ := new(big.Rat).SetString(number)

	// 将单位转换为Rat
	inPutUnitsValueBigInt, _ := new(big.Int).SetString(unitMultipliers[uints], 10)
	unitValueRat := new(big.Rat).SetInt(inPutUnitsValueBigInt)

	// 将Rat相乘
	numberInWeiRat := new(big.Rat).Mul(inputNubmerArt, unitValueRat)

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

// ========== 地址检查 ===============
// 使用 FNV 哈希函数来生成颜色
func charToColor(c byte) string {
	h := fnv.New32a()
	h.Write([]byte{c})
	hash := h.Sum32()

	// 基于 hash 值生成 RGB 颜色
	r := (hash & 0xFF0000) >> 16
	g := (hash & 0x00FF00) >> 8
	b := hash & 0x0000FF

	// ANSI 256色格式 "\033[38;2;R;G;Bm"
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// 为输入的地址生成特有的颜色
func checkAddressColor(input string) {
	var result string
	for i := 0; i < len(input); i++ {
		color := charToColor(input[i])
		result += fmt.Sprintf("%s%c\033[0m", color, input[i])
	}
	fmt.Println(result)
}

func compareAndPrintStrings(str1, str2 string) {
	// 两个字符串长度必须相等
	if len(str1) != len(str2) {
		fmt.Println("字符串长度不一致")
		return
	}

	// 用于保存输出的两行
	var line1, line2 string

	// 逐字符对比两个字符串
	for i := 0; i < len(str1); i++ {
		char1 := str1[i]
		char2 := str2[i]

		if char1 != char2 {
			// 对不同的字符添加颜色
			color1 := charToColor(char1)
			color2 := charToColor(char2)
			line1 += fmt.Sprintf("%s%c\033[0m", color1, char1)
			line2 += fmt.Sprintf("%s%c\033[0m", color2, char2)
		} else {
			// 相同字符直接添加到两行中
			line1 += fmt.Sprintf("%c", char1)
			line2 += fmt.Sprintf("%c", char2)
		}
	}

	// 打印两行字符串
	fmt.Println(line1)
	fmt.Println(line2)
}
