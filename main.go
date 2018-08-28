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
	Short: "Версия программы",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Версия %s\n", Version)
	},
}

var remapCmd = &cobra.Command{
	Use:   "remap map file",
	Short: "Заменяет одно на другое",
	Long: `Производит замену в файлах и вывожу результат в stdout

	map - путь до файла с мапкой для замены. Пример содержимого файла:

		<api_url> = https://api.com/

	В этом случае, если в файле [file] встретится строка "<api_url>",
	то она будет заменена на "https://api.com/"

	Здесь важно то, что ключ - это имя параметра, который нужно заменить.
	Как ключ написан, так и будет производиться поиск. 
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
