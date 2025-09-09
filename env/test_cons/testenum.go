package testcons

// test_enum.go - содержит константы описывающие набор значений опции для тестирования обработки перечислимого типа
// и перечислимый тип значениями которого могут быть эти константы

//go:generate stringer -type=TestEnum
//go:generate enummethods -type=TestEnum
//type level string

// TestEnum - тип необходимый для сопоставления набора строковых значений их индексам
type TestEnum int

const (
	string1 = TestEnum(iota)
	string2
)

const (
	STRING1 = string1
	STRING2 = string2
)
