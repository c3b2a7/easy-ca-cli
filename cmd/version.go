/*
Copyright © 2023 c3b2a <c3b2a@qq.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"runtime"
)

var (
	version = "dev"
	date    = "unknown"
	commit  = "none"
	builtBy = "unknown"
)

//go:embed banner.txt
var banner []byte

func Version() string {
	return version
}

func VersionTemplate() string {
	buf := new(bytes.Buffer)
	buf.Write(banner)
	buf.WriteString(`{{with .Name}}{{printf "%s: A fast, simple certificate generation utility built in Go.\n" .}}{{end}}`)
	buf.WriteString("https://github.com/c3b2a7/easy-ca-cli\n\n")
	fmt.Fprintln(buf, "Version:", version)
	fmt.Fprintln(buf, "BuildDate:", date)
	fmt.Fprintln(buf, "GitCommit:", commit)
	fmt.Fprintln(buf, "BuiltBy:", builtBy)
	fmt.Fprintln(buf, "GoVersion:", runtime.Version())
	fmt.Fprintln(buf, "Compiler:", runtime.Compiler)
	fmt.Fprintln(buf, "Platform:", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	return buf.String()
}
