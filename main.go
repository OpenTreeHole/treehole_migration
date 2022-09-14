package main

import (
	"treehole_migration/config"
	"treehole_migration/models"
	"treehole_migration/tasks"
)

func main() {
	config.InitConfig()
	models.InitDB()
	tasks.StartTasks()
}
