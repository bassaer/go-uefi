package main

type OutputStringFunc func(this *EfiSimpleTextOutputProtocol,str []rune) (int64)
type ClearScreenFunc func(this *EfiSimpleTextOutputProtocol) (int64)

type EfiInputKey struct {
	ScanCode [2]byte
	UnicodeChar [2]byte
}

type EfiSimpleTextInputProtocol struct {
	tmp int64;
	ReadKeyStroke *int64
}

type EfiSimpleTextOutputProtocol struct {
	tmp1 int64
	OutputString OutputStringFunc
	tmp2 [4]int64
	ClearScreen ClearScreenFunc
}

type EfiSystemTable struct {
	tmp1 [44]byte
	ConIN *EfiSimpleTextInputProtocol
	tmp2 int64
	ConOut *EfiSimpleTextOutputProtocol
}

func efi_main(ImageHandle uintptr, SystemTable *EfiSystemTable) {
	SystemTable.ConOut.ClearScreen(SystemTable.ConOut);
	msg := []rune("hello")
	SystemTable.ConOut.OutputString(SystemTable.ConOut, msg);
	for {}
}

func main() {
	var p uintptr
	efi_main(p, nil);
}

