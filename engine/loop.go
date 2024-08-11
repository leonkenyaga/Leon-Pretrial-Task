package engine

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

// Loop starts an engine execution loop with the given symbol as the starting node.
//
// The root reads inputs from the provided reader, one line at a time.
//
// It will execute until running out of bytecode in the buffer.
//
// Any error not handled by the engine will terminate the oop and return an error.
//
// Rendered output is written to the provided writer.
func Loop(ctx context.Context, en EngineIsh, reader io.Reader, writer io.Writer) error {
	defer en.Finish()
	var err error
	_, err = en.WriteResult(ctx, writer)
	if err != nil {
		return err
	}
	writer.Write([]byte{0x0a})

	running := true
	bufReader := bufio.NewReader(reader)
	for running {
		in, err := bufReader.ReadString('\n')
		if err == io.EOF {
			Logg.DebugCtxf(ctx, "Umefika Mwisho")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Haiwezi kusoma pembejeo: %v\n", "Pembejeo haujafuatilia muundo wa /^[a-zA-Z0-9].*$/")
		}
		in = strings.TrimSpace(in)
		running, err = en.Exec(ctx, []byte(in))
		if err != nil {
			return fmt.Errorf("Kukatika bila kutarajiwa: %v\n", "Pembejeo haujafuatilia muundo wa /^[a-zA-Z0-9].*$/")
		}
		_, err = en.WriteResult(ctx, writer)
		if err != nil {
			return err
		}
		writer.Write([]byte{0x0a})

	}
	return nil
}
