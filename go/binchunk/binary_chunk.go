package binchunk

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

type header struct {
	signature       [4]byte
	version         byte
	luacData        [6]byte
	cintSize        byte
	suzetSuze       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

type Prototype struct {
	Source          string
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalues
	Protos          []*Prototype
	LineInfo        []uint32
	LocVars         []LocVar
	UpvalueNames    []string
}
