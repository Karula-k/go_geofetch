package models

type EnvModel struct {
	JwtSecret   string `json:"jwt_secret"`
	Port        string `json:"port"`
	PostgresUser string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDB string `json:"postgres_db"`
	DatabaseUrl string `json:"database_url"`
}
