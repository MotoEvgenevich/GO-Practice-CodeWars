package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptThis(t *testing.T) {
	{
		result := EncryptThis("")
		expect := ""
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EncryptThis("A wise old owl lived in an oak")
		fmt.Println(result)
		expect := "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EncryptThis("The more he saw the less he spoke")
		expect := "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EncryptThis("The less he spoke the more he heard")
		expect := "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EncryptThis("Why can we not all be like that wise old bird")
		expect := "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := EncryptThis("Thank you Piotr for all your help")
		expect := "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}

}

/*
	Expect(EncryptThis("")).Should(Equal(""))
	Expect(EncryptThis("A wise old owl lived in an oak")).Should(Equal("65 119esi 111dl 111lw 108dvei 105n 97n 111ka"))
	Expect(EncryptThis("The more he saw the less he spoke")).Should(Equal("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"))
	Expect(EncryptThis("The less he spoke the more he heard")).Should(Equal("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"))
	Expect(EncryptThis("Why can we not all be like that wise old bird")).Should(Equal("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"))
	Expect(EncryptThis("Thank you Piotr for all your help")).Should(Equal("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"))
*/
