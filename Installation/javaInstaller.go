package installer

import (
	"fmt"
	"log"
	"os/exec"
)

func installJava(jdkPath String) {
	cmd := exec.Command(jdkPath, "/s")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Installation Java JDK 7u80 completed")
}
