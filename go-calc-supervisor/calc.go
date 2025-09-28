package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Request struct {
	Op    string
	A, B  int
	Reply chan Response
}

type Response struct {
	Result int
	Error  error
}

type Adder struct {
	inbox chan Request
	alive bool
	mu    sync.Mutex
}

func NewAdder() *Adder {
	return &Adder{
		inbox: make(chan Request, 10),
		alive: true,
	}
}

func (a *Adder) Run() {
	for req := range a.inbox {
		if (req.A > 0 && req.B > math.MaxInt64-req.A) ||
			(req.A < 0 && req.B < math.MinInt64-req.A) {
			select {
			case req.Reply <- Response{Error: fmt.Errorf("overflow")}:
			case <-time.After(10 * time.Millisecond):
			}
			a.mu.Lock()
			a.alive = false
			a.mu.Unlock()
			return
		}

		select {
		case req.Reply <- Response{Result: req.A + req.B}:
		case <-time.After(10 * time.Millisecond):
		}
	}
}

type Multiplier struct {
	inbox chan Request
	alive bool
	mu    sync.Mutex
}

func NewMultiplier() *Multiplier {
	return &Multiplier{
		inbox: make(chan Request, 10),
		alive: true,
	}
}

func (m *Multiplier) Run() {
	for req := range m.inbox {
		if m.wouldOverflow(req.A, req.B) {
			m.sendError(req.Reply)
			m.mu.Lock()
			m.alive = false
			m.mu.Unlock()
			return
		}

		select {
		case req.Reply <- Response{Result: req.A * req.B}:
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func (m *Multiplier) wouldOverflow(a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	result := a * b
	return result/a != b
}

func (m *Multiplier) sendError(reply chan Response) {
	select {
	case reply <- Response{Error: fmt.Errorf("overflow")}:
	case <-time.After(10 * time.Millisecond):
	}
}

type Calculator struct {
	adder       *Adder
	multiplier  *Multiplier
	restarts    map[string]int
	budget      int
	resetTicker *time.Ticker
	escalated   bool
	mu          sync.RWMutex
	done        chan struct{}
	wg          sync.WaitGroup
	stopped     bool
}

func NewCalculator() *Calculator {
	return &Calculator{
		restarts:    map[string]int{"adder": 0, "multiplier": 0},
		budget:      3,
		resetTicker: time.NewTicker(time.Minute),
		done:        make(chan struct{}),
	}
}

func (c *Calculator) Start() {
	c.adder = NewAdder()
	c.multiplier = NewMultiplier()

	c.wg.Add(2)
	go func() {
		defer c.wg.Done()
		c.adder.Run()
	}()
	go func() {
		defer c.wg.Done()
		c.multiplier.Run()
	}()

	go c.supervise()
}

func (c *Calculator) supervise() {
	for {
		select {
		case <-c.resetTicker.C:
			c.mu.Lock()
			c.budget = 3
			c.escalated = false
			c.mu.Unlock()

		case <-c.done:
			return
		}
	}
}

func (c *Calculator) Add(a, b int) (int, error) {
	req := Request{
		Op:    "add",
		A:     a,
		B:     b,
		Reply: make(chan Response, 1),
	}

	select {
	case c.adder.inbox <- req:
		select {
		case resp := <-req.Reply:
			if resp.Error != nil {
				c.restart("adder")
				return 0, resp.Error
			}
			return resp.Result, nil
		case <-time.After(100 * time.Millisecond):
			c.restart("adder")
			return 0, fmt.Errorf("timeout")
		}
	case <-time.After(10 * time.Millisecond):
		c.restart("adder")
		return 0, fmt.Errorf("agent busy")
	}
}

func (c *Calculator) Multiply(a, b int) (int, error) {
	req := Request{
		Op:    "mul",
		A:     a,
		B:     b,
		Reply: make(chan Response, 1),
	}

	select {
	case c.multiplier.inbox <- req:
		select {
		case resp := <-req.Reply:
			if resp.Error != nil {
				c.restart("multiplier")
				return 0, resp.Error
			}
			return resp.Result, nil
		case <-time.After(100 * time.Millisecond):
			c.restart("multiplier")
			return 0, fmt.Errorf("timeout")
		}
	case <-time.After(10 * time.Millisecond):
		c.restart("multiplier")
		return 0, fmt.Errorf("agent busy")
	}
}

func (c *Calculator) restart(agent string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.budget <= 0 {
		c.escalated = true
		return
	}

	c.restarts[agent]++
	c.budget--

	switch agent {
	case "adder":
		if c.adder != nil {
			close(c.adder.inbox)
		}
		c.adder = NewAdder()
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			c.adder.Run()
		}()
	case "multiplier":
		if c.multiplier != nil {
			close(c.multiplier.inbox)
		}
		c.multiplier = NewMultiplier()
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			c.multiplier.Run()
		}()
	}
}

func (c *Calculator) RestartCount(agent string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.restarts[agent]
}

func (c *Calculator) IsEscalated() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.escalated
}

func (c *Calculator) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.stopped {
		return
	}
	c.stopped = true

	c.stopSupervision()
	c.closeAgentChannel(&c.adder.inbox)
	c.closeAgentChannel(&c.multiplier.inbox)
}

func (c *Calculator) stopSupervision() {
	if c.done != nil {
		close(c.done)
	}
	if c.resetTicker != nil {
		c.resetTicker.Stop()
	}
}

func (c *Calculator) closeAgentChannel(ch *chan Request) {
	if ch != nil && *ch != nil {
		select {
		case <-*ch:
		default:
			close(*ch)
		}
		*ch = nil
	}
}
