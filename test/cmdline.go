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

	fmt.Println("cmdline test - end:")
}

//	_ = xTestVariadic("-opt1", "value1")
//	_ = xTestVariadic("-opt2", "value2", "file2")
//	_ = xTestVariadic("-opt3", "value3", "file3", "tags3")
//	_ = xTestVariadic("-opt4", &("value4"), "", "tags")

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
	xTestAddArgsFile(l_cmdline, "cmdline-test01.opt")
	//l_cmdline.AddArgsFile("cmdline-test01.opt")
	//l_cmdline.AddArgsFile(".\\cmdline-test01.opt")
	//l_cmdline.AddArgsFile("~\\cmdline-test01.opt")
	l_cmdline.Dump()

	l_cmdline.Debug(false)
	xTestHasOpt(l_cmdline, "-mailrcsmtp")
	xTestHasOpt(l_cmdline, "-mailrcsmtpx")

	xTestIsOpt(l_cmdline, "-mailrcsmtp")
	xTestIsOpt(l_cmdline, "-mailrcsmtpx")

	xTestGetOpt(l_cmdline, "-mailrcsmtp")
	xTestGetOpt(l_cmdline, "-mailrcsmtpx")

	fmt.Println("xTest02 -- end:")
}
func xTestAddArgsFile(a_cmdline cmdline.CmdLine, a_file string) {
	a_cmdline.AddArgsFile(a_file)
}
func xTestGetOpt(a_cmdline cmdline.CmdLine, a_opt string) {
	l_val, l_found := a_cmdline.GetOptValue(a_opt)
	fmt.Println("get opt =", a_opt, ", value =", l_val, ", found =", l_found)
}
func xTestHasOpt(a_cmdline cmdline.CmdLine, a_opt string) {
	fmt.Println("has opt =", a_opt, ", found =", a_cmdline.HasOpt(a_opt))
}
func xTestIsOpt(a_cmdline cmdline.CmdLine, a_opt string) {
	fmt.Println("is opt =", a_opt, ", found =", a_cmdline.IsOpt(a_opt))
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
