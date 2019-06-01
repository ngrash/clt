package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var printHelp bool
	var separator string
	var templateFile string

	flag.BoolVar(&printHelp, "h", false, "print this help")
	flag.StringVar(&separator, "sep", " ", "split input lines by `separator`")
	flag.StringVar(&templateFile, "t", "", "read template from `file`")
	flag.Parse()

	if printHelp {
		printUsage()
		os.Exit(0)
	}

	if len(templateFile) == 0 {
		log.Fatal("template file required (-t), see help (-h)")
	}

	template := read_template(templateFile)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		values := split_input_line(line, separator)
		rendered := render_template(template, values)
		fmt.Print(rendered)
	}

	if err := stdin.Err(); err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n  Input is read from stdin until EOF is reached.\n")
	fmt.Fprintf(os.Stderr, "\n  Available options:\n\n")
	flag.PrintDefaults()
}

func render_template(template []string, values []string) string {
	var rendered strings.Builder
	for _, line := range template {
		ansi_line := unescape_ansi_colors(line)
		rendered_line := render_template_line(ansi_line, values)
		rendered.WriteString(rendered_line + "\n")
	}
	return rendered.String()
}

func unescape_ansi_colors(src string) string {
	return strings.ReplaceAll(src, "\\033", "\033")
}

func render_template_line(template_line string, values []string) string {
	placeholder := regexp.MustCompile(`\$[0-9]+`)
	return placeholder.ReplaceAllStringFunc(template_line, func(placeholder string) string {
		value_index := placeholder_to_value_index(placeholder)
		if value_index >= len(values) {
			return "[UNKNOWN]"
		}

		return values[value_index]
	})
}

func placeholder_to_value_index(placeholder string) int {
	i, err := strconv.Atoi(placeholder[1:])
	if err != nil {
		log.Fatal("non-integer placeholder: ", placeholder)
	}

	return i
}

func read_template(path string) []string {
	template := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		template = append(template, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return template
}

func split_input_line(line string, sep string) []string {
	parts := make([]string, 0)

	for _, part := range strings.Split(line, sep) {
		trimmed_path := strings.TrimSpace(part)
		if len(trimmed_path) > 0 {
			parts = append(parts, trimmed_path)
		}
	}

	return parts
}
