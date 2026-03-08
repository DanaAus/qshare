package handlers

import (
	"context"
	"io"

	"github.com/schollz/progressbar/v3"
)

// ProgressReader wraps an io.Reader and updates a progress bar as data is read.
// It is context-aware and will stop reading if the context is cancelled.
type ProgressReader struct {
	reader io.Reader
	bar    *progressbar.ProgressBar
	ctx    context.Context
}

// NewProgressReader creates a new ProgressReader with a context.
func NewProgressReader(ctx context.Context, r io.Reader, bar *progressbar.ProgressBar) *ProgressReader {
	return &ProgressReader{
		reader: r,
		bar:    bar,
		ctx:    ctx,
	}
}

// Read implements the io.Reader interface.
func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	if err := pr.ctx.Err(); err != nil {
		return 0, err
	}

	n, err = pr.reader.Read(p)
	if n > 0 {
		_ = pr.bar.Add(n)
	}
	return n, err
}

// ProgressReadSeeker wraps an io.ReadSeeker and updates a progress bar as data is read.
// It is context-aware and will stop reading if the context is cancelled.
type ProgressReadSeeker struct {
	rs  io.ReadSeeker
	bar *progressbar.ProgressBar
	ctx context.Context
}

// NewProgressReadSeeker creates a new ProgressReadSeeker with a context.
func NewProgressReadSeeker(ctx context.Context, rs io.ReadSeeker, bar *progressbar.ProgressBar) *ProgressReadSeeker {
	return &ProgressReadSeeker{
		rs:  rs,
		bar: bar,
		ctx: ctx,
	}
}

// Read implements the io.Reader interface.
func (prs *ProgressReadSeeker) Read(p []byte) (n int, err error) {
	if err := prs.ctx.Err(); err != nil {
		return 0, err
	}

	n, err = prs.rs.Read(p)
	if n > 0 {
		_ = prs.bar.Add(n)
	}
	return n, err
}

// Seek implements the io.Seeker interface.
func (prs *ProgressReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return prs.rs.Seek(offset, whence)
}

// ProgressWriter wraps an io.Writer and updates a progress bar as data is written.
// It is context-aware and will stop writing if the context is cancelled.
type ProgressWriter struct {
	writer io.Writer
	bar    *progressbar.ProgressBar
	ctx    context.Context
}

// NewProgressWriter creates a new ProgressWriter with a context.
func NewProgressWriter(ctx context.Context, w io.Writer, bar *progressbar.ProgressBar) *ProgressWriter {
	return &ProgressWriter{
		writer: w,
		bar:    bar,
		ctx:    ctx,
	}
}

// Write implements the io.Writer interface.
func (pw *ProgressWriter) Write(p []byte) (n int, err error) {
	if err := pw.ctx.Err(); err != nil {
		return 0, err
	}

	n, err = pw.writer.Write(p)
	if n > 0 {
		_ = pw.bar.Add(n)
	}
	return n, err
}
