package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Version is a version of a program
var Version = "dev"

var kvSep = "="

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Remap version: %s\n", Version)
	},
}

var remapCmd = &cobra.Command{
	Use:   "remap map file",
	Short: "Remap just replaces all keys by it's values in a passed file and send it to stdout",
	Long: `How to use

	remap <map> <file>

	Remap just replaces all keys by it's values in a passed file and send it to stdout

	For example, we have a file file.conf:

		API_URL=<api_url>
		API_KEY=<api_key>

	We need to replace all placeholders (<api_url> and <api_key>) by specific value

	We create a file with specific values - the map file (map.conf):

		<api_url> = https://example.com
		<api_key> = kksdfo93204jkljJKHJKsdf

	Then we call remap:

		remap map.conf file.conf

	Result in stdout:

		API_URL=https://example.com
		API_KEY=kksdfo93204jkljJKHJKsdf

	Remap just replaces all keys by it's values in a passed file.
`,
	Args: cobra.ExactArgs(2),
	RunE: remap,
}

func remap(cmd *cobra.Command, args []string) error {
	mapFile, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer mapFile.Close()

	srcFile, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer srcFile.Close()

	oldnew := fetchReplacements(mapFile)
	r := strings.NewReplacer(oldnew...)

	replace(srcFile, r, os.Stdout)

	return nil
}

func fetchReplacements(r io.Reader) []string {
	s := bufio.NewScanner(r)
	data := []string{}

	for s.Scan() {
		str := s.Text()
		d := strings.SplitN(str, kvSep, 2)
		if len(d) < 2 {
			continue
		}

		data = append(data, strings.TrimSpace(d[0]))
		data = append(data, strings.TrimSpace(d[1]))
	}

	return data
}

func replace(r io.Reader, rep *strings.Replacer, w io.Writer) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintln(w, rep.Replace(s.Text()))
	}
}

func main() {
	remapCmd.AddCommand(versionCmd)
	remapCmd.Execute()
}
