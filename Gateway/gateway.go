package Gateway

import (
	"database/sql"
	"time"

	"github.com/SRkuma/design_pattern/Money"
)

type Gateway struct {
	DB sql.DB
}

type Contract struct {
	ID         int64
	Product    int64
	Revenue    float64
	DateSigned time.Time
}

type RevenueRecognitions struct {
	Contract     int64
	Amount       float64
	RecognizedOn time.Time
}

//// TODO: asOf to date
func (g *Gateway) FindRecognitionsFor(contractID int64, asOf string) (*[]RevenueRecognitions, error) {
	// findrecognitionstatemet渡してpreparedStatement作る
	findRecognitionsStatement := `
SELECT amount
FROM
revenueRecognitions
WHERE contract = ?
AND recognizedOn <= ?;
`
	rows, err := g.DB.Query(findRecognitionsStatement, contractID, asOf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		contract     int64
		amount       float64
		recognizedOn time.Time
	)

	var recognitions []RevenueRecognitions

	for rows.Next() {
		if err := rows.Scan(&contract, &amount, &recognizedOn); err != nil {
			return nil, err
		}
		tmp := RevenueRecognitions{
			Contract:     contract,
			Amount:       amount,
			RecognizedOn: recognizedOn,
		}

		recognitions = append(recognitions, tmp)
	}

	return &recognitions, nil
}

func (g *Gateway) FindContract(contractID int64) (*[]Contract, error) {

	findContractStatement := `
SELECT * FROM
contracts c,
product p
WHERE ID = ?
AND c.product = p.ID;
`
	rows, err := g.DB.Query(findContractStatement, contractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id, product int64
		revenue     float64
		dateSigned  time.Time
	)

	var contracts []Contract

	for rows.Next() {
		if err := rows.Scan(&id, &product, &revenue, &dateSigned); err != nil {
			return nil, err
		}
		tmp := Contract{
			ID:         id,
			Product:    product,
			Revenue:    revenue,
			DateSigned: dateSigned,
		}

		contracts = append(contracts, tmp)
	}

	return &contracts, nil
}

//// TODO: asOf to date
func (g *Gateway) InsertRecognition(contractID int64, amount Money.Money, asOf string) (*int64, error) {

	tx, err := g.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	insertRecognitionStatement := `
INSERT INTO revenueRecognitions VALUES (?,?,?)
`

	result, err := g.DB.Exec(insertRecognitionStatement, contractID, amount, asOf)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}
