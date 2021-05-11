package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	//client, err := docker.NewClient("http://172.18.2.224:2375")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//createOpts := docker.CreateExecOptions{}
	//createOpts.AttachStdin = true
	//createOpts.AttachStdout = true
	//createOpts.AttachStderr = true
	//createOpts.Tty = true
	//createOpts.Cmd = []string{"sh", "-c", "echo '\x21\x21' > /tmp/aaa2"}
	//createOpts.Container = "5914f8671813"
	//
	//exec, err := client.CreateExec(createOpts)
	//if err != nil {
	//	log.Fatalf("create exec error: %v\n", err)
	//}
	//
	//// start exec
	//startOpts := docker.StartExecOptions{}
	//startOpts.Tty = true
	//startOpts.RawTerminal = true
	//startOpts.Detach = false
	//startOpts.ErrorStream = os.Stderr
	//startOpts.InputStream = os.Stdin
	//startOpts.OutputStream = os.Stdout
	//
	//err = client.StartExec(exec.ID, startOpts)
	//if err != nil {
	//	log.Fatalf("start exec error: %v\n", err)
	//}
	dockerurl := "http://172.18.2.224:2375/containers/5914f8671813/exec"
	data := url.Values{}
	data.Set("AttachStdin","false")
	data.Set("AttachStdout","true")
	data.Set("AttachStderr","true")
	data.Set("DetachKeys","ctrl-p,ctrl-q")
	data.Set("Tty","false")
	data.Set("Cmd","date")
	//var jsonbody = []byte(`{"AttachStdin": false,"AttachStdout": true,"AttachStderr": true,"DetachKeys": "ctrl-p,ctrl-q","Tty": false,- "Cmd": ["date"],- "Env": ["FOO=bar","BAZ=quux"]}`)
	b := bytes.NewBufferString(data.Encode())
	Container,_ := http.NewRequest("POST",dockerurl, b)
	Res,_ := http.DefaultClient.Do(Container)
	fmt.Println(Res)
}


