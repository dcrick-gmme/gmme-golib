/*
//==============================================================================
// GMMEGoLib for Go/Golang
// Copyright (c) 2023, GMM Enterprises, LLC.
// Licensed under the GMM Software License
// All rights reserved
//==============================================================================
//	Author:	David Crickenberger
// -----------------------------------------------------------------------------
// Package:
//		gmme-golib/utils/cmdline
//
// 	Description:
//		Command line processor package.
//
//==============================================================================
*/
package cmdline

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// -----------------------------------------------------------------------------
// -- AddArgsArray()
// -----------------------------------------------------------------------------
func (c *sCmdLine) AddArgsArray(a_args []string) {
	//--------------------------------------------------------------------------
	//-- if debug on dump a_args
	if c.m_dbgOn {
		fmt.Println("DBG-utils.cmdline.AddArgsArray - beg:")
		fmt.Println("DBG-a_args -- beg:")
		for l_i, l_arg := range a_args {
			fmt.Printf("[%d] = %s\n", l_i, l_arg)
		}
		fmt.Println("DBG-a_args -- end:")

		defer func() {
			fmt.Println("DBG-utils.cmdline.AddArgsArray - end:")
		}()
	}

	//--------------------------------------------------------------------------
	//-- process
	var l_file string = ""

	var l_i = 0
	for l_i < len(a_args) {
		//----------------------------------------------------------------------
		//-- pull option
		l_arg := a_args[l_i]
		if l_arg[0] == '-' || l_arg[0] == '/' {
			var l_opt = l_arg
			var l_val string = ""
			if (l_i + 1) < len(a_args) {
				//-- make sure next value is not an option
				l_arg = a_args[l_i+1]
				if len(l_arg) == 0 || (l_arg[0] != '-' && l_arg[0] != '/' && l_arg[0] != '@') {
					//-- we have a value with the option
					l_val = l_arg
					l_i++
				}
			}

			//------------------------------------------------------------------
			//-- determine opt and tags if any exists, then add item to map
			var l_opt2, l_tags = xCheckOptForTags(l_opt)
			c.m_opts[strings.ToUpper(l_opt2)] = newSOptItem(l_opt2, l_val, l_file, l_tags)
		} else if l_arg[0] == '@' {
			//------------------------------------------------------------------
			//-- we have an include file
			c.AddArgsFile(l_arg[1:])

			fmt.Printf("NEED TO ADD SUPPORT for include file to load = %s\n", l_arg)
		}
		l_i++
	}

	c.m_isInit = true
}

// -----------------------------------------------------------------------------
// -- AddArgsFile()
// -----------------------------------------------------------------------------
func (c *sCmdLine) AddArgsFile(a_file string) {
	l_func := "DBG-utils.cmdline.AddArgsFile"

	//--------------------------------------------------------------------------
	//-- if debug on dump a_args, and setup defer
	if c.m_dbgOn {
		fmt.Println(l_func + "::a_file => " + a_file + " - beg:")

		l_cwdPath, l_err := os.Getwd()
		if l_err != nil {
			log.Fatal(l_err)
		}
		fmt.Println(l_func + "::current folder is => " + l_cwdPath)

		defer func() {
			fmt.Println(l_func + "::a_file =>" + a_file + " - end:")
		}()
	}

	//--------------------------------------------------------------------------
	//-- deterine file to open and try to read contents
	l_file, l_err := xExpandUser(a_file)
	if l_err != nil {
		log.Fatal(l_err)
	}
	if c.m_dbgOn {
		fmt.Println(l_func, "Open file ==", l_file, "...")
	}

	//--------------------------------------------------------------------------
	//-- read contents
	l_optFile, l_err := os.ReadFile(l_file)
	if errors.Is(l_err, os.ErrNotExist) {
		log.Fatal(l_err)
	}
	if c.m_dbgOn {
		fmt.Print(string(l_optFile))
	}

	//--------------------------------------------------------------------------
	//-- process each line in the file
	l_line := strings.Builder{}

	l_pos := 0
	l_optFileLen := len(l_optFile)
	for l_pos < l_optFileLen {
		//----------------------------------------------------------------------
		//-- reset line
		l_line.Reset()
		l_line.Grow(100)

		//----------------------------------------------------------------------
		//-- pull a single line and see if its a comment line
		for l_pos < l_optFileLen && l_optFile[l_pos] != '\r' && l_optFile[l_pos] != '\n' {
			l_line.WriteByte(l_optFile[l_pos])
			l_pos++
		}
		for l_pos < l_optFileLen && (l_optFile[l_pos] == '\r' || l_optFile[l_pos] == '\n') {
			l_pos++
		}
		l_line2 := strings.TrimSpace(l_line.String())
		if len(l_line2) > 1 && l_line2[0] == '#' {
			continue
		}
		if len(l_line2) > 2 && l_line2[0] == '/' && l_line2[1] == '/' {
			continue
		}

		c.AddArgsLine(l_line2)
	}
}

// -----------------------------------------------------------------------------
// -- AddArgsLine()
// -----------------------------------------------------------------------------
func (c *sCmdLine) AddArgsLine(a_line string) {
	l_func := "DBG-utils.cmdline.AddArgsLine::"

	//--------------------------------------------------------------------------
	//-- if debug on dump a_args, and setup defer
	if c.m_dbgOn {
		fmt.Println(l_func, a_line, "- beg:")

		defer func() {
			fmt.Println(l_func, a_line, "- end:")
		}()
	}

	//--------------------------------------------------------------------------
	//-- process line
	var l_array []string

	l_splitCh := " "
	l_line := a_line
	for len(l_line) > 0 {
		//----------------------------------------------------------------------
		//-- split based on current split character
		l_split := strings.SplitN(l_line, l_splitCh, 2)
		l_array = append(l_array, l_split[0])

		//----------------------------------------------------------------------
		//-- stirp spaces and determine next split character
		l_tmp := ""
		if len(l_split) == 2 {
			l_tmp = strings.TrimRight(l_split[1], " ")
			l_splitCh = " "
			if l_tmp != "" {
				if l_tmp[0] == '"' || string(l_tmp[0]) == "'" {
					l_splitCh = string(l_tmp[0])
					l_tmp = l_tmp[1:]
				}
			}
		}
		l_line = l_tmp
	}

	c.AddArgsArray(l_array)
}

// =============================================================================
// -- private methods
// =============================================================================
// -- xCheckOptForTags
var oReTestForTags = regexp.MustCompile(`(^.*)(\#\{(.*)\}$)`)
var oReTestForTagsSplit = regexp.MustCompile(`[ ,\:\|]`)

func xCheckOptForTags(a_opt string) (string, []string) {
	var l_opt string

	//-- see if we have any tagss for the given opt
	var l_reAll = oReTestForTags.FindAllStringSubmatch(a_opt, -1)
	if l_reAll == nil || len(l_reAll[0]) != 4 {
		return a_opt, nil
	}

	//-- see if there are any tags to pull
	l_opt = l_reAll[0][1]
	if l_opt == "" || len(l_opt) == 0 {
		return a_opt, nil
	}

	var l_tags = make([]string, 0)
	l_tagsTmp := oReTestForTagsSplit.Split(l_reAll[0][3], -1)
	if len(l_tagsTmp) > 0 {
		l_haveHidden := false
		for _, l_tag := range l_tagsTmp {
			if strings.Contains("HIDE|HIDDEN|SECRET", l_tag) {
				if !l_haveHidden {
					l_tags = append(l_tags, "HIDE")
					l_haveHidden = true
				}
			} else {
				l_tags = append(l_tags, l_tag)
			}
		}
	}

	return l_opt, l_tags
}

// -----------------------------------------------------------------------------
// -- xExpandUser
func xExpandUser(a_path string) (string, error) {
	//--------------------------------------------------------------------------
	//-- setup tildesep
	l_sep := string(os.PathSeparator)
	l_tildesep := "~" + l_sep
	l_tildesepLen := len(l_tildesep)

	//--------------------------------------------------------------------------
	//-- get os home directory
	l_home, l_err := os.UserHomeDir()
	if l_err != nil {
		return "", l_err
	}

	if a_path == "~" {
		return l_home, nil
	} else if strings.HasPrefix(a_path, l_tildesep) {
		return l_home + l_sep + a_path[l_tildesepLen:], nil
	}

	return a_path, nil
}

// -----------------------------------------------------------------------------
// -- xSubEnv
func xSubEnv(a_str string) string {
	//--------------------------------------------------------------------------
	//-- see if anything to check need at least string of length 4 to process
	if len(a_str) < 4 {
		return ""
	}
	var l_str = a_str

	//--------------------------------------------------------------------------
	//-- see if we are going to sub the opt with environment values
	//	for strings.Index(l_str, "${") > -1 {
	for strings.Contains(l_str, "${") {
		l_p1 := strings.Index(l_str, "${")
		l_p2 := strings.Index(l_str[l_p1+2:], "}")
		if l_p1 >= 0 && l_p2 > l_p1+4 {
			l_envSub := strings.ToUpper(l_str[l_p1+2 : l_p1+l_p2+2])

			l_strSubstr := strings.Builder{}
			l_strSubstr.Grow(len(l_str) * 2)
			l_strSubstr.WriteString(l_str[:l_p1])
			l_strSubstr.WriteString(os.Getenv(l_envSub))
			l_strSubstr.WriteString(l_str[l_p1+l_p2+3:])
			l_str = l_strSubstr.String()
		}
	}

	return l_str
}
