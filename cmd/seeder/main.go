package main

import (
	"flag"

	"github.com/WeTrustPlatform/charity-management-serv/db"
	"github.com/WeTrustPlatform/charity-management-serv/seed"
)

func main() {
	var dryRun bool
	flag.BoolVar(&dryRun, "dryrun", false, "dryrun")

	var filename string
	flag.StringVar(&filename, "data", "../seed/data.txt", "name of the data file inside <root>/seed/ folder")

	flag.Parse()

	dbInstance := db.Connect(false)
	defer dbInstance.Close()
	dbInstance.AutoMigrate(&db.Charity{})

	seed.Populate(dbInstance, filename, dryRun)
	db.UpdateSearchIndex(dbInstance)
}
