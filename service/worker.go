package service

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"sync"

	"github.com/wangkebin/ai-code-reviewer/models"
)

func collector(fmetas <-chan models.FData, res chan<- string) error {
	var filemetas strings.Builder
	for f := range fmetas {
		filemetas.WriteString(f.Content)
		filemetas.WriteString("\n")
	}
	res <- filemetas.String()

	return nil
}

func get_file_contents(filepath string) string {
	// Read the file content
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filepath, err)
	}
	return "### {" + filepath + "}\n{" + string(data) + "}\n\n"
}

func traversal(startPath string, fdatas chan<- models.FData) error {
	visit := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}
		f := models.FData{
			Path:    path,
			Content: get_file_contents(path),
		}
		fdatas <- f

		return nil
	}
	return filepath.WalkDir(startPath, visit)
}

func Walk(cfg *models.Config) string {
	//workers := 2 * runtime.GOMAXPROCS(0)
	fmetas := make(chan models.FData)
	fres := make(chan string)
	//done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		traversal(cfg.StartPath, fmetas)
		close(fmetas)
	}()

	go func() {
		collector(fmetas, fres)
		wg.Done()
	}()
	wg.Wait()
	result := <-fres
	// for i := 0; i < workers; i++ {
	// 	go collectHashes(fmetas, done)
	// }
	return result
}
