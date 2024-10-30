/*
Copyright © 2024 DracoYan-111 <yanlong2944@gmail.com>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

const (
	ARB  = "https://arb1.arbitrum.io/rpc"
	ETH  = "https://1rpc.io/eth"
	BNB  = "https://rpc.ankr.com/bsc"
	BASE = "https://mainnet.base.org"
	OP   = "https://mainnet.optimism.io"
	POL  = "https://polygon.llamarpc.com"
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("transaction called")
		err := processContent()
		return err
		//transactionUtil()
	},
}

var NetWork string
var Private string
var To string
var Amount string
var Uints string
var Nonce string
var GasLimit string
var GasPrice string
var Data string

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.Flags().StringVarP(&NetWork, "netWork", "N", "", "需要使用的网络")
	transactionCmd.Flags().StringVarP(&Private, "private", "p", "", "需要使用的私钥")
	transactionCmd.Flags().StringVarP(&To, "to", "t", "", "需要使用的目标地址")
	transactionCmd.Flags().StringVarP(&Amount, "amount", "a", "", "需要使用的数量")
	transactionCmd.Flags().StringVarP(&Uints, "uints", "u", "", "需要使用的单位")
	transactionCmd.Flags().StringVarP(&Nonce, "nonce", "n", "", "需要使用的nonce")
	transactionCmd.Flags().StringVarP(&GasLimit, "gasLimit", "g", "", "需要使用的gasLimit")
	transactionCmd.Flags().StringVarP(&GasPrice, "gasPrice", "G", "", "需要使用的gasPrice")
	transactionCmd.Flags().StringVarP(&Data, "data", "d", "", "需要使用的data")
}

// 定义单位精度
var unitMultipliers = map[string]string{
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

type Transaction struct {
	To       *common.Address
	Amount   *big.Int
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
	Data     []byte
}

//netWork string, private string, to string,nonce string,amount *big.Int ,gasLimit uint64,gasPrice *big.Int,data []byte

func processContent() error {
	var transaction Transaction
	if NetWork == "" || Private == "" || To == "" || Amount == "" {
		fmt.Println("网络、私钥、目标地址、数量不能为空")
	} else {

		// 判断网络
		switch NetWork {
		case "ETH":
			NetWork = ETH
		case "ARB":
			NetWork = ARB
		case "BNB":
			NetWork = BNB
		case "BASE":
			NetWork = BASE
		case "OP":
			NetWork = OP
		case "POL":
			NetWork = POL
		default:
			if !strings.HasPrefix(NetWork, "https://") {
				return errors.New("请检查网络的格式是否正确")
			}
		}
		client, err := ethclient.Dial(NetWork)
		if err != nil {
			return errors.New("请检查网络是否正确")
		}
		if _, err := client.ChainID(context.Background()); err != nil {
			return errors.New("请检查网络是否正确")
		}

		// 判断私钥
		Private = strings.TrimLeft(Private, "0x")
		privateKey, err := crypto.HexToECDSA(Private)
		if err != nil {
			return errors.New("请检查私钥的格式是否正确")
		}
		fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

		// 判断目标地址
		if len(To) != 42 {
			return errors.New("请检查地址的格式是否正确")
		} else {
			if transaction.To == nil {
				transaction.To = new(common.Address)
			}
			*transaction.To = common.HexToAddress(To)
		}

		// 判断数量是否正确并转为rat
		amountRat, bol := new(big.Rat).SetString(Amount)
		if !bol {
			return errors.New("请检查数量的格式是否正确")
		}

		// 判断是否存在单位
		if Uints != "" {
			if _, ok := unitMultipliers[Uints]; !ok {
				return errors.New("请检查单位的格式是否正确")
			} else {
				if Uints == "wei" {
					if strings.HasPrefix(Amount, "0.") {
						return errors.New("单位为wei时不可以为小数")
					}
				}
				// 获取输入的单位精度
				uintsInt, _ := new(big.Int).SetString(unitMultipliers[Uints], 10)
				// 将单位精度转换为Rat
				uintsRat, _ := new(big.Rat).SetString(uintsInt.String())
				// 将number与单位精度相乘
				numberInWeiRat := new(big.Rat).Mul(amountRat, uintsRat)

				// 将wei单位的精度相除
				uintsIntWei, _ := new(big.Int).SetString(unitMultipliers["wei"], 10)
				result := new(big.Rat).Quo(numberInWeiRat, new(big.Rat).SetInt(uintsIntWei)).FloatString(30)
				//  检查result是否包含".""如何包含将去除末尾的0与.
				if strings.Contains(result, ".") {
					result = strings.TrimRight(result, "0")
					result = strings.TrimRight(result, ".")
				}
				transaction.Amount, _ = new(big.Int).SetString(result, 10)
			}
		} else {
			if strings.HasPrefix(Amount, "0.") {
				return errors.New("单位为wei时不可以为小数")
			}
			amounts, _ := new(big.Int).SetString(Amount, 10)
			transaction.Amount = amounts
		}

		// 判断nonce是否存在
		if Nonce != "" {
			nonce, _ := new(big.Int).SetString(Nonce, 10)
			transaction.Nonce = nonce.Uint64()
		} else {
			transaction.Nonce, _ = client.PendingNonceAt(context.Background(), fromAddress)
		}

		// 判断data是否存在
		if Data != "" {
			transaction.Data = []byte(Data)
		} else {
			transaction.Data = nil
		}

		// 判断gasprice是否存在
		if GasPrice != "" {
			gasPrice, _ := new(big.Int).SetString(GasPrice, 10)
			transaction.GasPrice = gasPrice
		} else {
			transaction.GasPrice, _ = client.SuggestGasPrice(context.Background())

			// 判断gaslimit是否存在
			if GasLimit != "" {
				gasLimit, _ := new(big.Int).SetString(GasLimit, 10)
				transaction.GasLimit = gasLimit.Uint64()
			} else {
				callMsg := ethereum.CallMsg{
					From:     fromAddress,
					To:       transaction.To,
					GasPrice: transaction.GasPrice,
					Value:    transaction.Amount,
					Data:     transaction.Data,
				}

				estimateGas, err := client.EstimateGas(context.Background(), callMsg)
				if err != nil {
					return errors.New("获取gaslimit失败")
				}

				transaction.GasLimit = estimateGas
			}
		}
		fmt.Println("交易信息如下:")
		fmt.Println("交易发起者->", fromAddress)
		fmt.Println("目标地址->", *transaction.To)
		fmt.Println("交易数量->", transaction.Amount, "wei")
		fmt.Println("交易nonce->", transaction.Nonce)
		fmt.Println("交易gasprice->", transaction.GasPrice)
		fmt.Println("交易gaslimit->", transaction.GasLimit)
		fmt.Println("交易数据->", Data)
	}

	transactionUtil(NetWork, Private, &transaction)
	return nil
}

func transactionUtil(netWork string, private string, transaction *Transaction) error {
	// 1. 连接到以太坊节点 (可以使用 Infura 或本地节点)
	client, err := ethclient.Dial(netWork)
	if err != nil {
		return errors.New("连接以太坊节点失败")
	}

	// 2. 发送方的私钥
	privateKeyAvailable := strings.TrimLeft(private, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyAvailable)
	if err != nil {
		return errors.New("获取私钥失败")
	}

	// // 3. 获取发送方地址
	// fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// // 4. 获取nonce
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	return errors.New("获取nonce失败")
	// }

	// 5. 设置交易参数
	// transaction := &Transaction{
	// 	Nonce:    nonce,
	// 	To:       common.HexToAddress(to),
	// 	Amount:   big.NewInt(10000000000000),
	// 	GasPrice: big.NewInt(10000000000000),
	// 	GasLimit: 2200000,
	// 	Data:     []byte("HELLO WORLD!"),
	// }

	tx := types.NewTransaction(transaction.Nonce, *transaction.To, transaction.Amount, transaction.GasLimit, transaction.GasPrice, transaction.Data)

	// 6. 创建交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return errors.New("获取网络ID失败")
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return errors.New("签名交易失败")
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return errors.New("发送交易失败")
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	return nil
}
