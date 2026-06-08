// Package config loads gorag configuration from the execution environment.
package config

import "os"

// Config contains external service configuration.
type Config struct {
	Neo4j Neo4jConfig
}

// Neo4jConfig contains Neo4j connection settings.
type Neo4jConfig struct {
	URI      string
	User     string
	Password string
	Database string
}

// FromEnv loads configuration from environment variables.
func FromEnv() Config {
	return Config{
		Neo4j: Neo4jConfig{
			URI:      getEnv("GORAG_NEO4J_URI", "neo4j://localhost:7687"),
			User:     getEnv("GORAG_NEO4J_USER", "neo4j"),
			Password: getEnv("GORAG_NEO4J_PASSWORD", "gorag"),
			Database: getEnv("GORAG_NEO4J_DATABASE", "neo4j"),
		},
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
