package checks

import (
	"os/exec"
	"log"
	"strconv"
	"errors"
	"fmt"
)

func CheckDockerPool(okSize int) (error) {
	out, err := exec.Command("bash", "-c", "lvs -o data_percent,metadata_percent,LV_NAME --noheadings --units G --nosuffix | grep docker-pool").Output()
	if err != nil {
		msg := "Could not parse docker pool size: " + err.Error()
		log.Println(msg)
		return errors.New(msg)
	}

	isOk := isLvsSizeOk(string(out), okSize)
	if (!isOk) {
		return fmt.Errorf("Docker pool size is above: %v", strconv.Itoa(okSize))
	}
	return nil
}
