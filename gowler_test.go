package gowler

import "testing"
import "reflect"

func TestGowler(t *testing.T) {
	/*
		log.SetPrefix("gowler.go: ")
		log.Println("Starting the gowler web crawler ...")
		flag.Parse()
		args := flag.Args()
		log.Println("Given args are", args)

		if len(args) != 2 {
			log.Println("gowler.go [url] [maxDepth]")
			os.Exit(1)
		}

		maxDepth, parseError := strconv.Atoi(args[1])

		if parseError != nil {
			log.Println(parseError)
		} else {
			log.Println("Input value is: ", maxDepth)
		} */
	const in = "http://www.aishraj.com"

	out := make([]string, 0)
	out = append(out, "http://github.com/aishraj/")
	scrapResult, e := ScrapLinks("http://www.aishraj.com")
	if e == nil && !reflect.DeepEqual(out, scrapResult) {
		t.Errorf("ScrapLinks(%v) = %v, want %v", in, scrapResult, out)
	}
}
