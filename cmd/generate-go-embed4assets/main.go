package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/f9n/generate-go-embed4assets/internal/template"
	"github.com/f9n/generate-go-embed4assets/pkg/files"
)

var (
	version   = "none"
	gitCommit = "none"
	buildDate = "none"
	buildUser = "none"
)

var (
	versionFlag     = flag.Bool("version", false, "print the current version")
	directoryFlag   = flag.String("directory", "gen/go", "where to find assets")
	fileFormatsFlag = flag.String("file-formats", ".json,.yaml,.yml", "evaluate these file format types")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf("version %v, gitCommit: %s, buildDate: %s, buildUser: %s\n", version, gitCommit, buildDate, buildUser)
		os.Exit(0)
	}

	files, err := files.FindByExtension(*directoryFlag, strings.Split(*fileFormatsFlag, ","), files.Endless)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot find files by extentions. Error: %s", err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Printf("[+] Generating %s\n", fmt.Sprintf("%s.embed.go", file))
		err := generateFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot generate go embed file for '%s'. Error: %s", file, err)
			os.Exit(1)
		}
	}
}

// generateFile generates a _embed.go file containing embed definitions.
func generateFile(filePath string) error {
	fileAbsDir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)

	goPackageName, err := getGoPackageName(filePath)
	if err != nil {
		return err
	}

	if goPackageName == "go" {
		goPackageName = "_go"
	}

	goEmbedVariableName := getGoEmbedVariableName(filePath)

	templateData := template.Data{
		ToolVersion:    version,
		File:           filePath,
		Filename:       fileName,
		GoPackageName:  goPackageName,
		GoVariableName: goEmbedVariableName,
	}

	embedFilename := fileName + ".embed.go"
	destFile := filepath.Join(fileAbsDir, embedFilename)
	return template.Process(destFile, templateData)
}

var goPackageRegexCaps = `package\s(\w+)\s.*`

func getGoPackageName(filePath string) (string, error) {
	goPackageName := ""
	fileAbsDir := filepath.Dir(filePath)
	fileDirName := filepath.Base(fileAbsDir)

	goFiles, err := files.FindByExtension(fileAbsDir, []string{".go"}, files.OneShot)
	if err != nil {
		return "", err
	}
	if len(goFiles) == 0 {
		goPackageName = fileDirName
		return goPackageName, nil
	}

	re := regexp.MustCompile(goPackageRegexCaps)

	for _, goFile := range goFiles {
		fileContent, err := os.ReadFile(goFile)
		if err != nil {
			return "", err
		}

		matches := re.FindStringSubmatch(string(fileContent))

		if len(matches) > 1 {
			goPackageName = matches[1]
			break
		}

	}

	return goPackageName, nil
}

func getGoEmbedVariableName(file string) string {
	fileName := filepath.Base(file)

	return camelCase(strings.Split(fileName, "."))
}

func camelCase(words []string) string {
	var ret []string
	caser := cases.Title(language.AmericanEnglish)

	for _, word := range words {
		ret = append(ret, caser.String(word))
	}
	return strings.Join(ret, "")
}
