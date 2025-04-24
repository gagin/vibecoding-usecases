Vibecoding Use Cases
===================

Overview
--------

This repository defines ``USECASES.jsonc``, a standardized JSONC-based format for documenting use cases in small to mid-scale vibecoding projects. Vibecoding blends AI-assisted coding (e.g., Cursor, Aider, Roo Code) with rapid iteration, typically for 1–10 developers building CLI tools, web apps, or scripts. The format offers:

- **Human-readable documentation**: Converts to Markdown with Mermaid diagrams for clarity.
- **Programmatic access**: Parseable by AI agents (e.g., for code generation) or CI/CD pipelines (e.g., test automation).
- **Simplicity**: Lightweight schema, editable in text editors, with comments for context.

Unlike traditional SWE frameworks (e.g., UML, Agile), ``USECASES.jsonc`` is tailored for vibecoding’s lightweight, AI-driven workflows. See `COMPARISON.rst`_ for a detailed comparison.

.. _COMPARISON.rst: COMPARISON.rst

USECASES.jsonc Format
---------------------

The ``USECASES.jsonc`` file stores use cases as a JSON array, with each use case containing:

- **id**: Unique identifier (e.g., "UC-VIBE-001").
- **actors**: Entities involved (e.g., "Developer", "AI Agent").
- **description**: Brief summary of the use case.
- **preconditions**: Conditions before execution.
- **postconditions**: Expected outcomes.
- **basic_flow**: Step-by-step interaction sequence.
- **core_scenario**: High-level goal.
- **feature_specific_scenario**: Context for specific features.
- **needs**: Requirements or motivations.
- **choices_compromises_assumptions**: Rationale for design decisions.
- **solution**: Expected output or behavior (e.g., command, result).
- **mermaid_flow**: Mermaid syntax for flow diagrams.

JSONC (JSON with Comments) ensures parseability while allowing human-readable comments. See ``USECASES.jsonc`` for an example meta-use case.

Usage
-----

1. **Create Use Cases**:
   - Write use cases in ``USECASES.jsonc`` following the schema.
   - Example: Document a CLI command, web endpoint, or script behavior.

2. **Convert to Documentation**:
   - Use ``jsonc2md`` (forthcoming) to generate Markdown with Mermaid diagrams.
   - Example: ``go run jsonc2md -input USECASES.jsonc -output usecases.md``

3. **Generate Diagrams**:
   - Use ``jsonc2mmd`` (forthcoming) to extract Mermaid diagrams.
   - Example: ``go run jsonc2mmd -input USECASES.jsonc -output diagrams.mmd``

4. **Integrate with AI/Automation**:
   - Parse ``USECASES.jsonc`` with AI agents for code generation or analysis.
   - Use in CI/CD for test generation or feature validation.

Contributing
------------

Contributions are welcome! Submit issues or pull requests for:

- Schema enhancements to ``USECASES.jsonc``.
- Converter tools (e.g., ``jsonc2md``, ``jsonc2mmd``).
- Examples for vibecoding projects (CLI, web, scripts).

License
-------

MIT License. See ``LICENSE`` for details.

Contact
-------

Open an issue or use GitHub Discussions for questions.