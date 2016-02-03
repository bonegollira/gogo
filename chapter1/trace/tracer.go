package trace

import (
  "io"
  "fmt"
)

func New (w io.Writer) Tracer {
  return &tracer{out: w}
}

func Off () Tracer {
  return &nilTracer{}
}

type Tracer interface {
  Trace(...interface{})
}

type tracer struct {
  out io.Writer
}

func (t *tracer) Trace (a ...interface{}) {
  t.out.Write([]byte(fmt.Sprint(a...)))
  t.out.Write([]byte("\n"))
}

type nilTracer struct {
}

func (t *nilTracer) Trace (a ...interface{}) {
}

