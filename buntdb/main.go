package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var db *buntdb.DB

func main() {
	InitDB()
	//readAndUpdate()

	//db.Update(func(tx *buntdb.Tx) error {
	//	tx.DeleteAll()
	//	return nil
	//})
	db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("", func(key, value string) bool {
			log.Printf("key %s,value ==>%s\n", key, value)
			return true
		})
		return nil
	})
	//var delKeys []string
	//var sum int
	//db.Update(func(tx *buntdb.Tx) error {
	//	tx.AscendKeys("block:*", func(key, value string) bool {
	//		if value == "" {
	//			delKeys = append(delKeys, key)
	//		}
	//		return true
	//	})
	//	for _, k := range delKeys {
	//		if _, err := tx.Delete(k); err != nil {
	//			return err
	//		}
	//		sum = sum + 1
	//
	//	}
	//	return nil
	//})
	//log.Println(sum)
	//TestIndexJson()
	//db.CreateIndex("block", "block:*", buntdb.IndexJSON("block.Timestamp"))
	//ViewAscend("block:*:HEAD")
	// 运行后关闭数据库
	defer db.Close()
}

// InitDB 初始化数据库
func InitDB() {

	var err error

	if db, err = buntdb.Open("data.db"); err != nil {
		log.Panic("数据库初始化错误" + err.Error())
	} else {
		log.Println("数据库初始化成功！")
	}
}

type block struct {
	Timestamp string
	Value     string
}

func TestIndexJson() {
	for i := 0; i < 100; i++ {
		log.Println("create stream " + strconv.Itoa(i))
		go update(i, uuid.NewString())
	}
}

func update(i int, uuid string) {
	db.Update(func(tx *buntdb.Tx) error {
		for i := 0; i < 20; i++ {
			key := "block:" + uuid + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			tx.Set(key, "1111111", nil)
			log.Println("Stream " + strconv.Itoa(i) + "create = Set->" + key)
			time.Sleep(time.Second)
		}
		tx.Set("block:"+uuid+":HEAD", "1111111", nil)

		log.Println("Stream " + strconv.Itoa(i) + "create = Set->" + "block:" + uuid + ":HEAD")

		return nil
	})
}

func ViewAscend(indexKey string) {
	log.Println("开始查询数据...")
	db.View(func(tx *buntdb.Tx) error {
		tx.DescendLessOrEqual("", `"Value":"1111"`, func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			return true
		})
		return nil
	})
}

func readAndUpdate() {
	file, err := os.Open("45052.txt")
	if err != nil {
		log.Println("文件打开失败！")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		if line, _, err := reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			db.Update(func(tx *buntdb.Tx) error {
				key := getKey()
				tx.Set(key, string(line), nil)
				log.Printf("Set %s ===> %s", key, string(line))
				return nil
			})
		}

	}
}

func getKey() string {
	return "block:" + uuid.NewString() + ":" + strconv.FormatInt(time.Now().Unix(), 10)
}
