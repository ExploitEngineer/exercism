package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	uppercaseCustomerName := strings.ToUpper(customer)
	return "Welcome to the Tech Palace," + " " + uppercaseCustomerName
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	fancyBorder := strings.Repeat("*", numStarsPerLine)
	finalFancyOutput := fancyBorder + "\n" + welcomeMsg + "\n" + fancyBorder
	return finalFancyOutput
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutStarsString := strings.ReplaceAll(oldMsg, "*", "")
	cleanMessage := strings.TrimSpace(withoutStarsString)
	return cleanMessage
}
