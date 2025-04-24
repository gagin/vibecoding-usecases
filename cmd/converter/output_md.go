package main
import (
    "fmt"
    "strings"
)
// GenerateMarkdown converts use cases to Markdown with Mermaid diagrams
func GenerateMarkdown(useCases []UseCase, outputPath string) error {
    var builder strings.Builder
    builder.WriteString("# Vibecoding Use Cases\n\n")
    for _, uc := range useCases {
        builder.WriteString(fmt.Sprintf("## Use Case %s: %s\n\n", uc.ID, uc.Description))
        builder.WriteString(fmt.Sprintf("ID: %s\n", uc.ID))
        builder.WriteString(fmt.Sprintf("Actors: %s\n", strings.Join(uc.Actors, ", ")))
        builder.WriteString(fmt.Sprintf("Description: %s\n\n", uc.Description))
        if len(uc.Preconditions) > 0 {
            builder.WriteString("Preconditions:\n")
            for _, pre := range uc.Preconditions {
                builder.WriteString(fmt.Sprintf("- %s\n", pre))
            }
            builder.WriteString("\n")
        }
        if len(uc.Postconditions) > 0 {
            builder.WriteString("Postconditions:\n")
            for _, post := range uc.Postconditions {
                builder.WriteString(fmt.Sprintf("- %s\n", post))
            }
            builder.WriteString("\n")
        }
        if len(uc.BasicFlow) > 0 {
            builder.WriteString("Basic Flow:\n")
            for i, step := range uc.BasicFlow {
                builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, step))
            }
            builder.WriteString("\n")
        }
        builder.WriteString(fmt.Sprintf("Core Scenario: %s\n", uc.CoreScenario))
        builder.WriteString(fmt.Sprintf("Feature-Specific Scenario: %s\n\n", uc.FeatureSpecificScenario))
        if len(uc.Needs) > 0 {
            builder.WriteString("Needs:\n")
            for _, need := range uc.Needs {
                builder.WriteString(fmt.Sprintf("- %s\n", need))
            }
            builder.WriteString("\n")
        }
        if len(uc.ChoicesCompromises) > 0 {
            builder.WriteString("Choices, Compromises, Assumptions:\n")
            for k, v := range uc.ChoicesCompromises {
                builder.WriteString(fmt.Sprintf("- %s: %s\n", k, v))
            }
            builder.WriteString("\n")
        }
        if len(uc.Solution) > 0 {
            builder.WriteString("Solution:\n")
            for k, v := range uc.Solution {
                builder.WriteString(fmt.Sprintf("- %s: %s\n", k, v))
            }
            builder.WriteString("\n")
        }
        if uc.MermaidFlow != "" {
            builder.WriteString("Flow Diagram:\nmermaid\n")             builder.WriteString(uc.MermaidFlow)             builder.WriteString("\n\n\n")
        }
    }
    return WriteFile(outputPath, builder.String())
}
