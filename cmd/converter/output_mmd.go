package main
import (
    "fmt"
    "strings"
)
// GenerateMermaid converts use cases to standalone Mermaid files
func GenerateMermaid(useCases []UseCase, outputPath string) error {
    var builder strings.Builder
    for i, uc := range useCases {
        if uc.MermaidFlow != "" {
            if i > 0 {
                builder.WriteString("\n---\n")
            }
            builder.WriteString(fmt.Sprintf("%% Use Case %s: %s\n", uc.ID, uc.Description))
            builder.WriteString(uc.MermaidFlow)
            builder.WriteString("\n")
        }
    }
    return WriteFile(outputPath, builder.String())
}
