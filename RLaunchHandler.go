package main

import (
	"fmt"
    "bufio"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	"os/exec"
	"syscall"
    "github.com/kardianos/osext"
)

func pauseProgramForExit() {
	fmt.Println("Press enter to exit.");
	bufio.NewReader(os.Stdin).ReadString('\n');
}

func main() {
	fmt.Println("RBLXDev Launch Handler");
	
	if len(os.Args) != 2 {
		fmt.Println("Invalid arguments.");
		pauseProgramForExit()
		os.Exit(1);
	}

	argsClean := strings.Replace(os.Args[1], "/", "", -1);
	argsClean = strings.Replace(argsClean, ":", "", -1);
	argsClean = strings.Replace(argsClean, "rblxhueten", "", -1);
	
	var params []string = strings.Split(argsClean, "+");

	if(len(params) != 2) {
		fmt.Println("Invalid arguments.");
		pauseProgramForExit()
		os.Exit(1);
	}

	fmt.Println(params[0], params[1]);

	client := &http.Client{};

	req, err := http.NewRequest("GET", ("https://nobelium.xyz/user/checktoken/" + params[0] + "+" + params[1]), nil);
	if err != nil {
		fmt.Println("An error occurred.");
		pauseProgramForExit();
		os.Exit(1);
	}

	req.Header.Set("User-Agent", "RBLXhue/1.0 (Launcher)");

	resp, err := client.Do(req);
	if err != nil {
		fmt.Println("An error occurred.");
		os.Exit(1);
	}

	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		fmt.Println("An error occurred.");
		os.Exit(1);
	}


	if(string(body) != "true") {
		fmt.Println("Invalid token.");
		pauseProgramForExit();
		os.Exit(1);
	} else {

		fmt.Println("Starting client...");

		url := "http://rblxdev.pw/client/join/" + params[0] + "+" + params[1];
		
		folderPath, eferr := osext.ExecutableFolder();
		if eferr != nil {
			fmt.Println("An error occurred.");
			pauseProgramForExit();
			os.Exit(1);
		}
		os.Chdir(folderPath);
		
		cmd := exec.Command(`RobloxApp.exe`);
		cmd.SysProcAttr = &syscall.SysProcAttr{};
		cmd.SysProcAttr.CmdLine = `RobloxApp.exe -china "script('` + url + `')"`;
		cmd.Start();
	}

	os.Exit(0);
}
