package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	sentry "github.com/getsentry/sentry-go"
)

var help bool

func main() {
	flag.BoolVar(&help, "h", false, "show help message")
	flag.Parse()

	if help {
		basename := filepath.Base(os.Args[0])
		fmt.Printf("Usage: \n")
		fmt.Printf("\tenv SENTRY_DSN=dsn_url %v\n", basename)
		fmt.Printf("\t%v dsn_url\n", basename)
		fmt.Printf("Examples: \n")
		fmt.Printf("\tenv SENTRY_DSN=http://xxxxxx %v\n", basename)
		fmt.Printf("\t%v http://xxxxxx\n", basename)
		return
	}

	opt := sentry.ClientOptions{
		Debug: true,
	}
	if dsn := flag.Arg(0); dsn != "" {
		opt.Dsn = dsn
	}
	log.Printf("begin init sentry with opts=%+v", opt)
	err := sentry.Init(opt)
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
	sentry.CaptureException(fmt.Errorf("this is an test error, err=%w", os.ErrNotExist))
	log.Println("main exited")
}
