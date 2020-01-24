package RecognitionService

import (
	"fmt"

	"github.com/SRkuma/design_pattern/Gateway"
	"github.com/SRkuma/design_pattern/Money"
	_ "github.com/go-sql-driver/mysql"
)

type RecognitionService struct {
	DB Gateway.Gateway
}

// TODO: asOf をdate型に変える
func (rs *RecognitionService) recognizedRevenue(contractNumber int64, asOf string) Money.Money {
	result := Money.Dollars(0)
	resultSet := rs.DB.FindRecognitionsFor(contractNumber, asOf)
	// resultset回してrevenueを足し上げする
	// リザルトを返す

	// エラー起きたらApplicetionExceptionを投げて死ぬ
	fmt.Println(asOf)
	return Money.Dollars(0)
}


func (rs *RecognitionService) calculateRevenueRecognitions(contractNumber int64) {
	contracts := rs.DB.FindContract(contractNumber)
	fmt.Println(contractNumber)
	return
}
