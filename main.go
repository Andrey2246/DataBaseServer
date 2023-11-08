package main

// go run main.go arr.go bst.go hash.go q.go set.go stac.go dbInputHandler.go
// go run *.go

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	containers "github.com/Andrey2246/containers"
)

type DataBase struct {
	s     map[string]*containers.Stack
	h     map[string]*containers.HashMap
	q     map[string]*containers.Queue
	set   map[string]*containers.Set
	b     map[string]*containers.Bst
	a     map[string]*containers.Arr
	mutex sync.Mutex
}

func input(scanner *bufio.Reader, arr *containers.Arr) { // разбивает строку на слова
	s1, s2, s3 := "", "", ""
	str, _ := scanner.ReadString('\n')
	where := 1
	for i, l := range str {
		if i == len(str)-1 {
			break
		}
		if l == ' ' {
			where += 1
			continue
		}
		switch where {
		case 1:
			s1 += string(l)
		case 2:
			s2 += string(l)
		case 3:
			s3 += string(l)
		}
	}
	arr.Set(0, s1)
	arr.Set(1, s2)
	arr.Set(2, s3)
}

func (db *DataBase) handleConnection(conn net.Conn) {
	scanner := bufio.NewReader(conn)
	ans := ""
	conn.Write([]byte("Enter your password: "))
	password, _ := bufio.NewReader(conn).ReadString('\n')
	commands := new(containers.Arr)
	input(scanner, commands)
	for i := 0; ans != "exit" && conn != nil && i < 10000; i++ {
		db.mutex.Lock()
		ans = db.execute(commands, password)
		db.mutex.Unlock()
		conn.Write([]byte(ans + "\n"))
		input(scanner, commands)
	}
	conn.Close()
}

func main() {
	sock, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalln("conn messed up", err.Error())
		panic(err)
	}
	defer sock.Close()
	db := new(DataBase)
	db.a = make(map[string]*containers.Arr)
	db.b = make(map[string]*containers.Bst)
	db.h = make(map[string]*containers.HashMap)
	db.q = make(map[string]*containers.Queue)
	db.s = make(map[string]*containers.Stack)
	db.set = make(map[string]*containers.Set)
	fmt.Println("Server is up and ready")
	for {
		ln, err := sock.Accept()
		if err != nil {
			log.Fatalln("read messed up", err.Error())
			panic(err)
		}
		fmt.Println(ln.RemoteAddr(), "connected")
		go db.handleConnection(ln)
	}
}
