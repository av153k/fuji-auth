package constants

const (
	// For database related ENV values
	DbHost     string = "DB_HOST"
	DbType     string = "DB_TYPE"
	DbName     string = "DB_NAME"
	DbUser     string = "DB_USER"
	DbPassword string = "DB_PASSWORD"

	//For JWT related ENV values
	JwtSecret                 string = "JWT_SECRET"
	JwtValidityInMinutes      string = "JWT_VALIDITY_IN_MINUTES"
	JwtRefreshKey             string = "JWT_REFRESH_KEY"
	JwtRefreshValidityInHours string = "JWT_REFRESH_VALIDITY_IN_HOURS"
)
