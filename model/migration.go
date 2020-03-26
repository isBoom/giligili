package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Video{}).
		AutoMigrate(&Comment{})
	DB.Model(&Comment{}).AddForeignKey("video_id", "videos(id)", "no action", "no action")
	//DB.Model(&Comment{}).AddForeignKey("parent_user_id", "users(id)", "no action", "no action")
	DB.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "no action", "no action")
}
