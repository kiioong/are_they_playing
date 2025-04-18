package Database

import (
	"log"
	"os"
	"time"

	hash "github.com/kiioong/are_they_playing/internal/Hash"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{}, &League{}, &Team{}, &Sport{}, &UserTeam{}, &Game{})

	password, err := hash.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return
	}

	DB.Create(&User{Username: "admin", Password: password})
}

type User struct {
	ID        uint64    `gorm:"primaryKey"`
	Username  string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Teams []Team `gorm:"many2many:user_teams;"`
}

type League struct {
	ID      uint32 `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(255);not null"`
	SportID uint   `gorm:"not null"`

	Sport Sport
	Teams []Team `gorm:"many2many:league_teams;"`
}

type Team struct {
	ID         uint32 `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255);not null,index:idx_team,unique"`
	Gender     string `gorm:"type:varchar(6);not null,index:idx_team,unique"`
	PathToLogo string `gorm:"type:varchar(255)"`

	Leagues []League `gorm:"many2many:league_teams;"`
}

type Sport struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);not null"`

	Leagues []League `gorm:"foreignKey:SportID"`
}

type UserTeam struct {
	UserID uint   `gorm:"primaryKey"`
	TeamID uint32 `gorm:"primaryKey"`

	User User `gorm:"foreignKey:UserID"`
	Team Team `gorm:"foreignKey:TeamID"`
}

type LeagueTeam struct {
	LeagueID uint `gorm:"primaryKey"`
	TeamID   uint `gorm:"primaryKey"`

	User User `gorm:"foreignKey:LeagueID"`
	Team Team `gorm:"foreignKey:TeamID"`
}

type Game struct {
	ID         uint `gorm:"primaryKey"`
	HomeTeamID uint `gorm:"index:idx_game,unique"`
	AwayTeamID uint `gorm:"index:idx_game,unique"`
	LeagueID   uint `gorm:"index:idx_game,unique"`

	StartTime time.Time

	HomeTeam Team   `gorm:"foreignKey:HomeTeamID"`
	AwayTeam Team   `gorm:"foreignKey:AwayTeamID"`
	League   League `gorm:"foreignKey:LeagueID"`
}
