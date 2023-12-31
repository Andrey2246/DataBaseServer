package main

import (
	"bufio"
	"log"
	"net"

	containers "github.com/Andrey2246/containers" // это мой же модуль.
	// файл, который вы сейчас читаете доступен по ссылке github.com/Andrey2246/DataBaseServer
)

type serverDataBase struct {
	containers.DataBase
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

func (db *serverDataBase) handleConnection(conn net.Conn) {
	scanner := bufio.NewReader(conn)
	ans := ""
	conn.Write([]byte("Enter your login: "))           //один логин - одни контейнеры
	login, _ := bufio.NewReader(conn).ReadString('\n') //если логина раньше не было, он создается
	commands := new(containers.Arr)
	input(scanner, commands)
	for i := 0; ans != "exit" && conn != nil && i < 1000; i++ {
		db.Mutex.Lock()
		ans = db.Execute(commands, login) // критическая область
		db.Mutex.Unlock()
		conn.Write([]byte(ans + "\n"))
		input(scanner, commands)
	}
	conn.Close()
}

func main() {
	sock, err := net.Listen("tcp", ":6379") //  соединение tcp, порт - 6379
	if err != nil {
		log.Fatalln("conn messed up \n", err.Error())
		panic(err)
	}
	defer sock.Close()
	db := new(serverDataBase)
	db.Array = make(map[string]*containers.Arr)
	db.BST = make(map[string]*containers.Bst)
	db.HashMap = make(map[string]*containers.HashMap)
	db.Queue = make(map[string]*containers.Queue)
	db.Stack = make(map[string]*containers.Stack)
	db.Set = make(map[string]*containers.Set)
	log.Println("Server is up and ready")
	for {
		ln, err := sock.Accept()
		if err != nil {
			log.Fatalln("read messed up", err.Error())
			panic(err)
		}
		log.Println(ln.RemoteAddr(), "connected")
		go db.handleConnection(ln)
	}
}
