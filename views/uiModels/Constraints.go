package uiModels

type Entity interface {
	GetFields() []string
	GetFieldsValues() []string
	GetInstanceName() string
}
