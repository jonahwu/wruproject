package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//res, _ := exec.Command("ls", "/root").Output()
	//	/root/reference/etcd2.2/etcd-v2.2.2-linux-amd64/etcdctl ls
	res, _ := exec.Command("/root/reference/etcd2.2/etcd-v2.2.2-linux-amd64/etcdctl", "ls").Output()
	fmt.Println(string(res))
	listres := strings.Split(string(res), "\n")
	for i, key := range listres {
		fmt.Println(i, key)
		res2, _ := exec.Command("/root/reference/etcd2.2/etcd-v2.2.2-linux-amd64/etcdctl", "ls", key).Output()
		listres2 := strings.Split(string(res2), "\n")
		for _, key2 := range listres2 {
			if key2 != "" {
				fmt.Println("------", key2)
				//now i kill sub path
				_, _ = exec.Command("/root/reference/etcd2.2/etcd-v2.2.2-linux-amd64/etcdctl", "rmdir", key2).Output()
			}
		}
		//now kill root path
		_, _ = exec.Command("/root/reference/etcd2.2/etcd-v2.2.2-linux-amd64/etcdctl", "rmdir", key).Output()

	}

}
