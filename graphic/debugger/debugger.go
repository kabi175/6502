package debugger

/*
import (
	"fmt"

	cpu "github.com/kabi175/6502/cpu6502"
	"github.com/veandco/go-sdl2/sdl"
)

type Debugger struct {
	window    *sdl.Window
	renderer  *sdl.Renderer
	IsRunning bool
	Quit      chan bool
}

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
)

func NewDebugger() *Debugger {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("Debugger", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	return &Debugger{
		window:    window,
		renderer:  renderer,
		IsRunning: true,
		Quit:      make(chan bool),
	}
}

func (d *Debugger) Close() {
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}

func (d *Debugger) HandleEvent() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			d.IsRunning = false
			d.Quit <- true
		}
	}

}
func (d *Debugger) Logger(c *cpu.Cpu6502) {
	fmt.Printf("A: %X\n", c.A.Get())
	fmt.Printf("X: %X\n", c.X.Get())
	fmt.Printf("Y: %X\n", c.Y.Get())
	fmt.Printf("SP: %X\n", c.SP.Get())
	fmt.Printf("PC: %X\n", c.PC.Get())
	fmt.Printf("Addr: %X, Operand: %X\n", c.Addr, c.Operand)
	fmt.Println("-------------------------------------------")
}

func (d *Debugger) Render(c *cpu.Cpu6502) {
	d.HandleEvent()
	d.Logger(c)
}
*/
