package sui

import "encoding/binary"

const (
	Object1 = iota
	Pure1
	ForcedNonUniquePure1
)

type BuilderArg struct {
	Type   int
	Object interface{}
}

type BuilderArgObject struct {
	ObjectId ObjectId
}

type BuilderArgPure struct {
	Data []byte
}

type BuilderArgForceNonUniquePure struct {
	// todo
	Index uint32
}

type ProgrammableTransactionBuilder struct {
	Inputs   []CallArg
	Inputs1  map[BuilderArg]CallArg
	Commands []Command
}

func (pt *ProgrammableTransactionBuilder) finish() ProgrammableTransaction {
	return ProgrammableTransaction{
		Inputs:   pt.Inputs,
		Commands: pt.Commands,
	}
}

func (pt *ProgrammableTransactionBuilder) transferSui(recipient Address, amount uint64) {
	amountArg := pt.pure(amount)
	recipientArg := pt.pure(recipient)
	// split coins
	splitCoins := Command{
		Type: SplitCoins,
		Object: CommandSplitCoins{
			Src: Argument{
				Type:   GasCoin,
				Object: ArgumentGasCoin{},
			},
			Dst: []Argument{amountArg},
		},
	}
	coinArg := pt.command(splitCoins)
	// transfer
	transferObjects := Command{
		Type: TransferObjects,
		Object: CommandTransferObjects{
			Objects: []Argument{coinArg},
			Receipt: recipientArg,
		},
	}
	pt.command(transferObjects)
}

func (pt *ProgrammableTransactionBuilder) pure(value interface{}) Argument {
	switch value.(type) {
	case Address:
		addr := value.(Address)
		return pt.pureBytes(addr.Data[:], false)
	case uint64:
		v := value.(uint64)
		data := make([]byte, 8)
		binary.LittleEndian.PutUint64(data, v)
		return pt.pureBytes(data, false)
	default:
		panic("not supported")
	}
}

func (pt *ProgrammableTransactionBuilder) pureBytes(data []byte, forceSeparate bool) Argument {
	var builderArg BuilderArg
	if forceSeparate {
		// todo, not supported.
		builderArg.Type = ForcedNonUniquePure1
		builderArg.Object = BuilderArgForceNonUniquePure{Index: uint32(len(pt.Inputs))}
	} else {
		builderArg.Type = Pure1
		builderArg.Object = BuilderArgPure{Data: data}
	}
	callArg := CallArg{
		Type:   Pure,
		Object: CallArgPure{Data: data},
	}
	//pt.Inputs[builderArg] = callArg
	pt.Inputs = append(pt.Inputs, callArg)
	index := len(pt.Inputs) - 1
	argument := Argument{
		Type:   Input,
		Object: ArgumentInput{Input: uint16(index)},
	}
	return argument
}

func (pt *ProgrammableTransactionBuilder) command(command Command) Argument {
	pt.Commands = append(pt.Commands, command)
	index := len(pt.Commands) - 1
	return Argument{
		Type:   Result,
		Object: ArgumentResult{Result: uint16(index)},
	}
}
