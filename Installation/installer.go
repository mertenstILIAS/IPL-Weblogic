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

func installWeblogic(weblogicPath String) {

	temppath := `C:\temp\wls1036_generic.jar`
	cmd := exec.Command("java.exe -jar " + temppath)
	err := cmd.Run()
	fmt.Printf("Command finished with error: %v", err)

}
