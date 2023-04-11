// Code generated by ent, DO NOT EDIT.

package masteraccountcategory

import (
	"time"
)

const (
	// Label holds the string label denoting the masteraccountcategory type in the database.
	Label = "master_account_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountTypeID holds the string denoting the account_type_id field in the database.
	FieldAccountTypeID = "account_type_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// EdgeAccType holds the string denoting the acctype edge name in mutations.
	EdgeAccType = "accType"
	// Table holds the table name of the masteraccountcategory in the database.
	Table = "master_account_categories"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "master_accounts"
	// AccountsInverseTable is the table name for the MasterAccount entity.
	// It exists in this package in order to avoid circular dependency with the "masteraccount" package.
	AccountsInverseTable = "master_accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "category_id"
	// AccTypeTable is the table that holds the accType relation/edge.
	AccTypeTable = "master_account_categories"
	// AccTypeInverseTable is the table name for the MasterAccountType entity.
	// It exists in this package in order to avoid circular dependency with the "masteraccounttype" package.
	AccTypeInverseTable = "master_account_types"
	// AccTypeColumn is the table column denoting the accType relation/edge.
	AccTypeColumn = "account_type_id"
)

// Columns holds all SQL columns for masteraccountcategory fields.
var Columns = []string{
	FieldID,
	FieldAccountTypeID,
	FieldName,
	FieldDescription,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)