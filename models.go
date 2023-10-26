type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}

type Video struct {
    gorm.Model
    UserID uint
    Title  string
    URL    string
}
