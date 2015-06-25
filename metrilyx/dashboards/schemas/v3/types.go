package v3

type SchemaType string

const (
	SCHEMA_TYPE_DASHBOARD  SchemaType = "dashboard"
	SCHEMA_TYPE_POD        SchemaType = "pod"
	SCHEMA_TYPE_PANEL      SchemaType = "panel"
	SCHEMA_TYPE_DATASOURCE SchemaType = "datasource"
)

var SCHEMA_TYPES []SchemaType = []SchemaType{
	SCHEMA_TYPE_DASHBOARD,
	SCHEMA_TYPE_POD,
	SCHEMA_TYPE_PANEL,
	SCHEMA_TYPE_DATASOURCE,
}
