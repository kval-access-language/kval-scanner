package kvalscanner

// Token represents a lexical token.
type Token int

// Tokens to scan for that satisfy the KVAL (Key Value Access Language) specification
const (
   ILLEGAL Token = iota // Spacial token, ILLEGAL: Token is an illegal character
   EOF                  // Spacial token, EOF: Token signals the end of input
   WS                   // Spacial token, WS: Token signals whitespace has been found

   LITERAL // LITERAL: String literal discovered

   BUCKEY // Other operator, BUCKKEY: >>>> Bucket to Key syntax for KVAL
   BUCBUC // Other operator, BUCBUC: >> Bucket to Bucket syntax for KVAL
   KEYVAL // Other operator, KEYVALL :: Key to Value syntax for KVAL
   ASSIGN // Other operator, ASSIGN: => Assignment operator for KVAL renames

   USCORE // Single character operator, USCORE: _ Return unknown Key or Value
   OPATT  // Single character operator, OPATT: { Open a regular expression pattern
   CPATT  // Single character operator, COATT: } Close a regular expression pattern

   INS // Keyword, INS: Insert capability of KVAL
   GET // Keyword, GET: Get capability of KVAL
   LIS // Keyword, LIS: LIS capability of KVAL
   DEL // Keyword, DEL: Delete capability of KVAL
   REN // Keyword, REN: Rename capability of KVAL

   REGEX // Regular expression, REGEX: {PATT} ANy regex pattern inside OPATT and CPATT
)

// Mapped values aiding in validation of KVAL strings
var KeywordMap = map[string]int{
   "INS": 0x1, // Insert
   "ins": 0x1,

   "GET": 0x2, // Get values
   "get": 0x2,

   "LIS": 0x3, // Check existence
   "lis": 0x3,

   "DEL": 0x4, // Delete values
   "del": 0x4,

   "REN": 0x5, // Rename values
   "ren": 0x5,
}
