package nodefindergo_test

import (
	"github.com/zxjsdp/nodefinder-go/utils"
	"testing"
	"io/ioutil"
	"os"
)

var (
	rawTree string = "((a, ((b, c), (ddd,\t e))), (f, g));\n\n"
	cleanTree string = "((a,((b,c),(ddd,e))),(f,g));"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// stringUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func Test_CleanElements(t *testing.T) {
	dirtyElements := []string{" \tsome\n", " \ndirty", "elements  \t\n"}
	expectedCleanElements := []string{"some", "dirty", "elements"}

	if len(utils.CleanElements(dirtyElements)) != len(expectedCleanElements) {
		t.Error("CleanElements (stringUtil.go): Should remove blanks, tab, and newlines in string elements")
	}
}

func Test_RemoveChar(t *testing.T) {
	expectedCleanString := "((a, ((b, c), (ddd, e))), (f, g));\n\n"
	result := utils.RemoveChar(rawTree, '\t')

	if (result != expectedCleanString) {
		t.Error("RemoveChar (stringUtil.go): Should remove specific rune from string")
	}
}

func Test_RemoveBlankChars(t *testing.T) {
	expectedReplacedString := "((a,_((b,_c),_(ddd,__e))),_(f,_g));__"

	result := utils.ReplaceBlankChars(rawTree)

	if (result != expectedReplacedString) {
		t.Error("ReplaceBlankChars (stringUtil.go): Should remove all blank runes")
	}
}

func Test_CheckRuneExistsInString(t *testing.T) {
	subString := "a,((b,c)"

	if (!utils.CheckSubStringExistsInString(cleanTree, subString)) {
		t.Error("CheckSubStringExistsInString (stringUtil.go): Check substring in string failed")
	}
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// stringUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func Test_IOUtils(t *testing.T) {
	expectedLineNum := 3

	lines := utils.ReadLines("calibration.txt")

	if (len(lines) != expectedLineNum) {
		t.Error("ReadLines (ioUtil.go): Line number in file not match")
	}
}

func Test_ReadContent(t *testing.T) {
	expectedContent := "b, ddd, >0.1<0.2\nb, c, $5\nf, #3"

	result := utils.ReadContent("calibration.txt")

	if (result != expectedContent) {
		t.Error("ReadContent (ioUtil.go): Read content failed")
	}
}

func Test_ReadCleanContent(t *testing.T) {
	expectedCleanContent := cleanTree

	result := utils.ReadCleanContent("input.nwk", []rune{' ', '\t', '\n'})

	if (result != expectedCleanContent) {
		t.Error("ReadCleanContent (ioUtil.go): ReadCleanContent failed")
	}
}

func Test_CheckFileExists(t *testing.T) {
	utils.CheckFileExists("input.nwk", "description", "usage")
}

func Test_WriteContent(t *testing.T) {
	contentToBeWritten := "the fox jumped over the lazy dog"
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "testWrite")
	defer os.Remove(tmpFile.Name())

	utils.WriteContent(tmpFile.Name(), contentToBeWritten)

	utils.CheckFileExists(tmpFile.Name(), "description", "usage")
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// sliceUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
