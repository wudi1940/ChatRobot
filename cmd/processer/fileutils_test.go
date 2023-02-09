package processer

import (
	"fmt"
	"testing"
)

func TestOpenFile(t *testing.T) {

	fileName := "/root/MyProject/ChatRobot/docs/pic/AI.jpg"

	file, err := OpenFile(fileName)
	if err != nil {
		return
	}

	fmt.Println(file)

}
