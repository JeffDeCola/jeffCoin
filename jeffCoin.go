// jeffCoin.go

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("Hi Jeff")

    /mnt/c/Windows/System32/cmd.exe /c "/mnt/c/Users/Jeff/Desktop/t-rex-0.14.4-win-cuda10.0/XZC-2miners-POOL.bat"

	cmd := "/mnt/c/Windows/System32/cmd.exe"
	args := []string{"/C", "/mnt/c/Users/Jeff/Desktop/t-rex-0.14.4-win-cuda10.0/XZC-2miners-POOL.bat"}
	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

}
