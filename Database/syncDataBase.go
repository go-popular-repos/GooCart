package Database

import "github.com/mohdjishin/fiberRESTApi/model"

func SyncDatabase() {
	db := OpenDb()
	defer CloseDb(db)
	db.AutoMigrate(&model.Admin{})

	db.AutoMigrate(&model.Products{})
	db.AutoMigrate(&model.ProductImage{})

	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Address{})
}
