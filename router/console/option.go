package router

import (
	"regexp"
	"strings"
)

type OptionConfig struct {
	required  bool
	optional  bool
	mandatory bool
}

type Option struct {
	flags       string
	description string
	config      OptionConfig

	shortFlag string
	longFlag  string
}

func NewOption(config OptionConfig, flags string, description string) Option {
	if strings.Contains(flags, "<") {
		config.required = true
		config.optional = false
	}
	if strings.Contains(flags, "[") {
		config.required = false
		config.optional = true
	}

	shortFlag, longFlag := makeEachFlags(flags)

	return Option{
		flags:       flags,
		description: description,
		shortFlag:   shortFlag,
		config:      config,
		longFlag:    longFlag,
	}
}

func makeEachFlags(flags string) (string, string) {
	var shortFlag string
	var longFlag string

	regex := regexp.MustCompile("[ |,]+")
	flagSlice := regex.Split(flags, -1)

	if len(flagSlice) > 1 && regexp.MustCompile("^[[<]").FindString(flagSlice[1]) == "" {
		shortFlag, flagSlice = shift(flagSlice)
	}

	longFlag, flagSlice = shift(flagSlice)

	if shortFlag == "" && regexp.MustCompile("^-[^-]$").FindString(longFlag) != "" {
		shortFlag = longFlag
		longFlag = ""
	}

	return shortFlag, longFlag
}

func shift(result []string) (string, []string) {
	if len(result) < 1 {
		return "", make([]string, 0)
	}

	shifted := result[0:1]

	if len(shifted) < 1 {
		return "", make([]string, 0)
	}

	return shifted[0], result[1:]
}
