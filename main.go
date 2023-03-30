package raid

import (
    "os/exec"
)

const tool = "/opt/MegaRAID/storcli/storcli64"

type raid struct{}

func (r raid) Hello() (string, error) {
    // /opt/MegaRAID/storcli/storcli64 show
    cmd := exec.Command(tool, "show")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
