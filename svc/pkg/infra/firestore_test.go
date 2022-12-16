package infra

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
	"ynufes-mypage-backend/svc/pkg/infra/writer"

	"cloud.google.com/go/firestore"
)

func TestFirestore(t *testing.T) {
	client := newFirestoreTestClient(context.Background())
	w := writer.NewUser(client.Collection("users"))
	test1 := genTestCase()
	err := w.Create(context.Background(), test1[0])
	if err != nil {
		t.Fatal(err)
	}
	assert.IsEqual(err, nil)

	r := reader.NewUser(client)
	u, err := r.GetByID(context.Background(), 1234)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func newFirestoreTestClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, "ynufes-mypage")
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
	}

	return client
}

func TestMain(m *testing.M) {
	// noinspection
	var result int
	if runtime.GOOS == "windows" {
		killer := launchEmulatorOnWindows()
		defer killer(&result)
	}
	// now it's running, we can run our unit tests
	result = m.Run()
}

// refer this page for following code.
// https://www.captaincodeman.com/unit-testing-with-firestore-emulator-and-go
func launchEmulatorOnWindows() (killer func(result *int)) {
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
	return func(result *int) {
		// run command to kill the emulator (Windows)
		err = exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(cmd.Process.Pid)).Run()
		if err != nil {
			log.Fatal("ERROR ON KILLING " + err.Error())
		}
		os.Exit(*result)
	}
}

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"
