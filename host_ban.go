package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	DefaultIP     = "0.0.0.0"
	DefaultHost   = ""
	DefaultAction = "add"
)

type HostWriter interface {
	Write(ip, host, action string) error
	Filename() string
}

type WindowsHostWriter struct{}
type DarwinHostWriter struct{}
type LinuxHostWriter struct{}

func (hw WindowsHostWriter) Write(ip, host, action string) error {
	filename := hw.Filename()
	lineSeparator := "\r\n"
	err := Write(filename, lineSeparator, ip, host, action)
	return err
}

func (hw WindowsHostWriter) Filename() string {
	systemRoot := os.Getenv("windir")
	filename := systemRoot + "\\System32\\drivers\\etc\\hosts"
	return filename
}

func (hw DarwinHostWriter) Write(ip, host, action string) error {
	filename := "/etc/hosts"
	separator := "\n"
	err := Write(filename, separator, ip, host, action)
	return err
}

func (hw DarwinHostWriter) Filename() string {
	filename := "/etc/hosts"
	return filename
}

func (hw LinuxHostWriter) Write(ip, host, action string) error {
	filename := "/etc/hosts"
	separator := "\n"
	err := Write(filename, separator, ip, host, action)
	return err
}

func (hw LinuxHostWriter) Filename() string {
	filename := "/etc/hosts"
	return filename
}

func Write(filename, lineSeparator, ip, host, action string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	text := string(data)
	lines := strings.Split(text, lineSeparator)
	// search file whether already exist the host, delete it if exists
	if strings.Contains(text, host) {
		var lineIdx = -1
		for i, line := range lines {
			if strings.Contains(line, host) {
				lineIdx = i
				break
			}
		}
		if lineIdx >= 0 {
			var tmpLines []string
			tmpLines = append(tmpLines, lines[:lineIdx]...)
			tmpLines = append(tmpLines, lines[lineIdx+1:]...)
			lines = tmpLines
		}
	}
	if action == "add" {
		var line string
		line = fmt.Sprintf("%s %s", ip, host)
		lines = append(lines, line)
	}
	text = strings.Join(lines, lineSeparator)
	err = ioutil.WriteFile(filename, []byte(text), os.ModePerm)
	return err
}

func GetHostWriter() (hw HostWriter, err error) {
	switch runtime.GOOS {
	case "windows":
		hw = new(WindowsHostWriter)
	case "darwin":
		hw = new(DarwinHostWriter)
	case "linux":
		hw = new(LinuxHostWriter)
	default:
		err = errors.New(fmt.Sprintf("Not supported: %s", runtime.GOOS))
	}
	return
}

func main() {
	var ip = flag.String("ip", DefaultIP, "ip addr to redirect")
	var host = flag.String("host", DefaultHost, "some host, eg: www.somesite.com")
	var action = flag.String("action", DefaultAction, "add or delete")
	flag.Parse()
	*action = strings.ToLower(*action)
	if *action != "add" && *action != "delete" {
		fmt.Println(fmt.Sprintf("action only can specified as \"add\" or \"delete\", not \"%s\"", *action))
		return
	}
	if *host == "" {
		return
	}

	if w, err := GetHostWriter(); err != nil {
		fmt.Println(err)
	} else {
		if err := w.Write(*ip, *host, *action); err != nil {
			fmt.Println("error occurred!!!", err)
		}
	}
}
