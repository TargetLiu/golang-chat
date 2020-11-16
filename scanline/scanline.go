package scanline

import (
	"bufio"
	"os"
	"strings"
)

//ScanLine 读取整行
func Read() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	return strings.Replace(input, "\n", "", -1)
}
