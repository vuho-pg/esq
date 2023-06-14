package main

import . "github.com/vuho-pg/esq/codegen/builder"

var queryInterface = Interface("Query").Func("IsQuery")

func interfaceBuild() error {
	return Run(queryInterface)
}
