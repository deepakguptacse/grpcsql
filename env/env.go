package env

import (
	"os"
	"strings"
)

type Environment int

const (
	Dev Environment = iota
	Prod
)

var currentEnv Environment

func init() {
	env := os.Getenv("ENV")
	switch strings.ToLower(env) {
	case "dev":
		currentEnv = Dev
	case "prod":
		currentEnv = Prod
	default:
		currentEnv = Dev // Default to Dev if no environment variable is set
	}
}

func GetCurrentEnv() Environment {
	return currentEnv
}

func (e Environment) String() string {
	return [...]string{"Dev", "Prod"}[e]
}
