package yugigo

/*
GetDataBaseInfo returns information regarding the latest database
*/
func GetDataBaseInfo() (Database, error) {
	var database []Database
	err := getData(databaseurl, &database)
	if err != nil {
		return Database{}, err
	}
	return database[0], nil
}

//Database Struct contains the version that the package is currently using, and the Last Update Time of the Database.
type Database struct {
	DatabaseVersion string `json:"database_version"`
	LastUpdate      string `json:"last_update"`
}
