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
	Integer  TokenType = iota // Token type for integers
	Plus                      // Token type for plus operator
	Multiply                  // Token type for multiply operator
	Minus                     // Token type for minus operator
	Invalid                   // Token type for invalid tokens
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
		case char == '+': // If the character is a plus operator
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
			tokens = append(tokens, Token{Type: Plus, Value: "+"}) // Add the plus operator token
		case char == '*': // If the character is a multiply operator
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
			tokens = append(tokens, Token{Type: Multiply, Value: "*"}) // Add the multiply operator token
		case char == '-': // If the character is a minus operator
			if currentToken.Len() == 0 { // If there is no current token
				currentToken.WriteRune(char) // Append the minus operator to the current token
			} else {
				if _, err := strconv.Atoi(currentToken.String()); err == nil { // If the current token is a valid integer
					tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
					currentToken.Reset()                                                        // Reset the current token
					currentToken.WriteRune(char)                                                // Append the minus operator to the current token
				} else {
					currentToken.WriteRune(char) // Append the minus operator to the current token
				}
			}
		case char == ' ': // If the character is a whitespace
			// Ignore whitespace
			if currentToken.Len() > 0 { // If there is a current token
				tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
				currentToken.Reset()                                                        // Reset the current token
			}
		default:
			currentToken.WriteRune(char) // Append the character to the current token
		}
	}

	if currentToken.Len() > 0 { // If there is a current token
		tokens = append(tokens, Token{Type: Integer, Value: currentToken.String()}) // Add the current token as an integer token
	}

	return tokens
}

// RecursiveDescentParser parses the input and determines if it's a valid expression
func RecursiveDescentParser(tokens []Token) bool {
	index := 0
	if exp(tokens, &index) && index == len(tokens) { // If the tokens represent a valid expression and all tokens have been consumed
		return true
	}
	return false
}

// exp checks if the tokens represent an expression
func exp(tokens []Token, index *int) bool {
	if term(tokens, index) { // If the tokens represent a term
		if *index < len(tokens) && (tokens[*index].Type == Plus || tokens[*index].Type == Minus) { // If there is a plus or minus operator token
			(*index)++                // Consume the operator token
			return exp(tokens, index) // Recursively check for more terms
		}
		return true
	}
	return false
}

// term checks if the tokens represent a term
func term(tokens []Token, index *int) bool {
	if number(tokens, index) { // If the tokens represent a number
		if *index < len(tokens) && tokens[*index].Type == Multiply { // If there is a multiply operator token
			(*index)++                 // Consume the operator token
			return term(tokens, index) // Recursively check for more factors
		}
		return true
	}
	return false
}

// number checks if the tokens represent a number
func number(tokens []Token, index *int) bool {
	if *index < len(tokens) && tokens[*index].Type == Integer { // If the token is an integer
		(*index)++ // Consume the integer token
		return true
	}
	if *index < len(tokens) && tokens[*index].Type == Minus { // If the token is a minus operator
		(*index)++                                                  // Consume the minus operator token
		if *index < len(tokens) && tokens[*index].Type == Integer { // If there is a following integer token
			(*index)++ // Consume the integer token
			return true
		}
	}
	return false
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
			fmt.Printf("%s contains tokens which are not valid\n", input) // Print error message for invalid tokens
		} else if RecursiveDescentParser(tokens) {
			fmt.Printf("%s is an expression\n", input) // Print message if the input is a valid expression
		} else {
			fmt.Printf("%s is not an expression\n", input) // Print message if the input is not a valid expression
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err) // Print error message for any input reading errors
	}
}
