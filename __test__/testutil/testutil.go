package testutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ViniSantos88/payments-api/framework/config"
)

const (
	maxTest            = 150
	maxGroupTest       = 100
	characterTest      = "="
	characterGroupTest = "/"
)

//Setup sets up environment variables for unit testing run
func Setup() {

	os.Setenv("PSQL_DB_PORT", "5432")
	os.Setenv("PSQL_DB_NAME", "payments")

	config.Init()

}

// PrepareHandler prepare request Handler
func PrepareHandler(functions http.HandlerFunc) (recorder *httptest.ResponseRecorder, handlerFunc http.HandlerFunc) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(functions)
	return rr, handler
}

// ValidateError skips a unit test in case of error
func ValidateError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
		t.Skip()
	}
}

// PrintTestStart prints out the start with a Header for a unit test
func PrintTestStart(testName string) {
	firstLine := strings.Repeat(characterTest, maxTest)
	println(firstLine)

	secondLine := strings.Repeat(characterTest, 19) + "BEGIN " + testName + " "
	secondLine = secondLine + strings.Repeat(characterTest, maxTest-len(secondLine))
	println(secondLine)
}

// PrintTestEnd prints out the end with a Footer for a unit test
func PrintTestEnd(testName string) {
	firstLine := strings.Repeat(characterTest, 19) + "END " + testName + " "
	firstLine = firstLine + strings.Repeat(characterTest, maxTest-len(firstLine))
	println(firstLine)
	secondLine := strings.Repeat(characterTest, maxTest)
	println(secondLine)
}

// PrintTestGroupStart prints out the end with a Header for a test group
func PrintTestGroupStart(groupTestName string) {
	fixedLine := strings.Repeat(characterGroupTest, maxGroupTest)
	println(fixedLine)
	println(fixedLine)

	textLine := strings.Repeat(characterGroupTest, 5) + "BEGIN " + groupTestName + " "
	textLine = textLine + strings.Repeat(characterGroupTest, maxGroupTest-len(textLine))
	println(textLine)

	println(fixedLine)
	println(fixedLine)
}

// PrintTestGroupEnd prints out the end with a Footer for a test group
func PrintTestGroupEnd(groupTestName string) {
	fixedLine := strings.Repeat(characterGroupTest, maxGroupTest)
	println(fixedLine)
	println(fixedLine)

	textLine := strings.Repeat(characterGroupTest, 5) + "END " + groupTestName + " "
	textLine = textLine + strings.Repeat(characterGroupTest, maxGroupTest-len(textLine))
	println(textLine)

	println(fixedLine)
	println(fixedLine)
}
