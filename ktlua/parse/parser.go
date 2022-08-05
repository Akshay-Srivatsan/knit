// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parse

import __yyfmt__ "fmt"

//line parser.go.y:2

import (
	"github.com/zyedidia/knit/ktlua/ast"
)

//line parser.go.y:34
type yySymType struct {
	yys   int
	token ast.Token

	stmts []ast.Stmt
	stmt  ast.Stmt

	funcname *ast.FuncName
	funcexpr *ast.FunctionExpr

	exprlist []ast.Expr
	expr     ast.Expr

	fieldlist []*ast.Field
	field     *ast.Field
	fieldsep  string

	namelist []string
	parlist  *ast.ParList
}

const TAnd = 57346
const TBreak = 57347
const TDo = 57348
const TElse = 57349
const TElseIf = 57350
const TEnd = 57351
const TFalse = 57352
const TFor = 57353
const TFunction = 57354
const TIf = 57355
const TIn = 57356
const TLocal = 57357
const TNil = 57358
const TNot = 57359
const TOr = 57360
const TReturn = 57361
const TRepeat = 57362
const TThen = 57363
const TTrue = 57364
const TUntil = 57365
const TWhile = 57366
const TRule = 57367
const TEqeq = 57368
const TNeq = 57369
const TLte = 57370
const TGte = 57371
const T2Comma = 57372
const T3Comma = 57373
const TIdent = 57374
const TNumber = 57375
const TString = 57376
const UNARY = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TAnd",
	"TBreak",
	"TDo",
	"TElse",
	"TElseIf",
	"TEnd",
	"TFalse",
	"TFor",
	"TFunction",
	"TIf",
	"TIn",
	"TLocal",
	"TNil",
	"TNot",
	"TOr",
	"TReturn",
	"TRepeat",
	"TThen",
	"TTrue",
	"TUntil",
	"TWhile",
	"TRule",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"T2Comma",
	"T3Comma",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'('",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'^'",
	"';'",
	"'$'",
	"'='",
	"','",
	"':'",
	"'.'",
	"'['",
	"']'",
	"'#'",
	"')'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:521

func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
	return string([]byte{byte(c)})
}

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 18,
	48, 32,
	49, 32,
	-2, 69,
	-1, 95,
	48, 33,
	49, 33,
	-2, 69,
}

const yyPrivate = 57344

const yyLast = 596

var yyAct = [...]uint8{
	25, 90, 52, 24, 47, 86, 58, 139, 67, 155,
	138, 115, 34, 54, 144, 56, 55, 136, 164, 33,
	134, 64, 65, 50, 63, 110, 111, 113, 108, 107,
	51, 67, 40, 41, 49, 157, 83, 84, 85, 43,
	44, 168, 93, 140, 133, 50, 97, 94, 48, 46,
	45, 106, 51, 101, 79, 80, 81, 23, 82, 82,
	87, 67, 69, 22, 108, 109, 152, 21, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 74, 75, 73, 72, 76, 32,
	167, 150, 10, 141, 135, 70, 71, 77, 78, 79,
	80, 81, 151, 82, 143, 146, 145, 148, 147, 50,
	39, 149, 50, 18, 150, 62, 51, 154, 153, 51,
	112, 27, 189, 38, 40, 41, 49, 26, 36, 99,
	98, 61, 57, 28, 96, 64, 156, 42, 93, 158,
	104, 159, 30, 22, 29, 40, 41, 21, 186, 20,
	181, 35, 170, 171, 169, 95, 180, 174, 165, 166,
	161, 102, 53, 1, 172, 37, 100, 173, 137, 175,
	69, 66, 177, 176, 89, 132, 31, 19, 9, 60,
	184, 183, 59, 3, 68, 185, 162, 4, 2, 0,
	188, 0, 74, 75, 73, 72, 76, 0, 0, 0,
	0, 0, 0, 70, 71, 77, 78, 79, 80, 81,
	27, 82, 38, 0, 0, 0, 26, 36, 0, 0,
	0, 114, 28, 0, 69, 0, 0, 0, 0, 0,
	0, 30, 91, 29, 40, 41, 21, 0, 68, 0,
	35, 0, 0, 0, 0, 0, 74, 75, 73, 72,
	76, 0, 92, 69, 37, 0, 88, 70, 71, 77,
	78, 79, 80, 81, 0, 82, 0, 68, 0, 0,
	0, 0, 0, 160, 0, 74, 75, 73, 72, 76,
	0, 0, 0, 0, 0, 0, 70, 71, 77, 78,
	79, 80, 81, 27, 82, 38, 0, 76, 0, 26,
	36, 0, 142, 0, 0, 28, 77, 78, 79, 80,
	81, 0, 82, 0, 30, 91, 29, 40, 41, 21,
	27, 0, 38, 35, 0, 0, 26, 36, 0, 0,
	0, 0, 28, 0, 69, 92, 178, 37, 0, 0,
	0, 30, 22, 29, 40, 41, 21, 0, 68, 0,
	35, 0, 0, 0, 0, 0, 74, 75, 73, 72,
	76, 0, 69, 0, 37, 0, 0, 70, 71, 77,
	78, 79, 80, 81, 0, 82, 68, 0, 0, 179,
	0, 0, 0, 0, 74, 75, 73, 72, 76, 0,
	69, 0, 187, 0, 0, 70, 71, 77, 78, 79,
	80, 81, 0, 82, 68, 0, 0, 163, 0, 0,
	0, 0, 74, 75, 73, 72, 76, 0, 69, 0,
	0, 0, 0, 70, 71, 77, 78, 79, 80, 81,
	0, 82, 68, 0, 0, 182, 0, 0, 0, 0,
	74, 75, 73, 72, 76, 0, 69, 0, 0, 0,
	0, 70, 71, 77, 78, 79, 80, 81, 0, 82,
	68, 0, 0, 105, 0, 0, 0, 0, 74, 75,
	73, 72, 76, 0, 69, 0, 103, 0, 0, 70,
	71, 77, 78, 79, 80, 81, 0, 82, 68, 0,
	0, 0, 0, 0, 0, 0, 74, 75, 73, 72,
	76, 0, 69, 0, 0, 0, 0, 70, 71, 77,
	78, 79, 80, 81, 0, 82, 68, 0, 0, 0,
	0, 0, 0, 0, 74, 75, 73, 72, 76, 0,
	0, 0, 0, 0, 0, 70, 71, 77, 78, 79,
	80, 81, 0, 82, 7, 11, 0, 0, 0, 0,
	15, 16, 14, 0, 17, 0, 0, 0, 6, 13,
	0, 0, 0, 12, 0, 0, 0, 0, 0, 0,
	0, 22, 0, 0, 0, 21, 74, 75, 73, 72,
	76, 0, 0, 0, 0, 5, 8, 70, 71, 77,
	78, 79, 80, 81, 0, 82,
}

var yyPact = [...]int16{
	-1000, -1000, 539, 11, -1000, -1000, 310, -1000, 112, -9,
	-2, -1000, 310, -1000, 310, 100, 99, 103, -1000, -1000,
	-1000, 310, -1000, -1000, -18, 498, -1000, -1000, -1000, -1000,
	-1000, -1000, -2, -1000, -1000, 310, 310, 310, 24, -1000,
	-1000, 200, -1000, 310, 31, 310, 98, -1000, 97, 111,
	-1000, -1000, 152, -1000, 470, 117, 442, 3, 15, 24,
	-25, -1000, 88, -21, -1000, 166, -44, 310, 310, 310,
	310, 310, 310, 310, 310, 310, 310, 310, 310, 310,
	310, 310, 310, 14, 14, 14, -1000, -11, -1000, -39,
	-1000, -5, 310, 498, -18, -1000, -2, 249, -1000, 90,
	-1000, -41, -1000, -1000, 310, -1000, 310, 310, 82, -1000,
	70, 34, 24, 310, -1000, -1000, 498, 58, 550, 267,
	267, 267, 267, 267, 267, 267, 13, 13, 14, 14,
	14, 14, -46, -1000, -1000, -14, -1000, 283, -1000, -1000,
	310, 220, -1000, -1000, -1000, 151, 498, -1000, 358, 12,
	-1000, -1000, -1000, -1000, -18, -1000, 150, 59, -1000, 498,
	-7, -1000, 145, 310, -1000, 148, -1000, -1000, 310, -1000,
	-1000, 310, 330, 147, -1000, 498, 141, 414, -1000, 310,
	-1000, -1000, -1000, 139, 386, -1000, -1000, -1000, 113, -1000,
}

var yyPgo = [...]uint8{
	0, 162, 188, 2, 187, 186, 183, 182, 179, 178,
	110, 6, 3, 0, 19, 89, 149, 177, 4, 176,
	5, 175, 12, 174, 1, 168,
}

var yyR1 = [...]int8{
	0, 1, 1, 1, 2, 2, 2, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 5, 5, 6, 6, 6, 7, 7,
	8, 8, 9, 9, 10, 10, 10, 11, 11, 12,
	12, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 14, 15,
	15, 15, 15, 17, 16, 16, 18, 18, 18, 18,
	19, 20, 20, 21, 21, 21, 22, 22, 23, 23,
	23, 24, 24, 24, 25, 25,
}

var yyR2 = [...]int8{
	0, 1, 2, 3, 0, 2, 2, 1, 2, 3,
	1, 3, 5, 4, 6, 8, 9, 11, 7, 3,
	4, 4, 2, 0, 5, 1, 2, 1, 1, 3,
	1, 3, 1, 3, 1, 4, 3, 1, 3, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 1, 1,
	1, 1, 3, 3, 2, 4, 2, 3, 1, 1,
	2, 5, 4, 1, 1, 3, 2, 3, 1, 3,
	2, 3, 5, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, -6, -4, 46, 19, 5, 47, -9,
	-15, 6, 24, 20, 13, 11, 12, 15, -10, -17,
	-16, 36, 32, 46, -12, -13, 16, 10, 22, 33,
	31, -19, -15, -14, -22, 40, 17, 54, 12, -10,
	34, 35, 25, 48, 49, 52, 51, -18, 50, 36,
	-22, -14, -3, -1, -13, -3, -13, 32, -11, -7,
	-8, 32, 12, -11, 32, -13, -16, 49, 18, 4,
	37, 38, 29, 28, 26, 27, 30, 39, 40, 41,
	42, 43, 45, -13, -13, -13, -20, 36, 56, -23,
	-24, 32, 52, -13, -12, -10, -15, -13, 32, 32,
	55, -12, 9, 6, 23, 21, 48, 14, 49, -20,
	50, 51, 32, 48, 55, 55, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -21, 55, 31, -11, 56, -25, 49, 46,
	48, -13, 53, -18, 55, -3, -13, -3, -13, -12,
	32, 32, 32, -20, -12, 55, -3, 49, -24, -13,
	53, 9, -5, 49, 6, -3, 9, 31, 48, 9,
	7, 8, -13, -3, 9, -13, -3, -13, 6, 49,
	9, 9, 21, -3, -13, -3, 9, 6, -3, 9,
}

var yyDef = [...]int8{
	4, -2, 1, 2, 5, 6, 25, 27, 0, 0,
	10, 4, 0, 4, 0, 0, 0, 0, -2, 70,
	71, 0, 34, 3, 26, 39, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 0, 0, 0, 0, 69,
	68, 0, 8, 0, 0, 0, 0, 74, 0, 0,
	78, 79, 0, 7, 0, 0, 0, 37, 0, 0,
	28, 30, 0, 22, 37, 0, 71, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 80, 0, 86, 0,
	88, 34, 0, 93, 9, -2, 0, 0, 36, 0,
	76, 0, 11, 4, 0, 4, 0, 0, 0, 19,
	0, 0, 0, 0, 72, 73, 40, 50, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 0, 4, 83, 84, 87, 90, 94, 95,
	0, 0, 35, 75, 77, 0, 13, 23, 0, 0,
	38, 29, 31, 20, 21, 4, 0, 0, 89, 91,
	0, 12, 0, 0, 4, 0, 82, 85, 0, 14,
	4, 0, 0, 0, 81, 92, 0, 0, 4, 0,
	18, 15, 4, 0, 0, 24, 16, 4, 0, 17,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 54, 47, 43, 3, 3,
	36, 55, 41, 39, 49, 40, 51, 42, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 46,
	38, 48, 37, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 52, 3, 53, 45, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 56,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 44,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:73
		{
			yyVAL.stmts = yyDollar[1].stmts
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:93
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:96
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:99
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:109
		{
			yyVAL.stmt = &ast.RuleStmt{Contents: yyDollar[2].token.Str}
			yyVAL.stmt.SetLine(yyDollar[2].token.Pos.Line)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:113
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].exprlist, Rhs: yyDollar[3].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].exprlist[0].Line())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:118
		{
			if _, ok := yyDollar[1].expr.(*ast.FuncCallExpr); !ok {
				yylex.(*Lexer).Error("parse error")
			} else {
				yyVAL.stmt = &ast.FuncCallStmt{Expr: yyDollar[1].expr}
				yyVAL.stmt.SetLine(yyDollar[1].expr.Line())
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:126
		{
			yyVAL.stmt = &ast.DoBlockStmt{Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].token.Pos.Line)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:131
		{
			yyVAL.stmt = &ast.WhileStmt{Condition: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:136
		{
			yyVAL.stmt = &ast.RepeatStmt{Condition: yyDollar[4].expr, Stmts: yyDollar[2].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].expr.Line())
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:141
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[6].token.Pos.Line)
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:151
		{
			yyVAL.stmt = &ast.IfStmt{Condition: yyDollar[2].expr, Then: yyDollar[4].stmts}
			cur := yyVAL.stmt
			for _, elseif := range yyDollar[5].stmts {
				cur.(*ast.IfStmt).Else = []ast.Stmt{elseif}
				cur = elseif
			}
			cur.(*ast.IfStmt).Else = yyDollar[7].stmts
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[8].token.Pos.Line)
		}
	case 16:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:162
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[9].token.Pos.Line)
		}
	case 17:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:167
		{
			yyVAL.stmt = &ast.NumberForStmt{Name: yyDollar[2].token.Str, Init: yyDollar[4].expr, Limit: yyDollar[6].expr, Step: yyDollar[8].expr, Stmts: yyDollar[10].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[11].token.Pos.Line)
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:172
		{
			yyVAL.stmt = &ast.GenericForStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[7].token.Pos.Line)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:177
		{
			yyVAL.stmt = &ast.FuncDefStmt{Name: yyDollar[2].funcname, Func: yyDollar[3].funcexpr}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[3].funcexpr.LastLine())
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:182
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: []string{yyDollar[3].token.Str}, Exprs: []ast.Expr{yyDollar[4].funcexpr}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.stmt.SetLastLine(yyDollar[4].funcexpr.LastLine())
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:187
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: yyDollar[4].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:191
		{
			yyVAL.stmt = &ast.LocalAssignStmt{Names: yyDollar[2].namelist, Exprs: []ast.Expr{}}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:197
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:200
		{
			yyVAL.stmts = append(yyDollar[1].stmts, &ast.IfStmt{Condition: yyDollar[3].expr, Then: yyDollar[5].stmts})
			yyVAL.stmts[len(yyVAL.stmts)-1].SetLine(yyDollar[2].token.Pos.Line)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:206
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: nil}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:210
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprlist}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:214
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:220
		{
			yyVAL.funcname = yyDollar[1].funcname
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:223
		{
			yyVAL.funcname = &ast.FuncName{Func: nil, Receiver: yyDollar[1].funcname.Func, Method: yyDollar[3].token.Str}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:228
		{
			yyVAL.funcname = &ast.FuncName{Func: &ast.IdentExpr{Value: yyDollar[1].token.Str}}
			yyVAL.funcname.Func.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:232
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			fn := &ast.AttrGetExpr{Object: yyDollar[1].funcname.Func, Key: key}
			fn.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.funcname = &ast.FuncName{Func: fn}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:241
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:244
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:249
		{
			yyVAL.expr = &ast.IdentExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:253
		{
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:257
		{
			key := &ast.StringExpr{Value: yyDollar[3].token.Str}
			key.SetLine(yyDollar[3].token.Pos.Line)
			yyVAL.expr = &ast.AttrGetExpr{Object: yyDollar[1].expr, Key: key}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:265
		{
			yyVAL.namelist = []string{yyDollar[1].token.Str}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:268
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[3].token.Str)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:273
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:276
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:293
		{
			yyVAL.expr = &ast.NumberExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:297
		{
			yyVAL.expr = &ast.Comma3Expr{}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:304
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:307
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:310
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:313
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "or", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:317
		{
			yyVAL.expr = &ast.LogicalOpExpr{Lhs: yyDollar[1].expr, Operator: "and", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:321
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:329
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:333
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = &ast.RelationalOpExpr{Lhs: yyDollar[1].expr, Operator: "~=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.expr = &ast.StringConcatOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:361
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:365
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:369
		{
			yyVAL.expr = &ast.ArithmeticOpExpr{Lhs: yyDollar[1].expr, Operator: "^", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:373
		{
			yyVAL.expr = &ast.UnaryMinusOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:377
		{
			yyVAL.expr = &ast.UnaryNotOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:381
		{
			yyVAL.expr = &ast.UnaryLenOpExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetLine(yyDollar[2].expr.Line())
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:387
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.Str}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:393
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:396
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:399
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:402
		{
			if ex, ok := yyDollar[2].expr.(*ast.Comma3Expr); ok {
				ex.AdjustRet = true
			}
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:411
		{
			yyDollar[2].expr.(*ast.FuncCallExpr).AdjustRet = true
			yyVAL.expr = yyDollar[2].expr
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:417
		{
			yyVAL.expr = &ast.FuncCallExpr{Func: yyDollar[1].expr, Args: yyDollar[2].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:421
		{
			yyVAL.expr = &ast.FuncCallExpr{Method: yyDollar[3].token.Str, Receiver: yyDollar[1].expr, Args: yyDollar[4].exprlist}
			yyVAL.expr.SetLine(yyDollar[1].expr.Line())
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:427
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = []ast.Expr{}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:433
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.exprlist = yyDollar[2].exprlist
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:439
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:442
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:447
		{
			yyVAL.expr = &ast.FunctionExpr{ParList: yyDollar[2].funcexpr.ParList, Stmts: yyDollar[2].funcexpr.Stmts}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.expr.SetLastLine(yyDollar[2].funcexpr.LastLine())
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:454
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: yyDollar[2].parlist, Stmts: yyDollar[4].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[5].token.Pos.Line)
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:459
		{
			yyVAL.funcexpr = &ast.FunctionExpr{ParList: &ast.ParList{HasVargs: false, Names: []string{}}, Stmts: yyDollar[3].stmts}
			yyVAL.funcexpr.SetLine(yyDollar[1].token.Pos.Line)
			yyVAL.funcexpr.SetLastLine(yyDollar[4].token.Pos.Line)
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:466
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:469
		{
			yyVAL.parlist = &ast.ParList{HasVargs: false, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:473
		{
			yyVAL.parlist = &ast.ParList{HasVargs: true, Names: []string{}}
			yyVAL.parlist.Names = append(yyVAL.parlist.Names, yyDollar[1].namelist...)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:480
		{
			yyVAL.expr = &ast.TableExpr{Fields: []*ast.Field{}}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:484
		{
			yyVAL.expr = &ast.TableExpr{Fields: yyDollar[2].fieldlist}
			yyVAL.expr.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:491
		{
			yyVAL.fieldlist = []*ast.Field{yyDollar[1].field}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:494
		{
			yyVAL.fieldlist = append(yyDollar[1].fieldlist, yyDollar[3].field)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:497
		{
			yyVAL.fieldlist = yyDollar[1].fieldlist
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:502
		{
			yyVAL.field = &ast.Field{Key: &ast.StringExpr{Value: yyDollar[1].token.Str}, Value: yyDollar[3].expr}
			yyVAL.field.Key.SetLine(yyDollar[1].token.Pos.Line)
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:506
		{
			yyVAL.field = &ast.Field{Key: yyDollar[2].expr, Value: yyDollar[5].expr}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:509
		{
			yyVAL.field = &ast.Field{Value: yyDollar[1].expr}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:514
		{
			yyVAL.fieldsep = ","
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:517
		{
			yyVAL.fieldsep = ";"
		}
	}
	goto yystack /* stack new state and value */
}
