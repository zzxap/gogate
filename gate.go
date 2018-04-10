package main

import (
	"GateWay/HttpBusiness"
)

//jeeWah1El6oo
//cd ip2region项目根目录
//java -jar dbMaker-1.2.2.jar -src ./data/ip.merge.txt -region ./data/global_region.csv

func main() {

	HttpBusiness.StartHttpServer()
	//HttpBusiness.CheckIp()

}

/*
func readtxt() {
	curpath := public.GetCurDir()

	fi, err := os.Open(curpath + "\\ip.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fd, _ := ioutil.ReadAll(fi) //ioutil.ReadFile(path) //
	//fmt.Println(string(fd[:]))
	num := len(fd) / 5000

	for i := 0; i < num; i++ {
		from := i * 5000
		to := from + 5000
		//fmt.Println(string(fd[from:to]))

		topath := "D:\\GoProject\\src\\github.com\\mohong122\\ip2region\\data\\ip" + strconv.Itoa(i) + ".txt"
		public.WriteToFile(fd[from:to], topath)
	}
}
*/
