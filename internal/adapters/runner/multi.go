package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/realfabecker/kevin/internal/core/domain"
	"github.com/realfabecker/kevin/internal/core/ports"
)

type multi struct{}

func NewMulti() ports.ParallelRunner {
	return &multi{}
}

func (m *multi) Run(command string, pll int, mFlags []map[string]string, log domain.LogType) {
	start := time.Now()

	if len(mFlags) > 0 {
		m.runParallelWithFlags(command, pll, mFlags, log)
	} else {
		m.runParallel(command, pll, log)
	}

	end := time.Since(start)
	fmt.Println("Done in", end.String(), "(", end.Seconds(), "s)")
}

func (m *multi) runParallel(command string, pll int, log domain.LogType) {
	var wg sync.WaitGroup
	wg.Add(pll)

	var ch = make(chan struct{}, pll)
	defer close(ch)
	for i := 0; i < pll; i++ {
		go func(l domain.LogType) {
			defer wg.Done()
			ch <- struct{}{}
			if err := m.runCmd(command, nil, l); err != nil {
				fmt.Println(fmt.Errorf("runPll: %w", err))
			}
			<-ch
		}(log)
	}
	wg.Wait()
}

func (m *multi) runParallelWithFlags(command string, pll int, mFlags []map[string]string, log domain.LogType) {
	var wg sync.WaitGroup
	wg.Add(len(mFlags))

	var counter int64

	var ch = make(chan struct{}, pll)
	defer close(ch)

	for i, v := range mFlags {
		go func(flags map[string]string, n int, l domain.LogType) {
			defer wg.Done()
			ch <- struct{}{}

			current := atomic.AddInt64(&counter, 1)

			if l == domain.LogTool {
				if fgs, _ := json.Marshal(flags); fgs != nil {
					fmt.Printf("Script: %d of %d - %s\n", current, len(mFlags), string(fgs))
				} else {
					fmt.Printf("Script: %d of %d\n", current, len(mFlags))
				}
			}

			if err := m.runCmd(command, flags, l); err != nil {
				fmt.Println(fmt.Errorf("runPll: %w", err))
			}

			<-ch
		}(v, i, log)
	}
	wg.Wait()
}

func (m *multi) runCmd(command string, flags map[string]string, log domain.LogType) error {
	act := strings.Split(command, " ")
	if len(flags) > 0 {
		for k, v := range flags {
			act = append(act, fmt.Sprintf(`--%s='%s'`, k, v))
		}
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", strings.Join(act, " "))
	} else {
		cmd = exec.Command(act[:1][0], act[1:]...)
	}
	if log == domain.LogScript {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
