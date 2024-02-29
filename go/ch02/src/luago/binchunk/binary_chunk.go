package binchunk

// 定义二进制格式
type binaryChunk struct {
	header                  //头部
	sizeUpvalues byte       //主函数upvalues数量
	mainFunc     *Prototype //主函数原型
}

// header头部信息
type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luaNum          float64
}

// 相关常量
const (
	LUA_SIGNATURE = "\x1bLua"
	//LUAC_VERSION     = 0x53 //书上的版本
	LUAC_VERSION = 0x54 //我的版本
	LUAC_FORMAT  = 0
	LUAC_DATA    = "\x19\x93\r\n\x1a\n"
	CINT_SIZE    = 4
	CSZIET_SIZE  = 8
	//INSTRUCTION_SIZE = 4
	INSTRUCTION_SIZE = 8
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

// 函数原型结构体
type Prototype struct {
	Source          string
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalue
	Protos          []*Prototype
	LineInfo        []uint32
	LocVars         []LocVar
	UpvalueNames    []string
}

// tag 常量
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x3
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

// Upvalue 表
type Upvalue struct {
	Instack byte
	Idx     byte
}

// 局部变量表
type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

// 解析二进制chunk函数
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        //校验头部
	reader.readByte()           //跳过Upvalue数量
	return reader.readProto("") //读取函数原型
}
