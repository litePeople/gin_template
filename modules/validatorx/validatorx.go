package validatorx

func Struct(s interface{}) error {
	return cli.vali.Struct(s)
}

func Var(filed interface{}, tag string) error {
	return cli.vali.Var(filed, tag)
}

func Map(data, rules map[string]interface{}) map[string]interface{} {
	return cli.vali.ValidateMap(data, rules)
}
