// +build !windows

package golang_tts

import(
	"fmt"
)

func SpeakText(text string){
	fmt.Printf("Saying %s\n", text)
}
