package process

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
)

func TestExecProcess(t *testing.T) {

	binary, lookErr := exec.LookPath("bash")
	assertEq(nil, lookErr)
	assertNotEq("", binary)

	args := []string{"-c", "ls -a -l -h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env) //syscall.Exec not support by windows

	if runtime.GOOS == "windows" {
		assertNotEq(nil, execErr)
	} else {
		assertEq(nil, execErr)
	}

}

func execProcess() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
