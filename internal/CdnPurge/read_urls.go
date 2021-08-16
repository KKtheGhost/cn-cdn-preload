package CdnPurge
 
import (
	"os"
	"bufio"
	"strings"
)

//本函数用于读取urlFile中的所有Url，并保存为一个临时的[]string用于在之后赋值给变量UrlList
func ReadUrls(path string) (UrlList []string){
	var resUrlList []string
	UrlFile = path
	RawFileBytes, err := os.Open(UrlFile)
	if err != nil {
		loger.Println("File error.")
	} else {
		buf := bufio.NewScanner(RawFileBytes)
		for {
			if !buf.Scan() {
				break
			}
			line := buf.Text()
			line = strings.TrimSpace(line)
			resUrlList = append(resUrlList, line)
		}
	}
	return resUrlList
}