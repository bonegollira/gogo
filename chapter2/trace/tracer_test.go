package trace

import (
  "bytes"
  "testing"
)

func TestNew (t *testing.T) {
  var buf bytes.Buffer

  tracer := New(&buf)
  if tracer == nil {
    t.Error("Newからの戻り値が空です")
  } else {
    tracer.Trace("こんにちは、tracerパッケージ")
    if buf.String() != "こんにちは、tracerパッケージ\n" {
      t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
    }
  }
}

func TestOff (t *testing.T) {
  slientTracer := Off()
  slientTracer.Trace("データ")
}
