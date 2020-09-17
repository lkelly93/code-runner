package main

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type configSettings struct {
	hostname string
	rootLoc  string
}

func (cs *configSettings) setupInternalContainer() {
	mountProc(cs.rootLoc)
	mountSys(cs.rootLoc)
	changeHostName(cs.hostname)
	changeRoot(cs.rootLoc)
}

func changeHostName(name string) {
	must(syscall.Sethostname([]byte(name)))
}

func changeRoot(newRoot string) {
	must(syscall.Chroot(newRoot))
	must(os.Chdir("/"))
}

func mountProc(newRoot string) {
	source := "proc"
	fstype := "proc"
	target := filepath.Join(newRoot, "/proc")
	flags := uintptr(0)
	data := ""

	must(os.MkdirAll(target, 0755))
	must(syscall.Mount(source, target, fstype, flags, data))
}

func mountSys(rootLocation string) {
	source := "sysfs"
	target := filepath.Join(rootLocation, "/sys")
	fstype := "sysfs"
	flags := uintptr(0)
	data := ""

	os.MkdirAll(target, 0755)
	must(syscall.Mount(source, target, fstype, flags, data))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
