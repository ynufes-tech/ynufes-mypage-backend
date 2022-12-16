package infra

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"testing"
)

func newFirestoreTestClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, "ynufes-mypage")
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
	}

	return client
}

func TestMain(m *testing.M) {
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

	var result int
	defer func() {
		err := cmd.Process.Signal(syscall.SIGKILL)
		if err != nil {
			log.Fatal("KILL FAIL: " + err.Error())
		}
		os.Exit(result)
	}()

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

	// now it's running, we can run our unit tests
	result = m.Run()
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"
