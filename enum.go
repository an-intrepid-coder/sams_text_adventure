package main

type Enum struct {
    labels []string
}

func MakeEnum(ls []string) Enum {
    return Enum{ls}
}

func (e *Enum) ContainsValue(n int) bool {
    return n >= 0 && n < len(e.labels)
}

func (e *Enum) ContainsLabel(s string) bool {
    for i := range e.labels {
        if e.labels[i] == s { 
            return true
        }
    }
    return false
}

func (e *Enum) NumLabels() int {
    return len(e.labels)
}

func (e *Enum) Label(i int) string {
    return e.labels[i]
}

/* Returns a pointer to the enumerated value of the string, if it exists
   within the enumeration. If not, it returns nil.  */
func (e *Enum) Value(s string) *int {
    for i := 0; i < len(e.labels); i++ {
        if s == e.labels[i] {
            return &i
        }
    }
    return nil
}

func (e *Enum) Labels() []string {
    s := make([]string, len(e.labels))
    copy(s, e.labels)
    return s
}

func (e *Enum) Values() []int {
    s := make([]int, len(e.labels))
    for i := range e.labels {
        s[i] = i
    }
    return s
}


