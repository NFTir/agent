/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: loadEnvVars.go provide function that helps load environemnt variables
*/

// @package
package utils

// @import
import (
	"github.com/joho/godotenv"
)

// @dev Loads environment varables
func LoadEnvVars()  {
	err := godotenv.Load()
	HandleException(err);
}