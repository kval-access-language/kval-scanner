package kvalscanner

// Token represents a lexical token.
type Token int

//Tokens to scan for that satisfy the KVAL (Key Value Access Language) specification
const (
	// Special tokens
	ILLEGAL Token = iota    // ILLEGAL: Token is an illegal character
	EOF                     // EOF: Token signals the end of input
	WS                      // WS: Token signals whitespace has been found

	// Literals
	LITERAL                   // LITERAL: String literal discovered

   // Other Operators (in order)
   BUCKEY                  // BUCKKEY: >>>> Bucket to Key syntax for KVAL
   BUCBUC                  // BUCBUC: >> Bucket to Bucket syntax for KVAL
   KEYVAL                  // KEYVALL :: Key to Value syntax for KVAL
   ASSIGN                  // ASSIGN: => Assignment operator for KVAL renames

   // Single character operators
   USCORE                  // USCORE: _ Return unknown Key or Value 
   OPATT                   // OPATT: { Open a regular expression pattern
   CPATT                   // COATT: } Close a regular expression pattern

	// Keywords
   INS                     // INS: Insert capability of KVAL
   GET                     // GET: Get capability of KVAL
   LIS                     // LIS: LIS capability of KVAL
   DEL                     // DEL: Delete capability of KVAL
   REN                     // REN: Rename capability of KVAL

   // Regex
   REGEX                   // REGEX: {PATT} ANy regex pattern inside OPATT and CPATT
)

//Mapped values aiding in validation of KVAL strings
var KeywordMap = map[string]int{
   "INS":   0x1,         // Insert
   "ins":   0x1,

   "GET":   0x2,         // Get values
   "get":   0x2, 

   "LIS":   0x3,         // Check existence
   "lis":   0x3, 

   "DEL":   0x4,         // Delete values
   "del":   0x4,

   "REN":   0x5,         // Rename values   
   "ren":   0x5,
}
