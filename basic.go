package main

const (
	ProgramName  = "алг "
	ProgramStart = "нач"
	ProgramEnd   = "кон"
)

var ModuleNames = map[string]string{"robot": "Робот"}

func ConstructAlgorithm(name string) string {
	return ProgramName + name
}

func ConstructUse(module string) string {
	res := ""
	ok := false
	res, ok = ModuleNames[module]
	if !ok {
		PrintWarning(`Module name "%s" could not be resolved!`, module)
		res = module
	}
	return "использовать " + res
}
