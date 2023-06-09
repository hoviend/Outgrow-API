// Code generated by ent, DO NOT EDIT.

package organizationaccountcategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the organizationaccountcategory type in the database.
	Label = "organization_account_category"
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
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// Table holds the table name of the organizationaccountcategory in the database.
	Table = "organization_account_categories"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "organization_accounts"
	// AccountsInverseTable is the table name for the OrganizationAccount entity.
	// It exists in this package in order to avoid circular dependency with the "organizationaccount" package.
	AccountsInverseTable = "organization_accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "category_id"
	// TypeTable is the table that holds the type relation/edge.
	TypeTable = "organization_account_categories"
	// TypeInverseTable is the table name for the OrganizationAccountType entity.
	// It exists in this package in order to avoid circular dependency with the "organizationaccounttype" package.
	TypeInverseTable = "organization_account_types"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "account_type_id"
)

// Columns holds all SQL columns for organizationaccountcategory fields.
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

// Order defines the ordering method for the OrganizationAccountCategory queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountTypeID orders the results by the account_type_id field.
func ByAccountTypeID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAccountTypeID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAccountsCount orders the results by accounts count.
func ByAccountsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountsStep(), opts...)
	}
}

// ByAccounts orders the results by accounts terms.
func ByAccounts(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTypeField orders the results by type field.
func ByTypeField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newAccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
	)
}
func newTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TypeTable, TypeColumn),
	)
}
