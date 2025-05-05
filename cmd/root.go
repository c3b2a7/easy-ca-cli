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
	"github.com/c3b2a7/easy-ca-cli/cmd/gen"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "easy-ca-cli",
	Short: "easy-ca-cli is a certificate generator",
	Long: `The easy-ca-cli is a fast, simple certificate generation utility built in Go.
It serves as a CLI to the core [easy-ca](https://github.com/c3b2a7/easy-ca) library,
providing an accessible way to generate various certificate types with customizable parameters.

Complete documentation is available at https://github.com/c3b2a7/easy-ca-cli#easy-ca-cli`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(gen.NewGenCmd())
}
