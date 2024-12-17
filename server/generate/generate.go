package main

import (
	"server/models"

	"gorm.io/gen"
)

// generate code
func main() {
	models.InitDB()
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../models/dao",    //curd代码的输出路径
		ModelPkgPath:      "../models/entity", //model代码的输出路径
		Mode:              gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})
	g.UseDB(models.Db)
	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)
	// execute the action of code generation
	g.Execute()
}
