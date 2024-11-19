package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runner/types"
	"sync"

	"github.com/creack/pty"
)

func NewTerminal() *LocalTerminal {
	return &LocalTerminal{
		Terminal: &types.Terminal{
			Sessions: make(map[string]*types.TerminalSession),
		},
	}
}

type LocalTerminal struct {
	*types.Terminal
	Mu sync.Mutex
}

func (t *LocalTerminal) Init(id string, replID string, onData func(data string)) (*os.File, error) {
	cmd := exec.Command("zsh")

	ptyFile, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("failed to start PTY process: %w", err)
	}

	fmt.Println("Pty process spawned")
	t.Mu.Lock()
	t.Sessions[id] = &types.TerminalSession{
		Terminal: ptyFile,
		ReplID:   replID,
	}
	t.Mu.Unlock()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptyFile.Read(buf)
			if err != nil {
				fmt.Println("Error reading from pty:", err)
				break
			}
			onData(string(buf[:n]))
		}
	}()

	return ptyFile, nil
}

func (t *LocalTerminal) Write(id string, data string) {
	t.Mu.Lock()
	sessions, exists := t.Sessions[id]
	t.Mu.Unlock()

	if !exists {
		fmt.Printf("No terminal session found with id %s\n", id)
	}
	_, err := sessions.Terminal.Write([]byte(data))
	if err != nil {
		fmt.Printf("Error writing on terminal of id %s\n", id)
	} else {
		fmt.Println("Written data on terminal Sucessfully")
	}

}

func (t *LocalTerminal) Resize(id string, rows, cols int) {
	t.Mu.Lock()
	sessions, exists := t.Sessions[id]
	t.Mu.Unlock()

	if !exists {
		fmt.Printf("No terminal session found with id %s\n", id)
	}

	err := pty.Setsize(sessions.Terminal, &pty.Winsize{Cols: uint16(cols), Rows: uint16(rows)})

	if err != nil {
		fmt.Printf("failed to write in terminal wth id %s\n", id)
	} else {
		fmt.Printf("Resized terminal")
	}
}
func (t *LocalTerminal) Close(id string) {
	t.Mu.Lock()
	session, exists := t.Sessions[id]
	if exists {
		delete(t.Sessions, id)
	}
	t.Mu.Unlock()

	if !exists {
		fmt.Printf("No terminal session found with id %s\n", id)
		return
	}

	session.Terminal.Close()
	fmt.Printf("Closed terminal session %s\n", id)
}

func (t *LocalTerminal) Cleanup() {
	t.Mu.Lock()
	defer t.Mu.Unlock()

	for id, session := range t.Sessions {
		session.Terminal.Close()
		fmt.Printf("Closed terminal session %s during cleanup\n", id)
		delete(t.Sessions, id)
	}
	fmt.Println("All terminal sessions have been cleaned up.")
}
