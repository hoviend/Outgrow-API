package enum

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "DEBIT"
	TransactionTypeCredit TransactionType = "CREDIT"
)

func (TransactionType) Values() (kinds []string) {
	for _, s := range []TransactionType{TransactionTypeDebit, TransactionTypeCredit} {
		kinds = append(kinds, string(s))
	}
	return
}
