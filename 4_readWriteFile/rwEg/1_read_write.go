package rwEg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
Seek 设置下一次 Read 或 Write 的偏移量(offset)，它的解释取决于 whence
type Seeker interface {
    Seek(offset int64, whence int) (ret int64, err error)
}
SEEK_SET int = 0 //从文件的起始处开始设置 offset
SEEK_CUR int = 1 //从文件的指针的当前位置处开始设置 offset
SEEK_END int = 2 //从文件的末尾处开始设置 offset
*/

/*
ReadByte读取输入中的单个字节并返回。如果没有字节可读取，会返回错误。
ReadRune读取单个utf-8编码的字符，返回该字符和它的字节长度。如果没有有效的字符，会返回错误
WriteByte写入一个字节，如果写入失败会返回错误。
type ByteReader interface {
 ReadByte() (c byte, err error)
}
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
type ByteWriter interface {
    WriteByte(c byte) error
}
*/

/*
Read和Write函数不带文件偏移指针参数的读和写，以字节流方式读写，返回读写的字节数和状态，接口原型如下
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
*/
func wr() {
	f, _ := os.Create("rw.txt")
	defer f.Close()
	f.Write([]byte("Golang is a happy language(1)")) // 写入字节流方式
	f.Seek(0, os.SEEK_SET)                           // 将指针重置
	p := make([]byte, 2)                             // 读取 2 byte( len(buf)=2 )
	if _, err := f.Read(p); err != nil {
		log.Fatal("[F]", err)
	}
	fmt.Printf("read byte: \"%s\", len = %d byte\n", p, len(p)) //  /"输出引号
	p = make([]byte, 50)

	if _, err := f.Read(p); err != nil {
		if err != io.EOF { //忽略 EOF 错误
			log.Fatal("[F]", err)
		}
	}
	fmt.Printf("read byte: \"%s\", len = %d byte\n\n", p, len(p))
}

/*
ReaderAt和WriterAt带文件偏移指针参数的读和写，第二个参数是文件指针偏移量，以字节流方式读写，返回读写的字节数和状态，接口原型如下
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
WriteString是以字符串方式写入，原型
func WriteString(s string)(ret int, er error)
*/
func at() {
	f, _ := os.Create("rwat.txt")
	defer f.Close()
	f.WriteString("__Golang is a happy    language(2)") // 写入字符串方式
	f.WriteAt([]byte("pleasant"), 14)                   // 偏移 13byte 改写“happy”->“"pleasant”

	fi, _ := f.Stat()              //获取文件信息
	p := make([]byte, fi.Size()-2) //文件大小减去偏移值
	fmt.Println("filesize:", fi.Size())
	f.ReadAt(p, 2) //读文件，偏移 2 byte

	os.Stdout.Write(p)
	fmt.Printf("\n\n")
}

/*
ReadFrom() 从 r 中读取数据，直到 EOF 或发生错误。返回读取的字节数和 io.EOF 之外的其他错误。ReadFrom不会返回EOF错误
WriteTo() 将数据写入 w 中，直到没有数据可写或发生错误。返回写入的字节数和任何错误。
接口原型：
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
*/
func fromTo() {
	r := strings.NewReader("Golang is a happy language(3)") //创建一个 Reader
	fmt.Println("r =", r)

	w := bufio.NewWriter(os.Stdout) //创建一个 Writer

	w.ReadFrom(r) // w 一次性读取 r 的全部内容

	r.WriteTo(w) // r 一次性将内容写入 w 中

	fmt.Printf("w = ")
	w.Flush() // 输出内容
}

func ReadWriteFile() {
	wr()
	at()
	fromTo()
}
