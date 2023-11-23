package variables

var FoldersName = []string{
	"../src",
	"../migrations",
	"../.github",
	"../.github/workflows",
	"../controllers",
	"../services",
	"../repositories",
	"../domain",
	"../config",
}

// var FilesName = []string{
// 	"../src/main.go",
// 	"../.github/workflows/deploy.yml",
// 	"../controllers/controller.go",
// 	"../services/service.go",
// 	"../repositories/repository.go",
// 	"../domain/contracts.go",
// 	"../config/db.go",
// }

var ContentFiles = []map[string]string{
	{
		"fileName": "../src/main.go",
		"content":  `package main`,
	},
	{
		"fileName": "../.github/workflows/deploy.yml",
		"content":  ``,
	},
	{
		"fileName": "../controllers/controller.go",
		"content":  `package controllers`,
	},
	{
		"fileName": "../services/service.go",
		"content":  `package services`,
	},
	{
		"fileName": "../repositories/repository.go",
		"content":  `package repositories`,
	},
	{
		"fileName": "../domain/contracts.go",
		"content":  `package domain`,
	},
	{
		"fileName": "../config/db.go",
		"content":  `package config`,
	},
}
