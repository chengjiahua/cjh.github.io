package ansible

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-kratos/kratos/pkg/log"
)

//Command .
type Command struct {
}

//NewCmd new.
func NewCmd() (cmd *Command) {
	return &Command{}
}

//AnsibleRunCopy ansible copy.
func (c *Command) AnsibleRunCopy(hosts []string, srcFile string, destPath string) (err error) {
	hostPattern := strings.Join(hosts, ",")
	a := fmt.Sprintf("src=%s dest=%s", srcFile, destPath)
	log.Info("hostPattern=%s a=%s", hostPattern, a)
	cmd := exec.Command("ansible", "all", "-i", hostPattern+",", "-m", "copy", "-a", a)
	outPut, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf("ansible copy to %s error %v", hostPattern, err)
		return
	}
	log.Info(string(outPut))
	return
}

//AnsibleRunShell ansible shell.
func (c *Command) AnsibleRunShell(hosts []string, shellCmd string) (err error) {
	log.Info("ansible run shell %s", shellCmd)
	hostPattern := strings.Join(hosts, ",")
	cmd := exec.Command("ansible", "all", "-i", hostPattern+",", "-m", "shell", "-a", shellCmd)
	outPut, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf("ansible run shell on %s error %v", hostPattern, err)
		return
	}
	log.Info(string(outPut))
	return
}
