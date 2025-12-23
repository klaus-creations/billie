package token;

type TokenType string;

type Token struct {
    Type  TokenType
    Literal string
}



/// definig our token - names based on our token type
const (

    // the token which we can't identify
    ILLEGAL = "ILLEGAL"
    // which tells end of the file
    EOF = "EOF"

    IDENT = "IDEN"
    INT = "INT"


    ASSIGN = "="
    PLUS = "+"
	MINUS = "-"
	MULTIPLICATION = "*"
	DIVIDE = "/"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET = "LET"
)




