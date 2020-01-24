package Gateway

import (
	"database/sql"
	"fmt"
)
type Gateway struct {
	DB sql.DB
}

type ResultSet struct {

}
// TODO: asOf to date
func (g *Gateway) FindRecognitionsFor(contractID int64, asOf string) ResultSet {
	// findrecognitionstatemet渡してpreparedStatement作る
	fmt.Println(contractID)
	fmt.Println(asOf)
	return ResultSet{}
}

func (g *Gateway) FindContract(contractID int64) ResultSet {
	findContractStatement :=
```
SELECT * FROM
contracts c,
product p
WHERE ID = ?
AND c.product = p.ID;
```
	// findrecognitionstatemet渡してpreparedStatement作る
	fmt.Println(contractID)
	return ResultSet{}
}

