package debugger

import (
	"sort"
	"sync"

	"github.com/kabi175/6502/model"
)

type DebuggerState uint8

const (
	Running DebuggerState = iota
	Paused
	Stopped
)

type Debugger struct {
	c       *sync.Cond
	PauseAt []uint16
	StopAt  uint16
	State   DebuggerState
}

func NewDebugger(bp []uint16, end uint16) *Debugger {
	sort.Slice(bp, func(i, j int) bool {
		return bp[i] < bp[j]
	})
	return &Debugger{
		c:       sync.NewCond(&sync.Mutex{}),
		PauseAt: bp,
		StopAt:  end,
	}
}

func (d *Debugger) doPause(pc uint16) bool {
	d.c.L.Lock()
	defer d.c.L.Unlock()
	index := sort.Search(len(d.PauseAt), func(i int) bool {
		return d.PauseAt[i] >= pc
	})
	if index < len(d.PauseAt) && d.PauseAt[index] == pc {
		return true
	}
	return false
}

func (d *Debugger) Run(cpu model.CPU, bus model.Bus16) {

	if d.IsEnd(cpu.GetPC()) {
		cpu.Quit()
		return
	}

	if d.doPause(cpu.GetPC()) {
		d.Pause()
	}

	d.c.L.Lock()
	defer d.c.L.Unlock()
	if d.State == Paused {
		d.c.Wait()
	}
}

func (d *Debugger) Pause() {
	d.c.L.Lock()
	defer d.c.L.Unlock()
	d.State = Paused
}

func (d *Debugger) Resume() {
	d.c.L.Lock()
	defer d.c.L.Unlock()
	if d.State == Running {
		return
	}
	d.State = Running
	d.c.Broadcast()
}

func (d *Debugger) Stop() {
	d.c.L.Lock()
	defer d.c.L.Unlock()
	d.State = Stopped
}

func (d *Debugger) IsEnd(pc uint16) bool {
	d.c.L.Lock()
	defer d.c.L.Unlock()
	return d.StopAt == pc
}
