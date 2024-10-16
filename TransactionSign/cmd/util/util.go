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
	"golang.org/x/crypto/sha3"
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

// 单位转换 https://converter.murkin.me/
var ethereumConverter = &cobra.Command{
	Use:   "ethereum",
	Short: "ethereum 单位转换器",
	Long: figure.NewFigure("Converter", "larry3d", true).String() +
		`
=====这是将单位转换为各个以太单位的工具=====
`,
	Example: `
输入需要转换的数量和单位，例如以下的某一个单位:
-n number -u wei
-n number -u kwei
-n number -u mwei
-n number -u gwei
-n number -u szabo
-n number -u finney
-n number -u ether
-n number -u kether
-n number -u mether
-n number -u gether
-n number -u tether

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ethereum 单位转换器")
		errors := checkInput(args)
		return errors
	},
}

var colorAddressColor string

// 地址上色 https://eth-colored-address.dnevend.site/
var colorAddress = &cobra.Command{
	Use:   "color",
	Short: "给地址添加唯一的颜色",
	Long: figure.NewFigure("Color", "larry3d", true).String() +
		`
=====给地址上色，可以看到地址的唯一颜色表示，也可对比两个地址有个不同
`,
	Example: `
--address/-a:为地址赋予唯一的颜色
--left/-l:需要对比的左侧地址
--right/-r:需要对比的右侧地址
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// 为地址赋予唯一的颜色
		if colorAddressColor != "" {
			err := checkAddressColor(colorAddressColor)
			return err
		}

		// 地址区别对比
		left, _ := cmd.Flags().GetBool("left")
		right, _ := cmd.Flags().GetBool("right")

		if left {
			if right {
				err := compareAndPrintStrings(args[0], args[1])
				return err
			} else {
				return errors.New("请传入右侧需要对比的地址")
			}
		} else {
			return errors.New("请传入左侧需要对比的地址")
		}
	},
}

var keccak256Data string

// keccak256
var keccak256Cmd = &cobra.Command{
	Use:   "keccak256",
	Short: "keccak256 哈希函数",
	Long: figure.NewFigure("Keccak256", "larry3d", true).String() +
		`
=====keccak256 哈希函数=====
`,
	Example: `
--data/-a:需要计算的数据
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("keccak256 哈希函数")
		if keccak256Data == "" {
			return errors.New("请传入需要计算的数据")
		}
		errors := solidityKeccak256(keccak256Data)
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
	UtilCmd.AddCommand(colorAddress)
	UtilCmd.AddCommand(keccak256Cmd)

	ethereumConverter.Flags().BoolP("number", "n", false, "转换数量")
	ethereumConverter.Flags().BoolP("uint", "u", false, "数量单位")
	ethereumConverter.Flags().BoolP("help", "h", false, "显示帮助信息")

	colorAddress.Flags().StringVarP(&colorAddressColor, "address", "a", "", "要上色的地址")
	colorAddress.Flags().BoolP("left", "l", false, "需要对比的左侧地址")
	colorAddress.Flags().BoolP("right", "r", false, "需要对比的右侧地址")

	keccak256Cmd.Flags().StringVarP(&keccak256Data, "data", "d", "", "要计算的数据")
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
func checkAddressColor(input string) error {
	if len(input) != 42 {
		return errors.New("请输入正确的地址")
	}
	var result string
	for i := 0; i < len(input); i++ {
		color := charToColor(input[i])
		result += fmt.Sprintf("%s%c\033[0m", color, input[i])
	}
	fmt.Println(result)

	return nil
}

// 对比传入地址的区别并上色
func compareAndPrintStrings(str1, str2 string) error {
	// 两个字符串长度必须相等
	if len(str1) != len(str2) {
		return errors.New("字符串长度不一致")
	}

	if len(str1) != 42 && len(str2) != 42 {
		return errors.New("请检查地址的格式是否正确")
	}

	// 用于保存输出的两行
	var line1, line2 string
	var difference bool
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
			difference = true
		} else {
			// 相同字符直接添加到两行中
			line1 += fmt.Sprintf("%c", char1)
			line2 += fmt.Sprintf("%c", char2)
		}
	}

	// 打印两行字符串与是否存在区别
	fmt.Println("是否存在区别:", difference)
	fmt.Println("左侧地址:", line1)
	fmt.Println("右侧地址:", line2)
	return nil
}

// ========== keccak_256 ===============
func solidityKeccak256(data string) error {
	dataByte := []byte(data)
	hash := sha3.NewLegacyKeccak256() // 使用 Keccak 版本的 SHA3
	hash.Write(dataByte)
	fmt.Printf("Keccak-256 hash: %x\n", hash.Sum(nil))
	return nil
}
