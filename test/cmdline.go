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
	"gmme-golib/Utils/cmdline"
	"regexp"
)

//var o_cmdline *cmdline.SCMDLine = cmdline.New()

func main() {
	fmt.Println("cmdline test - beg:")

	//	xTestRe()
	xTest01()

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
// ------------------------------------------------------------------------------
// -- Test01 - process array of options like they came from the command line
// ------------------------------------------------------------------------------
func xTest01() {
	fmt.Println("xTest01 -- beg:")
	/*
		l_r := rand.New(rand.NewSource(time.Now().UnixNano()))
		min := 10
		max := 20
		fmt.Println(l_r.Intn(max-min+1) + min)
		fmt.Println(l_r.Intn(max-min+1) + min)
		fmt.Println(l_r.Intn(max-min+1) + min)
	*/
	/*
		fmt.Println(rand.Intn(100))
		fmt.Println(rand.Intn(100))
		fmt.Println(rand.Intn(100))

		l_chrstr := strings.Repeat("*", 10)
		fmt.Println("repeat:", l_chrstr)
	*/

	//---------------------------------------------------------------------------
	//-- create test args array
	var l_args []string

	l_args = append(l_args, "-testsubenv")
	l_args = append(l_args, "upis${USERPROFILE}")

	l_args = append(l_args, "-mailonerr")

	l_args = append(l_args, "-azSecret#{HIDDEN|SECRET}")
	l_args = append(l_args, "ueA8Q~9T_vN.tyF~SIA63HqjMuwq1aCCe4ttCaeM")

	l_args = append(l_args, "-mailonerr")

	l_args = append(l_args, "-mailrctxt")
	l_args = append(l_args, "this is a test")

	l_args = append(l_args, "-mailrcsmtp")
	l_args = append(l_args, "mail78.apmoller.net")

	//---------------------------------------------------------------------------
	//-- create cmdline object and process
	var l_cmdline *cmdline.SCMDLine = cmdline.NewSCMDLine()
	l_cmdline.SetDebug(true)
	l_cmdline.AddArgsArray(l_args)
	//	l_cmdline.AddArgsArray(os.Args[1:])
	l_cmdline.Dump()

	fmt.Println("xTest01 -- end:")
}

// ------------------------------------------------------------------------------
// ------------------------------------------------------------------------------
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
