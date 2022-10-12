package basics

import "fmt"

// This is not the idiomatic case where panic-recover can use
// instead of we can use error and let the flow decide what can be done
func recovered() {
	if err := recover(); err != nil {
		fmt.Println("error occured", err)
	}
}

func getAndPrintRequiredInfo(info *string) string {
	defer recovered()
	if info == nil || len(*info) == 0 {
		panic("Info is required param")
	}
	return fmt.Sprintf("Get info is %s", *info)
}
func PanicAndRecover() {

	fmt.Println("***************************Panic/Recover Section Start***************************")
	info := "This is required info"
	fmt.Println(getAndPrintRequiredInfo(&info))
	emptyInfo := ""
	fmt.Print(getAndPrintRequiredInfo(&emptyInfo))
	fmt.Print(getAndPrintRequiredInfo(nil))
	fmt.Println("***************************Panic/Recover Section Start***************************")

}
