package main // 声明 main 包，表明当前是一个可执行程序

import (
	"errors"
	"fmt" // 导入内置 fmt 包
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() { // main函数，是程序执行的入口
	fmt.Println("wcy 的小程序") // 在终端打印
	var ClassList []Class = InitClass()
	var relativePath string = "./excel"
	var result []Class
	//读出相对路径下的文件列表
	pathList, _ := ReadFiles(relativePath)
	// 获取班级数据结构体
	result, _ = GetClassRelativePath(ClassList, pathList)
	// 读取每个文件中的数据中"未返庆家庭成员"sheet页数据，将数据存入多维数组
	var tempList [][]string
	for i := 0; i < len(result); i++ {
		if result[i].RelativePath != "" {
			var temp = GetExcelDate(result[i].RelativePath, "未返庆家庭成员")
			tempList = append(tempList, temp...)
		}
	}
	fmt.Printf(" 查询出的数据：%v\n", tempList)
	// 导出一张整表Book1.xlxs
	SaveToExcel(tempList, "未返庆家庭成员")
}

// 读取文件名信息
func ReadFiles(relativePath string) (pathList []string, err error) {
	if relativePath == "" {
		return nil, errors.New("目录为空")
	}
	var List []string
	var FileInfo []os.FileInfo

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取  文件夹出错")
		return nil, errors.New("读取  文件夹出错")
	}

	for _, fileInfo := range FileInfo {
		// fmt.Println(fileInfo.Name())
		// //获取文件后缀
		fileSuffix := path.Ext(fileInfo.Name())
		// fmt.Println("fileSuffix =", fileSuffix)

		// //获取文件名
		// filenameOnly := strings.TrimSuffix(fileInfo.Name(), fileSuffix)
		// fmt.Println("filenameOnly =", filenameOnly)
		if fileSuffix == ".xlsx" {
			List = append(List, fileInfo.Name())
		}

	}
	// fmt.Println(List)
	return List, nil
}

type Class struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	RelativePath string `json:"relativePath"`
}

// 初始化班级列表
func InitClass() (list []Class) {
	name := [15]string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五"}
	var s []Class
	var class Class
	for i := 1; i <= 15; i++ {
		class = Class{
			Id:           i,
			Name:         "高一" + name[i-1] + "班",
			RelativePath: "",
		}
		s = append(s, class)
	}
	return s
}

// 把文件名和班级列表一一对应
func GetClassRelativePath(ClassList []Class, pathList []string) (result []Class, err error) {
	if len(ClassList) < 1 {
		return nil, nil
	}
	for i := 0; i < len(ClassList); i++ {
		for k := 0; k < len(pathList); k++ {
			if strings.Contains(pathList[k], ClassList[i].Name) {
				ClassList[i].RelativePath = pathList[k]
				// fmt.Println(ClassList[i])
			}
		}
	}

	return ClassList, nil
}

// 创建电子表格
func CreateExcel() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func OpenExcel(values []Class) {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.DeleteSheet("Sheet1")
	index := f.NewSheet("Sheet1")

	f.SetActiveSheet(index)
	f.SetCellValue("Sheet1", "A2", "Hello world1.")
	f.SetCellValue("Sheet1", "B2", "Hello world1.")
	f.SetCellValue("Sheet1", "A2", "002")
	f.SetCellValue("Sheet1", "B2", "Hello world.")
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)
	// 获取 Sheet1 上所有单元格
	rows, _ := f.GetRows("Sheet1")

	for _, row := range rows {
		fmt.Println(row)
		// for key, colCell := range row {
		// 	print(key, colCell, "\t")
		// }
	}
	// 根据指定路径保存文件
	// if err := f.SaveAs("Book1.xlsx.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

}

// 将数据存入表格
func SaveToExcel(values [][]string, sheetName string) {
	fmt.Println("11")
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if sheetName == "" {
		sheetName = "Sheet1"
	}
	f.DeleteSheet(sheetName)
	index := f.NewSheet(sheetName)

	f.SetActiveSheet(index)
	// 获取工作表中指定单元格的值
	// cell, err := f.GetCellValue(sheetName, "B2")
	// if err != nil {

	// 	println(err.Error())
	// 	return
	// }
	// println(cell)
	for i := 0; i < len(values); i++ {
		fmt.Println("存储数据：", values[i])
		err := f.SetSheetRow(sheetName, "A"+strconv.Itoa(i+1), &values[i])

		f.SetCellValue(sheetName, "A"+strconv.Itoa(i+1), i+1)
		// f.SetCellValue("Sheet1", "B1", "Hello world.")
		if err != nil {
			println(err.Error())
			return
		}
	}

	// 获取 Sheet1 上所有单元格
	rows, _ := f.GetRows(sheetName)

	for _, row := range rows {
		fmt.Println(row)
		// for key, colCell := range row {
		// 	print(key, colCell, "\t")
		// }
	}
	// 保存文件
	if err = f.Save(); err != nil {
		println(err.Error())
	}
	// 根据指定路径保存文件
	// if err := f.SaveAs("Book1.xlsx.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

}

// 获取各excel中的数据
func GetExcelDate(RelativePath string, sheetName string) (rowsList [][]string) {
	var temp [][]string
	var relativePath string = "./excel/" + RelativePath
	f, err := excelize.OpenFile(relativePath)
	if err != nil {
		fmt.Println(err)
		return temp
	}
	// 获取工作表中指定单元格的值
	//value := f.GetCellValue("data", "F1")

	//fmt.Println(value)

	// 获取 data 上所有单元格  data是我的工作表的名称
	rows, _ := f.GetRows(sheetName)

	for index, row := range rows {
		// fmt.Println(row) //[三层南走廊东 4000027 4000027 hkGxwgL6u03f]
		// fmt.Println("===============")

		//index的表格是第五列的时候，打印第五列的内容
		name, _ := f.GetCellValue(sheetName, "D"+strconv.Itoa(index+1))
		if index >= 5 && name != "" && !strings.Contains(name, "表格") {
			// fmt.Println("index", row)
			// fmt.Println("name", name)
			temp = append(temp, row)

			//fmt.Println(cell, "\t", index) //三层南走廊东   ，取到单元格的内容  “三层南走廊东”
		}
		// fmt.Println() //换一行
	}
	return temp
}
