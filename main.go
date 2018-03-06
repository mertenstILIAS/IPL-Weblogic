package main

const (
	beahome      = "D:/ilias"
	jdkHome      = `C:/Program Files/Java/jdk1.7.0_80`
	nodeMgrPort  = "1000"
	weblogicPath = "D:/distri/WebLogic-10.3.6/wls1036_generic.jar"
	jdkPath      = "D:/distri/jdk-7u80-windows-x64.exe"
)

func main() {

	installJava(jdkPath)

	/*
		path := "./silent_template.xml"
		var _, err = os.Stat(path)

		if os.IsExist(err) {
			fmt.Print("File doesn't exist")

		} else {
			contentXMLFile, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}

			ctx := plush.NewContext()
			ctx.Set("beahome", beahome)
			ctx.Set("jdkhome", jdkHome)
			ctx.Set("nodeMgrPort", nodeMgrPort)

			s, err := plush.Render(string(contentXMLFile[:]), ctx)
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile("./silent.xml", []byte(s), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}*/

	installWeblogic(weblogicPath)

}
