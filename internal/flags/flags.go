package flags

//Flags struct
type CLIFlags struct {
	Recursive bool
	Human bool
	All bool
}

//Flags struct getter
func Create(recursive, all, human bool) CLIFlags {
	return CLIFlags{
		Recursive: recursive,
		All: all,
		Human: human,
	}
}