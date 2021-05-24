package pathsanitize

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

const LCOV_KEYWORD = ":"

var trimString = ""
var lcov = false

func CreateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "path-sanitize",
		Short:   "Change file paths from Windows style to Unix style",
		Example: "ncli path-sanitize --lcov lcov.info",
		Args:    cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := fixPaths(args[0])
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	command.Flags().StringVarP(&trimString, "trim", "t", "", "Some substring to which will be strings trim to in every line. (example: trim 'bar' 'foo/bar/some'->'bar/some')")
	command.Flags().BoolVar(&lcov, "lcov", false, "Parsing file in lcov format")
	return command
}

func fixPaths(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if len(file) == 0 {
		return errors.New("empty file " + filename)
	}

	log.Println("Starting processing file " + filename)
	fileLines := strings.Split(string(file), "\n")
	fileLines = sanitize(fileLines)

	if trimString != "" {
		fileLines, err = trimFrom(fileLines)
		if err != nil {
			log.Println("[TRIM ERROR] " + err.Error())
		}
	}

	log.Println("Writing changed file content to " + filename)
	err = os.WriteFile(filename, []byte(strings.Join(fileLines, "\n")), 0600)
	if err != nil {
		return err
	}
	return nil
}

func sanitize(lines []string) []string {
	log.Println("Sanitizing lines to comply with Unix convention")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "\\", "/")
		result = append(result, line)
	}
	log.Println("Sanitized " + strconv.Itoa(len(result)) + " lines")
	return result
}

func trimFrom(content []string) ([]string, error) {
	lcovIndex := 0
	log.Println("Trimming lines. Removing characters before '" + trimString + "' in line")
	result := make([]string, 0, len(content))
	for _, line := range content {
		if len(line) == 0 {
			continue
		}
		if lcov {
			if line == "end_of_record" {
				continue
			}
			lcovIndex = strings.Index(line, LCOV_KEYWORD)
			if lcovIndex == -1 {
				return nil, errors.New("lcov file is not the right format")
			}
		}
		targetString := line
		trimIndex := strings.Index(line, trimString)
		if trimIndex != -1 {
			targetString = line[trimIndex:]
			if lcov {
				targetString = line[:(lcovIndex+1)] + targetString
			}
		}
		result = append(result, targetString)
	}
	return result, nil
}
