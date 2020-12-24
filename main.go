package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var abcd = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

type data struct {
	Data oneList
}

type oneList struct {
	// PractiveList []panduanOne `json:"practice_list"`
	PractiveList []danxuanone `json:"practice_list"`
}

type panduanOne struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Answer  string `json:"correct_answer"`
}

type danxuanone struct {
	ID          string   `json:"id"`
	Content     string   `json:"content"`        // <p> 题目 </p>
	Answer      []string `json:"correct_answer"` // 答案编号
	Option      []option `json:"option"`         // 选项列表
	AnswerCount int      `json:"correct_count"`
}

type option struct {
	ID      string `json:"id"`
	Describ string `json:"describ"`
}

// func read() (msgList []panduanOne) {
// 	file, err := os.Open("./data/2.json")
// 	defer file.Close()
// 	if err != nil {
// 		fmt.Println("read file error: " + err.Error())
// 	}

// 	r := io.Reader(file)
// 	msg := panduan{}
// 	if err = json.NewDecoder(r).Decode(&msg); err != nil {
// 		fmt.Println("get json msg error: " + err.Error())
// 	}
// 	return msg.Data.PractiveList
// }

func read() (msgList []danxuanone) {
	file, err := os.Open("./data/3.json")
	defer file.Close()
	if err != nil {
		fmt.Println("read file error: " + err.Error())
	}

	r := io.Reader(file)
	msg := data{}
	if err = json.NewDecoder(r).Decode(&msg); err != nil {
		fmt.Println("get json msg error: " + err.Error())
	}
	return msg.Data.PractiveList
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// func panduanti() {
// 	msg := read()

// 	fmt.Println(len(msg))
// 	println()

// 	filename := "判断题.txt"

// 	var file *os.File
// 	var err error

// 	if checkFileIsExist(filename) { //如果文件存在
// 		file, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
// 	} else {
// 		file, err = os.Create(filename) //创建文件
// 	}

// 	if err != nil {
// 		fmt.Println("创建写入文件异常：" + err.Error())
// 		return
// 	}

// 	io.WriteString(file, "一、判断题\n") //写入文件(字符串)
// 	for i := range msg {
// 		question := strings.TrimPrefix(msg[i].Content, "<p>")
// 		question = strings.TrimSuffix(question, "</p>")
// 		question = fmt.Sprintf("%d、%s  ", i+1, question)
// 		io.WriteString(file, question) //写入文件(字符串)
// 		if msg[i].Answer == "1" {
// 			io.WriteString(file, "（ 对 ）\n")
// 		} else {
// 			io.WriteString(file, "（ 错 ）\n")
// 		}
// 		fmt.Println(question)
// 	}

// }

func danxuanti() {
	msg := read()

	fmt.Println(len(msg))
	println()

	filename := "单选题.txt"

	var file *os.File
	var err error

	if checkFileIsExist(filename) { //如果文件存在
		file, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		file, err = os.Create(filename) //创建文件
	}

	if err != nil {
		fmt.Println("创建写入文件异常：" + err.Error())
		return
	}

	io.WriteString(file, "二、单选题\n") //写入文件(字符串)
	for i := range msg {
		question := strings.TrimPrefix(msg[i].Content, "<p>")
		question = strings.TrimSuffix(question, "</p>")
		question = fmt.Sprintf("%d、%s  ", i+1, question)
		io.WriteString(file, question) //写入文件(字符串)

		// fmt.Println(question)

		var answerStr string
		var answer string
		var option string

		if len(msg[i].Answer) == 1 {
			answerStr = msg[i].Answer[0]
		} else {
			fmt.Println("告警列表异常，break！！！！！！！！")
			fmt.Println(msg[i])
			continue
		}

		for j := range msg[i].Option {
			if answerStr == msg[i].Option[j].ID {
				answer += abcd[j]
			}

			option += fmt.Sprintf("%s、%s\n", abcd[j], msg[i].Option[j].Describ)
		}

		io.WriteString(file, answer+"\n"+option+"\n")

	}

}

func duoxuanti() {
	msg := read()

	fmt.Println(len(msg))
	println()

	filename := "多选题.txt"

	var file *os.File
	var err error

	if checkFileIsExist(filename) { //如果文件存在
		file, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		file, err = os.Create(filename) //创建文件
	}

	if err != nil {
		fmt.Println("创建写入文件异常：" + err.Error())
		return
	}

	io.WriteString(file, "三、多选题\n") //写入文件(字符串)
	for i := range msg {
		question := strings.TrimPrefix(msg[i].Content, "<p>")
		question = strings.TrimSuffix(question, "</p>")
		question = fmt.Sprintf("%d、%s  ", i+1, question)
		io.WriteString(file, question) //写入文件(字符串)

		// fmt.Println(question)

		var answer string
		var option string

		if len(msg[i].Answer) != msg[i].AnswerCount {
			fmt.Println("告警列表异常，break！！！！！！！！")
			fmt.Println(msg[i].Answer)
			fmt.Println(len(msg[i].Answer))
			fmt.Println(msg[i].AnswerCount)
			continue
		}

		for j := range msg[i].Option {

			for k := range msg[i].Answer {
				if msg[i].Answer[k] == msg[i].Option[j].ID {
					answer += abcd[j]
				}
			}

			option += fmt.Sprintf("%s、%s\n", abcd[j], msg[i].Option[j].Describ)
		}

		io.WriteString(file, answer+"\n"+option+"\n")

	}

}

func main() {
	fmt.Println("start")
	// panduanti()
	duoxuanti()
}
