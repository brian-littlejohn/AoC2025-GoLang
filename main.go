package main

func main() {
	testData := false
	var filename string
	if testData {
		filename = "testdata"
	} else {
		filename = "data"
	}
	day01a(readDataFile("./Day01/" + filename))
	day01b(readDataFile("./Day01/" + filename))
	day02a(readDataFile("./Day02/" + filename))
}
