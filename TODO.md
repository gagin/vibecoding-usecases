# Think about this one

Behavior-Driven Development (BDD):
Format: BDD uses “Given-When-Then” scenarios to define behavior, often as use case alternatives:
Given [context] (e.g., config file exists).

When [action] (e.g., run codecat -e jsonc).

Then [outcome] (e.g., output.txt created).

Storage: Text (Gherkin syntax in .feature files) or JSON/YAML for automation (e.g., Cucumber). Example:
json

{
  "scenario": "Default Concatenation",
  "given": ["Config exists", "Directory has .py files"],
  "when": ["Run codecat -e jsonc > output.txt"],
  "then": ["output.txt contains .py, .jsonc files"]
}

Relevance: BDD’s concise scenarios complement your JSONC basic_flow and solution fields. Gherkin-to-JSON converters exist, supporting programmatic access, but diagrams are less common.

