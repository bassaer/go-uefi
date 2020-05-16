TARGET    = BOOTX64.EFI

.PHONY: build clean

all: run

BOOTX64.EFI: main.go
	GOOS=windows GOARCH=amd64 go build -o $@

goos.img: BOOTX64.EFI
	dd if=/dev/zero of=$@ bs=1k count=1440
	mformat -i $@ -f 1440 ::
	mmd -i $@ ::/EFI
	mmd -i $@ ::/EFI/BOOT
	mcopy -i $@ BOOTX64.EFI ::/EFI/BOOT

run: goos.img
	qemu-system-x86_64 -bios /usr/share/ovmf/OVMF.fd -net none -usbdevice disk::$<

clean:
	rm -f BOOTX64.EFI
