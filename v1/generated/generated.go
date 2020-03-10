package generated

//go:generate genny -pkg=$GOPACKAGE -in ../templates/go_chain__once.go -out go_chain__once__generated.go gen "GoChain=Chain"
//go:generate genny -pkg=$GOPACKAGE -in ../templates/go_chain__some.go -out go_chain__some__generated.go gen "GoChain=Chain Some=Bytes,BUILTINS,interface{},struct{}"
//go:generate genny -pkg=$GOPACKAGE -in ../templates/go_chain__some_other.go -out go_chain__some_other__generated.go gen "GoChain=Chain Some=Bytes,BUILTINS,interface{},struct{} Other=Bytes,BUILTINS,interface{},struct{},time.Time"

type Bytes = []byte
