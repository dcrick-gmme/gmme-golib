package main

//==============================================================================
// GMMEGoLib for Go/Golang
// Copyright (c) 2023, GMM Enterprises, LLC.
// Licensed under the GMM Software License
// All rights reserved
//==============================================================================
//	Author:	David Crickenberger
// -----------------------------------------------------------------------------
// Packages:
//		Utils::CmdLine
//
// 	Description:
//		Command line processor module.
//
//==============================================================================
//==============================================================================

//------------------------------------------------------------------------------
//
//------------------------------------------------------------------------------

import (
	"fmt"
	"gmme-golib/utils/cmdline"
	"regexp"
)

func main() {
	l_tests := map[string]bool{
		"testre": false,
		"test01": false,
		"test02": true,
	}

	fmt.Println("cmdline test - beg:")

	//--------------------------------------------------------------------------
	//-- create test args array
	if l_tests["testre"] {
		xTestRe()
	}
	if l_tests["test01"] {
		xTest01()
	}
	if l_tests["test02"] {
		xTest02()
	}

	//	_ = xTestVariadic("-opt1", "value1")
	//	_ = xTestVariadic("-opt2", "value2", "file2")
	//	_ = xTestVariadic("-opt3", "value3", "file3", "tags3")
	//	_ = xTestVariadic("-opt4", &("value4"), "", "tags")

	fmt.Println("cmdline test - end:")
}

//func xTestVariadic(a_opt string, a_val *string, a_args ...string) *string {
//	fmt.Println("opt = %s", a_opt)
//	fmt.Println("val = %s", a_val)
//	fmt.Println("args cnt = %d", len(a_args))

//	return nil
//}

/*
9439 5301 0935 5000 0047 05
9439 5301 0935 5000 0047 29
9439 5301 0935 5000 0046 99
9439 5301 0935 5000 0047 43
*/
// -----------------------------------------------------------------------------
// -- Test01 - process array of options like they came from the command line
// -----------------------------------------------------------------------------
func xTest01() {
	fmt.Println("xTest01 -- beg:")

	//--------------------------------------------------------------------------
	//-- create test args array
	l_args := []string{
		"-testsubenv",
		"upis${USERPROFILE}",
		"-mailonerr",
		"-azSecret#{HIDDEN|SECRET}",
		"ueA8Q~9T_vN.tyF~SIA63HqjMuwq1aCCe4ttCaeM",
		"-mailonerr",
		"-mailrctxt",
		"this is a test",
		"-mailrcsmtp",
		"mail78.apmoller.net",
	}

	fmt.Println(l_args)

	//--------------------------------------------------------------------------
	//-- create cmdline object and process
	var l_cmdline = cmdline.NewCmdLine()
	l_cmdline.Debug(true)
	l_cmdline.AddArgsArray(l_args)
	l_cmdline.Dump()

	fmt.Println("xTest01 -- end:")
}

// -----------------------------------------------------------------------------
// -- Test02 - process array of options like they came from the command line
// -----------------------------------------------------------------------------
func xTest02() {
	fmt.Println("xTest02 -- beg:")

	var l_cmdline = cmdline.NewCmdLine()
	l_cmdline.Debug(true)
	l_cmdline.AddArgsFile("cmdline-test01.opt")
	l_cmdline.AddArgsFile(".\\cmdline-test01.opt")
	l_cmdline.AddArgsFile("~\\cmdline-test01.opt")
	l_cmdline.Dump()

	fmt.Println("xTest02 -- end:")
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
func xTestRe() {
	var l_args0 []string
	l_args0 = append(l_args0, "-azSecret#{HIDDEN|SECRET}")
	l_args0 = append(l_args0, "ueA8Q~9T_vN.tyF~SIA63HqjMuwq1aCCe4ttCaeM")

	var testre = regexp.MustCompile(`(^.*)(\#\{(.*)\}$)`)
	var check1 = testre.MatchString(l_args0[0])
	//	var split = testre.Split(l_args0[0], -1)
	//	var findstr = testre.FindString(l_args0[0])
	var findallsub = testre.FindAllStringSubmatch(l_args0[0], -1)
	//	var find0 = testre.FindAllString(l_args0[0], 0)
	//	var find1 = testre.FindAllString(l_args0[0], 1)
	//	var find2 = testre.FindAllString(l_args0[0], 2)
	//	var find3 = testre.FindAllString(l_args0[0], 3)
	fmt.Println(check1)
	//	fmt.Println(split)
	//	fmt.Println(findstr)
	fmt.Println(findallsub)
	//	fmt.Println(find0)
	//	fmt.Println(find1)
	//	fmt.Println(find2)
	//	fmt.Println(find3)
	fmt.Println("this is a test!")
}
