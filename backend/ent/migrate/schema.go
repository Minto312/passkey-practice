// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthHistoriesColumns holds the columns for the "auth_histories" table.
	AuthHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "method", Type: field.TypeString},
		{Name: "authenticated_at", Type: field.TypeTime},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "user_auth_histories", Type: field.TypeUUID},
	}
	// AuthHistoriesTable holds the schema information for the "auth_histories" table.
	AuthHistoriesTable = &schema.Table{
		Name:       "auth_histories",
		Columns:    AuthHistoriesColumns,
		PrimaryKey: []*schema.Column{AuthHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auth_histories_users_auth_histories",
				Columns:    []*schema.Column{AuthHistoriesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PasskeysColumns holds the columns for the "passkeys" table.
	PasskeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "credential_id", Type: field.TypeString, Unique: true},
		{Name: "public_key", Type: field.TypeBytes},
		{Name: "device_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_used_at", Type: field.TypeTime},
		{Name: "user_passkeys", Type: field.TypeUUID},
	}
	// PasskeysTable holds the schema information for the "passkeys" table.
	PasskeysTable = &schema.Table{
		Name:       "passkeys",
		Columns:    PasskeysColumns,
		PrimaryKey: []*schema.Column{PasskeysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passkeys_users_passkeys",
				Columns:    []*schema.Column{PasskeysColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "user_sessions", Type: field.TypeUUID},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthHistoriesTable,
		PasskeysTable,
		SessionsTable,
		UsersTable,
	}
)

func init() {
	AuthHistoriesTable.ForeignKeys[0].RefTable = UsersTable
	PasskeysTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
}
