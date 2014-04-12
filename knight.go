package knight

import (
	"fmt"
    "strings"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var mTimes = make(map[string]time.Time)

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(fmt.Sprintf("Error walking %s: ", path), err)
	}
	if !info.IsDir() && strings.HasSuffix(path, "go") {
		mTime := info.ModTime()
		oldTime, ok := mTimes[path]
		if ok {
			if oldTime.Before(mTime) {
				fmt.Println(" * Detected change, reloading")
				os.Exit(3)
			}
		} else {
			mTimes[path] = mTime
		}
	}
	return nil
}

func reloaderLoop(root string) {
	for {
		filepath.Walk(root, visit)
		time.Sleep(500 * time.Millisecond)
	}
}

type Knight struct {
	root string
}

func NewKnight(root string) *Knight {
    return &Knight{root:root}
}

func (knight Knight) ListenAndServe(addr string, handler http.Handler) error {
	if reloaderEnv := os.Getenv("KNIGHT_RELOADER"); reloaderEnv != "true" {
		fmt.Printf(" * Knight serving on %s\n", addr)
		for {
			fmt.Println(" * Restarting with reloader")
			arg := []string{"run"}
            _, file := filepath.Split(os.Args[0])
            file = file + ".go"
            arg = append(arg, file)
			arg = append(arg, os.Args[1:]...)
			command := exec.Command("go", arg...)
			command.Env = append(command.Env, "KNIGHT_RELOADER=true")
            command.Env = append(command.Env, os.Environ()...)
            command.Stdout = os.Stdout
			err := command.Run()
			if err == nil {
				return nil
			}
		}
	} else {
		go func() {
			http.ListenAndServe(addr, handler)
		}()
		reloaderLoop(knight.root)
	}
    return nil
}
