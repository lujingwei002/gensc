package gensc

type Config struct {
	GenResource    GenResourceConfig
	GenProtocol    GenProtocolConfig
	GenModel       GenModelConfig
	GenConst       GenConstConfig
	GenBehavior    GenBehaviorConfig
	GenApplication GenApplicationConfig
	Macro          MacroConfig
}

type GenResourceConfig struct {
	Module             string
	DocDir             string // gen/
	ResourceDir        string // resource/
	GenDir             string // gen/resource/
	SrcGenDir          string // src/gen/resource/
	ExcelDir           string // .
	RespositoryVersion string // 仓库版本
	Force              bool
}

type GenProtocolConfig struct {
	GenProtocolDir    string
	SrcGenProtocolDir string
	DocProtocolDir    string
}

type GenModelConfig struct {
	Module      string
	SrcGenDir   string // src/gen/model
	DocDir      string // doc/model
	GenProtoDir string // gen/model
}

type GenConstConfig struct {
	SrcGenDir string
	ExcelDir  string
	DocDir    string
}

type GenBehaviorConfig struct {
	Module    string
	SrcGenDir string
	DocDir    string
	GenXmlDir string
}

type GenApplicationConfig struct {
	SrcGenApplicationDir string
	DocDir               string
}

type MacroConfig struct {
	SrcDirs    []string
	SrcGenDir  string
	SrcTestDir string
	SrcDir     string
}
