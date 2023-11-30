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
	{
		"fileName": "../config/environment.go",
		"content":  `package config`,
	},
	{
		"fileName": "../.env",
		"content":  ``,
	},
	{
		"fileName": "../Dockerfile",
		"content":  ``,
	},
	{
		"fileName": "../docker-compose.yml",
		"content":  ``,
	},
	{
		"fileName": "makefile",
		"content":  ``,
	},
}
