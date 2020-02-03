package Gateway

import (
	"database/sql"
	"time"
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

//// TODO: asOf to date
//func (g *Gateway) FindRecognitionsFor(contractID int64, asOf string) ResultSet {
//	// findrecognitionstatemet渡してpreparedStatement作る
//	fmt.Println(contractID)
//	fmt.Println(asOf)
//	return ResultSet{}
//}

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
