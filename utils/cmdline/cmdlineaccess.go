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
	"strings"
)

// -----------------------------------------------------------------------------
// -- GetOpt()
// -----------------------------------------------------------------------------
func (c *sCmdLine) GetOpt(a_opt string) (string, bool) {
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		return "", false
	}

	//--------------------------------------------------------------------------
	//-- see if opt exists
	l_val := ""
	l_opt, l_found := c.m_opts[strings.ToUpper(a_opt)]
	if l_found {
		l_val = l_opt.m_val
	}

	return l_val, l_found
}

// -----------------------------------------------------------------------------
// -- GetOptDef()
// -----------------------------------------------------------------------------
func (c *sCmdLine) GetOptDef(a_opt string, a_def string) string {
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		return a_def
	}

	//--------------------------------------------------------------------------
	//-- see if opt exists
	l_val := a_def
	l_opt, l_found := c.m_opts[strings.ToUpper(a_opt)]
	if l_found {
		l_val = l_opt.m_val
	}

	return l_val
}

// -----------------------------------------------------------------------------
// -- GetPath()
// -----------------------------------------------------------------------------
func (c *sCmdLine) GetPath(a_opt string) (string, bool) {
	return a_opt, false
}
func (c *sCmdLine) GetPathDef(a_opt string, a_def string) string {
	return a_opt
}

/*
#---------------------------------------------------------------------------
#-- getPathOpt
#---------------------------------------------------------------------------
def GetPathOpt(self, a_opt, a_defValue = None, a_allowSub = None, a_subValue = None):

	if not self.m_isInit: return None

	#-----------------------------------------------------------------------
	#-- see if value exists
	l_str = a_defValue
	l_val = self.GetOptValue(a_opt)
	if l_val is not None:
		#-------------------------------------------------------------------
		#-- save value and see if substitution is allowed
		l_str = l_val
		if a_allowSub is not None:
			if (a_allowSub == True) and (a_subValue is not None):
				l_i = l_str.find('%')
				if l_i > -1:
					l_str2 = l_str
					if l_i > 0 : l_str = l_str2[0:l_i-1]
					l_str = l_str + a_subValue
					l_str = l_str + l_str2[l_i:]


	#-----------------------------------------------------------------------
	#-- make sure their is '\' on end of string
	if (l_str is not None) and l_str != '':
		l_sep = os.path.sep
#            if not l_str.endswith(l_sep) :
		if l_str[-1] != l_sep:
			l_str = l_str + l_sep
			l_str = os.path.expanduser(l_str)

	return l_str
*/

// -----------------------------------------------------------------------------
// -- HasOpt() and IsOpt()
// -----------------------------------------------------------------------------
func (c *sCmdLine) HasOpt(a_opt string) bool {
	return c.IsOpt(a_opt)
}
func (c *sCmdLine) IsOpt(a_opt string) bool {
	//--------------------------------------------------------------------------
	//-- make sure object is fully initialized
	if !c.m_isInit {
		return false
	}

	//--------------------------------------------------------------------------
	//-- see if opt exists
	_, l_found := c.m_opts[strings.ToUpper(a_opt)]

	return l_found
}
