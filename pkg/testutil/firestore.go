package testutil

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

func NewFirestoreTestClient(ctx context.Context) (client *firestore.Client, killer func()) {
	if runtime.GOOS == "windows" {
		killer = LaunchEmulatorOnWindows()
	}
	client, err := firestore.NewClient(ctx, "ynufes-mypage")
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
	}

	return client, killer
}

// LaunchEmulatorOnWindows refer this page for following code.
// https://www.captaincodeman.com/unit-testing-with-firestore-emulator-and-go
func LaunchEmulatorOnWindows() (killer func()) {
	// command to start firestore emulator
	cmd := exec.Command("gcloud", "beta", "emulators", "firestore", "start", "--host-port=localhost")
	// this makes it killable
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	// we need to capture it's output to know when it's started
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// we're going to wait until it's running to start
	var wg sync.WaitGroup
	wg.Add(1)

	// by starting a separate go routine
	go func() {
		// reading it's output
		buf := make([]byte, 256, 256)
		for {
			n, err := stderr.Read(buf[:])
			if err != nil {
				// until it ends
				if err == io.EOF {
					break
				}
				log.Fatalf("reading stderr %v", err)
			}

			if n > 0 {
				d := string(buf[:n])

				// only required if we want to see the emulator output
				log.Printf("%s", d)

				// checking for the message that it's started
				if strings.Contains(d, "Dev App Server is now running") {
					wg.Done()
					return
				}

				// and capturing the FIRESTORE_EMULATOR_HOST value to set
				pos := strings.Index(d, FirestoreEmulatorHost+"=")
				if pos > 0 {
					host := d[pos+len(FirestoreEmulatorHost)+1 : len(d)-2]
					log.Println("HOST: " + host)
					err := os.Setenv(FirestoreEmulatorHost, host)
					if err != nil {
						fmt.Println("ERROR")
						log.Fatalf("setting env var %v", err)
					}
				}
			}
		}
	}()
	// wait until the running message has been received
	wg.Wait()
	return func() {
		// run command to kill firestore emulator, which is running on port 8080
		err = exec.Command("powershell", "taskkill /PID $(netstat -ano | findstr.exe 8080 | grep LISTENING | awk '{print $5}' ) /F").Run()
		if err != nil {
			log.Fatal("ERROR ON KILLING " + err.Error())
		}
	}
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"
