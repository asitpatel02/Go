package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// TokenType represents the type of token
type TokenType int

const (
	Integer     TokenType = iota // Token type for integers
	At                           // Token type for '@' operator
	Caret                        // Token type for '^' operator
	Exclamation                  // Token type for '!' operator
	Invalid                      // Token type for invalid tokens
)

// Token represents a token in the input
type Token struct {
	Type  TokenType // Type of the token
	Value string    // Value of the token
}

// Scanner tokenizes the input string and identifies tokens
func Scanner(input string) []Token {
	var tokens []Token
	var currentToken strings.Builder

	for _, char := range input {
		switch {
		case unicode.IsDigit(char): // If the character is a digit
			currentToken.WriteRune(char) // Append the digit to the current token
		case char == '@': // If the character is '@' operator
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
			tokens = append(tokens, Token{Type: At, Value: "@"}) // Add the '@' operator token
		case char == '^': // If the character is '^' operator
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
			tokens = append(tokens, Token{Type: Caret, Value: "^"}) // Add the '^' operator token
		case char == '!': // If the character is '!' operator
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
			tokens = append(tokens, Token{Type: Exclamation, Value: "!"}) // Add the '!' operator token
		case char == ' ': // If the character is a whitespace
			// Ignore whitespace
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
		default:
			// If the character is not a valid lexeme, mark the token as invalid
			tokens = append(tokens, Token{Type: Invalid, Value: string(char)})
			return tokens
		}
	}

	if currentToken.Len() > 0 { // If there is a current token
		tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
	}

	return tokens
}

// Interpreter evaluates the expression represented by tokens
func Interpreter(tokens []Token) (int, bool) {
	var result int
	var currentNumber int
	var operator string

	for _, token := range tokens {
		switch token.Type {
		case Integer:
			num, err := strconv.Atoi(token.Value)
			if err != nil {
				return 0, false
			}
			switch operator {
			case "":
				currentNumber = num
			case "^":
				currentNumber *= num
			case "@":
				result += currentNumber
				currentNumber = num
			}
		case Exclamation:
			currentNumber *= -1
		default:
			operator = token.Value
		}
	}

	// Add the last number in case there's any remaining
	switch operator {
	case "@":
		result += currentNumber
	case "^":
		result *= currentNumber
	}

	return result, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read input from standard input
	for scanner.Scan() {                  // Read input line by line
		input := scanner.Text()  // Get the input line
		tokens := Scanner(input) // Tokenize the input line
		validTokens := true
		for _, token := range tokens { // Check if there are any invalid tokens
			if token.Type == Invalid {
				validTokens = false
				break
			}
		}
		if !validTokens {
			fmt.Printf("\"%s\" contains invalid lexemes and thus is not an expression.\n", input) // Print error message for invalid tokens
		} else {
			// Check if consecutive operators are present
			consecutiveOperators := false
			for i := 1; i < len(tokens); i++ {
				if (tokens[i-1].Type == At || tokens[i-1].Type == Caret || tokens[i-1].Type == Exclamation) &&
					(tokens[i].Type == At || tokens[i].Type == Caret || tokens[i].Type == Exclamation) {
					consecutiveOperators = true
					break
				}
			}
			if consecutiveOperators {
				fmt.Printf("\"%s\" is not an expression\n", input) // Print message if consecutive operators are present
			} else {
				result, valid := Interpreter(tokens)
				if valid {
					fmt.Println(result) // Print the result of the expression evaluation
				} else {
					fmt.Printf("\"%s\" is not an expression\n", input) // Print message if the input is not a valid expression
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err) // Print error message for any input reading errors
	}
}
