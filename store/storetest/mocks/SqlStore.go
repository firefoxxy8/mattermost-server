// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import gorp "github.com/mattermost/gorp"
import mock "github.com/stretchr/testify/mock"

import store "github.com/demisto/mattermost-server/store"

// SqlStore is an autogenerated mock type for the SqlStore type
type SqlStore struct {
	mock.Mock
}

// AlterColumnTypeIfExists provides a mock function with given fields: tableName, columnName, mySqlColType, postgresColType
func (_m *SqlStore) AlterColumnTypeIfExists(tableName string, columnName string, mySqlColType string, postgresColType string) bool {
	ret := _m.Called(tableName, columnName, mySqlColType, postgresColType)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string) bool); ok {
		r0 = rf(tableName, columnName, mySqlColType, postgresColType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Audit provides a mock function with given fields:
func (_m *SqlStore) Audit() store.AuditStore {
	ret := _m.Called()

	var r0 store.AuditStore
	if rf, ok := ret.Get(0).(func() store.AuditStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.AuditStore)
		}
	}

	return r0
}

// Channel provides a mock function with given fields:
func (_m *SqlStore) Channel() store.ChannelStore {
	ret := _m.Called()

	var r0 store.ChannelStore
	if rf, ok := ret.Get(0).(func() store.ChannelStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ChannelStore)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *SqlStore) Close() {
	_m.Called()
}

// ClusterDiscovery provides a mock function with given fields:
func (_m *SqlStore) ClusterDiscovery() store.ClusterDiscoveryStore {
	ret := _m.Called()

	var r0 store.ClusterDiscoveryStore
	if rf, ok := ret.Get(0).(func() store.ClusterDiscoveryStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ClusterDiscoveryStore)
		}
	}

	return r0
}

// Command provides a mock function with given fields:
func (_m *SqlStore) Command() store.CommandStore {
	ret := _m.Called()

	var r0 store.CommandStore
	if rf, ok := ret.Get(0).(func() store.CommandStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.CommandStore)
		}
	}

	return r0
}

// CommandWebhook provides a mock function with given fields:
func (_m *SqlStore) CommandWebhook() store.CommandWebhookStore {
	ret := _m.Called()

	var r0 store.CommandWebhookStore
	if rf, ok := ret.Get(0).(func() store.CommandWebhookStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.CommandWebhookStore)
		}
	}

	return r0
}

// Compliance provides a mock function with given fields:
func (_m *SqlStore) Compliance() store.ComplianceStore {
	ret := _m.Called()

	var r0 store.ComplianceStore
	if rf, ok := ret.Get(0).(func() store.ComplianceStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ComplianceStore)
		}
	}

	return r0
}

// CreateColumnIfNotExists provides a mock function with given fields: tableName, columnName, mySqlColType, postgresColType, defaultValue
func (_m *SqlStore) CreateColumnIfNotExists(tableName string, columnName string, mySqlColType string, postgresColType string, defaultValue string) bool {
	ret := _m.Called(tableName, columnName, mySqlColType, postgresColType, defaultValue)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) bool); ok {
		r0 = rf(tableName, columnName, mySqlColType, postgresColType, defaultValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateCompositeIndexIfNotExists provides a mock function with given fields: indexName, tableName, columnNames
func (_m *SqlStore) CreateCompositeIndexIfNotExists(indexName string, tableName string, columnNames []string) bool {
	ret := _m.Called(indexName, tableName, columnNames)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, []string) bool); ok {
		r0 = rf(indexName, tableName, columnNames)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateFullTextIndexIfNotExists provides a mock function with given fields: indexName, tableName, columnName
func (_m *SqlStore) CreateFullTextIndexIfNotExists(indexName string, tableName string, columnName string) bool {
	ret := _m.Called(indexName, tableName, columnName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(indexName, tableName, columnName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateIndexIfNotExists provides a mock function with given fields: indexName, tableName, columnName
func (_m *SqlStore) CreateIndexIfNotExists(indexName string, tableName string, columnName string) bool {
	ret := _m.Called(indexName, tableName, columnName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(indexName, tableName, columnName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateUniqueIndexIfNotExists provides a mock function with given fields: indexName, tableName, columnName
func (_m *SqlStore) CreateUniqueIndexIfNotExists(indexName string, tableName string, columnName string) bool {
	ret := _m.Called(indexName, tableName, columnName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(indexName, tableName, columnName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DoesColumnExist provides a mock function with given fields: tableName, columName
func (_m *SqlStore) DoesColumnExist(tableName string, columName string) bool {
	ret := _m.Called(tableName, columName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(tableName, columName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DoesTableExist provides a mock function with given fields: tablename
func (_m *SqlStore) DoesTableExist(tablename string) bool {
	ret := _m.Called(tablename)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(tablename)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DriverName provides a mock function with given fields:
func (_m *SqlStore) DriverName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Emoji provides a mock function with given fields:
func (_m *SqlStore) Emoji() store.EmojiStore {
	ret := _m.Called()

	var r0 store.EmojiStore
	if rf, ok := ret.Get(0).(func() store.EmojiStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.EmojiStore)
		}
	}

	return r0
}

// FileInfo provides a mock function with given fields:
func (_m *SqlStore) FileInfo() store.FileInfoStore {
	ret := _m.Called()

	var r0 store.FileInfoStore
	if rf, ok := ret.Get(0).(func() store.FileInfoStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.FileInfoStore)
		}
	}

	return r0
}

// GetAllConns provides a mock function with given fields:
func (_m *SqlStore) GetAllConns() []*gorp.DbMap {
	ret := _m.Called()

	var r0 []*gorp.DbMap
	if rf, ok := ret.Get(0).(func() []*gorp.DbMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gorp.DbMap)
		}
	}

	return r0
}

// GetCurrentSchemaVersion provides a mock function with given fields:
func (_m *SqlStore) GetCurrentSchemaVersion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMaster provides a mock function with given fields:
func (_m *SqlStore) GetMaster() *gorp.DbMap {
	ret := _m.Called()

	var r0 *gorp.DbMap
	if rf, ok := ret.Get(0).(func() *gorp.DbMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorp.DbMap)
		}
	}

	return r0
}

// GetMaxLengthOfColumnIfExists provides a mock function with given fields: tableName, columnName
func (_m *SqlStore) GetMaxLengthOfColumnIfExists(tableName string, columnName string) string {
	ret := _m.Called(tableName, columnName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(tableName, columnName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetReplica provides a mock function with given fields:
func (_m *SqlStore) GetReplica() *gorp.DbMap {
	ret := _m.Called()

	var r0 *gorp.DbMap
	if rf, ok := ret.Get(0).(func() *gorp.DbMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorp.DbMap)
		}
	}

	return r0
}

// GetSearchReplica provides a mock function with given fields:
func (_m *SqlStore) GetSearchReplica() *gorp.DbMap {
	ret := _m.Called()

	var r0 *gorp.DbMap
	if rf, ok := ret.Get(0).(func() *gorp.DbMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorp.DbMap)
		}
	}

	return r0
}

// Job provides a mock function with given fields:
func (_m *SqlStore) Job() store.JobStore {
	ret := _m.Called()

	var r0 store.JobStore
	if rf, ok := ret.Get(0).(func() store.JobStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.JobStore)
		}
	}

	return r0
}

// License provides a mock function with given fields:
func (_m *SqlStore) License() store.LicenseStore {
	ret := _m.Called()

	var r0 store.LicenseStore
	if rf, ok := ret.Get(0).(func() store.LicenseStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LicenseStore)
		}
	}

	return r0
}

// MarkSystemRanUnitTests provides a mock function with given fields:
func (_m *SqlStore) MarkSystemRanUnitTests() {
	_m.Called()
}

// OAuth provides a mock function with given fields:
func (_m *SqlStore) OAuth() store.OAuthStore {
	ret := _m.Called()

	var r0 store.OAuthStore
	if rf, ok := ret.Get(0).(func() store.OAuthStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.OAuthStore)
		}
	}

	return r0
}

// Plugin provides a mock function with given fields:
func (_m *SqlStore) Plugin() store.PluginStore {
	ret := _m.Called()

	var r0 store.PluginStore
	if rf, ok := ret.Get(0).(func() store.PluginStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PluginStore)
		}
	}

	return r0
}

// Post provides a mock function with given fields:
func (_m *SqlStore) Post() store.PostStore {
	ret := _m.Called()

	var r0 store.PostStore
	if rf, ok := ret.Get(0).(func() store.PostStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PostStore)
		}
	}

	return r0
}

// Preference provides a mock function with given fields:
func (_m *SqlStore) Preference() store.PreferenceStore {
	ret := _m.Called()

	var r0 store.PreferenceStore
	if rf, ok := ret.Get(0).(func() store.PreferenceStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PreferenceStore)
		}
	}

	return r0
}

// Reaction provides a mock function with given fields:
func (_m *SqlStore) Reaction() store.ReactionStore {
	ret := _m.Called()

	var r0 store.ReactionStore
	if rf, ok := ret.Get(0).(func() store.ReactionStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ReactionStore)
		}
	}

	return r0
}

// RemoveColumnIfExists provides a mock function with given fields: tableName, columnName
func (_m *SqlStore) RemoveColumnIfExists(tableName string, columnName string) bool {
	ret := _m.Called(tableName, columnName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(tableName, columnName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveIndexIfExists provides a mock function with given fields: indexName, tableName
func (_m *SqlStore) RemoveIndexIfExists(indexName string, tableName string) bool {
	ret := _m.Called(indexName, tableName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(indexName, tableName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveTableIfExists provides a mock function with given fields: tableName
func (_m *SqlStore) RemoveTableIfExists(tableName string) bool {
	ret := _m.Called(tableName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(tableName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RenameColumnIfExists provides a mock function with given fields: tableName, oldColumnName, newColumnName, colType
func (_m *SqlStore) RenameColumnIfExists(tableName string, oldColumnName string, newColumnName string, colType string) bool {
	ret := _m.Called(tableName, oldColumnName, newColumnName, colType)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string) bool); ok {
		r0 = rf(tableName, oldColumnName, newColumnName, colType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Role provides a mock function with given fields:
func (_m *SqlStore) Role() store.RoleStore {
	ret := _m.Called()

	var r0 store.RoleStore
	if rf, ok := ret.Get(0).(func() store.RoleStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.RoleStore)
		}
	}

	return r0
}

// Session provides a mock function with given fields:
func (_m *SqlStore) Session() store.SessionStore {
	ret := _m.Called()

	var r0 store.SessionStore
	if rf, ok := ret.Get(0).(func() store.SessionStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.SessionStore)
		}
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *SqlStore) Status() store.StatusStore {
	ret := _m.Called()

	var r0 store.StatusStore
	if rf, ok := ret.Get(0).(func() store.StatusStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StatusStore)
		}
	}

	return r0
}

// System provides a mock function with given fields:
func (_m *SqlStore) System() store.SystemStore {
	ret := _m.Called()

	var r0 store.SystemStore
	if rf, ok := ret.Get(0).(func() store.SystemStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.SystemStore)
		}
	}

	return r0
}

// Team provides a mock function with given fields:
func (_m *SqlStore) Team() store.TeamStore {
	ret := _m.Called()

	var r0 store.TeamStore
	if rf, ok := ret.Get(0).(func() store.TeamStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.TeamStore)
		}
	}

	return r0
}

// Token provides a mock function with given fields:
func (_m *SqlStore) Token() store.TokenStore {
	ret := _m.Called()

	var r0 store.TokenStore
	if rf, ok := ret.Get(0).(func() store.TokenStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.TokenStore)
		}
	}

	return r0
}

// TotalMasterDbConnections provides a mock function with given fields:
func (_m *SqlStore) TotalMasterDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// TotalReadDbConnections provides a mock function with given fields:
func (_m *SqlStore) TotalReadDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// TotalSearchDbConnections provides a mock function with given fields:
func (_m *SqlStore) TotalSearchDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// User provides a mock function with given fields:
func (_m *SqlStore) User() store.UserStore {
	ret := _m.Called()

	var r0 store.UserStore
	if rf, ok := ret.Get(0).(func() store.UserStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.UserStore)
		}
	}

	return r0
}

// UserAccessToken provides a mock function with given fields:
func (_m *SqlStore) UserAccessToken() store.UserAccessTokenStore {
	ret := _m.Called()

	var r0 store.UserAccessTokenStore
	if rf, ok := ret.Get(0).(func() store.UserAccessTokenStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.UserAccessTokenStore)
		}
	}

	return r0
}

// Webhook provides a mock function with given fields:
func (_m *SqlStore) Webhook() store.WebhookStore {
	ret := _m.Called()

	var r0 store.WebhookStore
	if rf, ok := ret.Get(0).(func() store.WebhookStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.WebhookStore)
		}
	}

	return r0
}
