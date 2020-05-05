package database

import "os"

func mongoConnectionString() string {
	return os.Getenv("MONGO")
}
func mongoServer() string {
	return os.Getenv("MONGO_SERVER")
}
func mongoDatabase() string {
	return os.Getenv("MONGO_DATABASE")
}
func mongoUsername() string {
	return os.Getenv("MONGO_USERNAME")
}
func mongoPassword() string {
	return os.Getenv("MONGO_PASSWORD")
}
func mongoAuthSource() string {
	return os.Getenv("MONGO_AUTH_SOURCE")
}
func mongoReplicaSet() string {
	return os.Getenv("MONGO_REPLICA_SET")
}