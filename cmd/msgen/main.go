// File: cmd/msgen/main.go
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "text/template"

    "github.com/spf13/cobra"

    "github.com/bocaletto-luca/msgen/internal/embed"
)

var (
    cfgPath string
    outDir  string
)

func main() {
    root := &cobra.Command{
        Use:   "msgen",
        Short: "Generate a persistent server daemon for Windows & Linux",
        RunE:  run,
    }
    root.Flags().StringVar(&cfgPath, "config", "config.yaml", "path to config.yaml")
    root.Flags().StringVar(&outDir, "out", "dist", "output directory for generated server")
    if err := root.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func run(cmd *cobra.Command, args []string) error {
    // 1) Load raw template
    tmplSrc := embed.ServerTmpl
    tmpl, err := template.New("server").Parse(string(tmplSrc))
    if err != nil {
        return fmt.Errorf("parse template: %w", err)
    }

    // 2) Read config.yaml
    cfgBytes, err := ioutil.ReadFile(cfgPath)
    if err != nil {
        return fmt.Errorf("read config: %w", err)
    }

    // 3) Prepare output dir
    if err := os.MkdirAll(outDir, 0755); err != nil {
        return fmt.Errorf("mkdir out: %w", err)
    }
    // Copy config.yaml
    if err := ioutil.WriteFile(filepath.Join(outDir, "config.yaml"), cfgBytes, 0644); err != nil {
        return fmt.Errorf("write config.yaml: %w", err)
    }

    // 4) Execute template
    buf := &bytes.Buffer{}
    data := struct {
        ConfigYAML string
        ModulePath string
    }{
        ConfigYAML: string(cfgBytes),
        ModulePath: "github.com/bocaletto-luca/msgen", // adjust if needed
    }
    if err := tmpl.Execute(buf, data); err != nil {
        return fmt.Errorf("render template: %w", err)
    }

    // 5) Write server.go
    outFile := filepath.Join(outDir, "server.go")
    if err := ioutil.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
        return fmt.Errorf("write server.go: %w", err)
    }

    fmt.Printf("Generated server in %s\n", outDir)
    fmt.Println("Now compile:")
    fmt.Println("  GOOS=windows GOARCH=amd64 go build -o dist/server.exe dist/server.go")
    fmt.Println("  GOOS=linux   GOARCH=amd64 go build -o dist/server     dist/server.go")
    return nil
}
