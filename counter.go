package main

type Counter struct {
    count int
    step int
}

func MakeCounter(a, b int) Counter {
    return Counter{
        count: a,
        step: b,
    }
}

func (c *Counter) GetAndIncrement() int {
    i := c.count
    c.count += c.step
    return i
}

func (c *Counter) Count() int {
    return c.count
}

func (c *Counter) Step() int {
    return c.step
}

func (c *Counter) Increment() {
    c.count += c.step
}

