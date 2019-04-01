package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	ip := "0.0.0.0"
	host := "www.somesite.com"
	action := "add"

	var w HostWriter
	switch runtime.GOOS {
	case "windows":
		w = new(WindowsHostWriter)
	case "darwin":
		w = new(DarwinHostWriter)
	case "linux":
		w = new(LinuxHostWriter)
	default:
		t.Error("Not supported", runtime.GOOS)
		return
	}

	err := w.Write(ip, host, action)
	if err != nil {
		t.Error(err)
		return
	}

	filename := w.Filename()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	text := string(data)
	line := fmt.Sprintf("%s %s", ip, host)
	if !strings.Contains(text, line) {
		t.Error("Add rule to hosts file failed!")
	}
}

func TestDelete(t *testing.T) {
	ip := "0.0.0.0"
	host := "www.somesite.com"
	action := "delete"

	var w HostWriter
	switch runtime.GOOS {
	case "windows":
		w = new(WindowsHostWriter)
	case "darwin":
		w = new(DarwinHostWriter)
	case "linux":
		w = new(LinuxHostWriter)
	default:
		t.Error("Not supported", runtime.GOOS)
		return
	}

	err := w.Write(ip, host, action)
	if err != nil {
		t.Error(err)
		return
	}

	filename := w.Filename()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	text := string(data)
	line := fmt.Sprintf("%s %s", ip, host)
	if strings.Contains(text, line) {
		t.Error("Delete rule from hosts file failed!")
	}
}
