package api

// enable secrets with http api,  under path :sys/mount/your secrtes
//https://www.vaultproject.io/api-docs/system/mounts#enable-secrets-engine
const (
	Address = "https://localhost:9200"
	Token   = "hvs.qYpjtjWtaTEu3dsYFZF7fFoP"
	//Todo Path to enable
	PathSecrets           = "sys/mounts/vaultkv" //to enbale
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

	// Todo Get data from path
	PathKvData = "vaultkv/data/user"
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
	valTransit = map[string]interface{}{
		"type": "transit",
	}
	ValKv = map[string]interface{}{
		"type": "kv-v2",
	}
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
