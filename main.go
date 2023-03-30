package raid

import (
    "os/exec"
)

//const tool = "/opt/MegaRAID/storcli/storcli64"
const tool = "ls"

type PluginService struct{}

func (p PluginService) Hello(name string) (string, error) {
    // /opt/MegaRAID/storcli/storcli64 show
    cmd := exec.Command(tool, "-l", "/var/log")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
