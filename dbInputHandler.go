package main

import (
	"strconv"
	containers "github.com/Andrey2246/containers"
)

func (db *DataBase) execute(commands *containers.Arr, password string) string {
	command := commands.Get(0)
	key := commands.Get(1)
	val := commands.Get(2)
	if command == "" {
		return "help"
	}
	switch command[0] {
	case 'H':
		{
			if db.h[password] == nil {
				db.h[password] = new(containers.HashMap)
			}
		}
	case 'S':
		{
			if command[1] != 'P' && db.set[password] == nil { // SP => SPUSH || SPOP
				db.set[password] = new(containers.Set)
			} else if db.s[password] == nil {
				db.s[password] = new(containers.Stack)
			}
		}
	case 'B':
		{
			if db.b[password] == nil {
				db.b[password] = new(containers.Bst)
			}
		}
	case 'Q':
		if db.q[password] == nil {
			db.q[password] = new(containers.Queue)
		}
	case 'A':
		{
			if db.a[password] == nil {
				db.a[password] = new(containers.Arr)
			}
		}
	}
	switch command {
	case "HSET":
		{
			err := db.h[password].Set(key, val)
			if err != nil {
				return err.Error()
			}
		}
	case "HGET":
		{
			st, err := db.h[password].Get(key)
			if err != nil {
				return err.Error()
			} else {
				return ("Value on key" + key + ":" + st)
			}
		}
	case "HDEL":
		{
			err := db.h[password].Del(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SPUSH":
		{
			db.s[password].Push(key)
		}
	case "SPOP":
		{
			st, err := db.s[password].Pop()
			if err != nil {
				return err.Error()
			} else {
				return st
			}
		}
	case "QPUSH":
		{
			db.q[password].Push(key)
		}
	case "QPOP":
		{
			st, err := db.q[password].Pop()
			if err != nil {
				return err.Error()
			} else {
				return st
			}
		}
	case "BADD":
		{
			err := db.b[password].Add(key)
			if err != nil {
				return err.Error()
			}
		}
	case "BPRINT":
		{
			return db.b[password].Print()
		}
	case "BDEL":
		{
			err := db.b[password].Del(key)
			if err != nil {
				return err.Error()
			}
		}
	case "BISMEM":
		{
			if db.b[password].IsMem(key) {
				return ("Your tree contains value\"" + key + "\"")
			} else {
				return ("Your tree does not contain value\"" + key + "\"")
			}
		}
	case "ASET":
		{
			keyInt, _ := strconv.Atoi(key)
			db.a[password].Set(keyInt, val)
		}
	case "AGET":
		{
			keyI, err := strconv.Atoi(key)
			if err != nil {
				return "key to array is an integer number, not" + key
			} else {
				db.a[password].Get(keyI)
			}
		}
	case "SADD":
		{
			err := db.set[password].ADD(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SREM":
		{
			err := db.set[password].Rem(key)
			if err != nil {
				return err.Error()
			}
		}
	case "SISMEM":
		{
			ans := db.set[password].IsMem(key)
			if ans {
				return "Your set contains value \"" + key + "\""
			}
			return "Your set does not contain value \"" + key + "\""
		}
	case "exit":
		{
			return "exit"
		}
	default:
		return "no command: " + command
	}
	return "OK"
}
