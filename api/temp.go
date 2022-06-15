package api

const (
	Address               = "http://127.0.0.1:8200"
	Token                 = "hvs.iUeLC47BGapeXFXSw98yEP1y"
	PathSecrets           = "secret/data/mydata"
	PathPolicy            = "sys/policies/acl/apps"
	PathCreateToken       = "/auth/token/create"
	PathReCreateToken     = "/database/rotate-role/education"
	PathRotateTokenByRoot = "/database/static-creds/education"
	// path to enable engies
	// patData := "/sys/mounts/database"

	// path to write config engies database
	// ppathConfigDb := "/database/config/postgresql"

	// path to setup rotate password

	PathRoot   = "/database/rotate-root/postgresql"
	PathStatic = "/database/static-roles/education"
)

var (
	// Todo: here is config database
	Value = map[string]interface{}{
		"plugin_name":    "postgresql-database-plugin",
		"allowed_roles":  "*",
		"connection_url": "postgresql://{{username}}:{{password}}@localhost:5432/postgres?sslmode=disable",
		"username":       "root",
		"password":       "rootpassword",
	}
	// ToDo : here is enable secrets engies
	// valueDB := map[string]interface{}{
	// 	"type": "database",
	// }
	// Todo: here is config static role
	vale       = "[ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';],"
	StaticRole = map[string]interface{}{
		"db_name":             "postgresql",
		"rotation_statements": "ALTER USER tayto WITH PASSWORD '{{password}}'",
		"username":            "tayto",
		"rotation_period":     "86400",
	}
	ValPolicy = map[string]interface{}{
		"policy": "path \"database/static-creds/education\" {capabilities = [ \"read\" ]}",
	}
	ValToken = map[string]interface{}{
		"policies": "apps",
	}
)
