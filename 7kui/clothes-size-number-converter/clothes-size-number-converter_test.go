package kata

import (
	"fmt"
	"testing"
)

func TestSizeToNumber(t *testing.T) {
	fmt.Println(SizeToNumber("xxxl"))
}

/*
dotest("", 0, false, "Blank string is invalid ('')")
        dotest("xm", 0, false, "Cannot apply modifier for medium size ('xm')")
        dotest("xxxm", 0, false, "Cannot apply modifier for medium size ('xxxm')")
        dotest("xxxx", 0, false, "No base size provided ('xxxx')")
        dotest("ssss", 0, false, "Only one base size is allowed ('ssss')")
        dotest("hello world", 0, false, "Any other text is invalid ('Hello world') is invalid")
        dotest("sm", 0, false, "'sm' should be invalid (two bases)")
        dotest("ml", 0, false, "'ml' should be invalid (two bases)")
        dotest("lm", 0, false, "'lm' should be invalid (two bases)")
        dotest("lx", 0, false, "'lx' should be invalid (base should be last)")
*/

/*
Random valid modifiers + base. Testing for 'xxxxxxxxxxxxxxxxxxxxxxxxxxxl'
Expected
    <int>: 0
to equal
    <int>: 94
*/
