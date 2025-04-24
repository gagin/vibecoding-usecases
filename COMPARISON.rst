Comparison to SWE Frameworks
============================

Overview
--------

``USECASES.jsonc`` is a JSONC-based format for documenting use cases in small to mid-scale vibecoding projects (1–10 developers, AI-assisted coding, rapid iteration). Unlike traditional software engineering (SWE) frameworks, it’s tailored for vibecoding’s lightweight, AI-driven workflows, balancing human-readable documentation (via Markdown/Mermaid) with programmatic access (for AI agents, CI/CD). This document compares ``USECASES.jsonc`` to common SWE frameworks, highlighting its advantages for vibecoding.

Frameworks Compared
------------------

- **UML/Unified Process (UP)**: Formal use cases with diagrams, used in enterprise settings.
- **Agile (User Stories)**: Lightweight stories for iterative development.
- **Use Case Driven Development (UCDD)**: Narrative use cases for object-oriented systems.
- **Behavior-Driven Development (BDD)**: Given-When-Then scenarios for testing.
- **Domain-Driven Design (DDD)**: Event-based modeling for complex domains.
- **CASE Tools**: Repository-based use cases with automation.

Comparison Table
----------------

+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| Framework               | Key Features        | Differences from USECASES.jsonc    | Fit for Vibecoding               | Pros/Cons                |
+=========================+=====================+====================================+==================================+==========================+
| UML/UP                  | ID, actors, flows,  | Adds alternative flows,            | Moderate: Detailed but heavy     | **Pros**: Comprehensive, |
|                         | preconditions, UML  | extensions, UML tools. Lacks       | for small teams. Tool dependency | visual.                  |
|                         | diagrams. Text/XML. | needs, rationale, Mermaid.         | clashes with rapid iteration.    | **Cons**: Complex, tool- |
|                         |                     |                                    |                                  | dependent.               |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| Agile (User Stories)    | As a [role], I want | Simpler, no flows or diagrams.     | High: Lightweight, but lacks     | **Pros**: Simple,        |
|                         | [action]. Markdown, | Lacks rationale, structure.        | depth for AI parsing.            | iterative.               |
|                         | JSON.               |                                    |                                  | **Cons**: Too minimal,   |
|                         |                     |                                    |                                  | no diagrams.            |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| UCDD                    | Actors, goal, main  | Similar, but lacks Mermaid,        | High: Narrative structure suits  | **Pros**: Detailed,      |
|                         | scenario, text/JSON.| needs, rationale. Has extensions.  | vibecoding. No diagrams.         | lightweight.             |
|                         |                     |                                    |                                  | **Cons**: Lacks visual   |
|                         |                     |                                    |                                  | aids.                   |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| BDD                     | Given-When-Then,    | Concise, no actors or rationale.   | Moderate: Scenarios useful, but  | **Pros**: Parseable,     |
|                         | Gherkin/JSON.       | Gherkin vs. JSONC. No diagrams.    | Gherkin unfamiliar to vibecoding.| test-friendly.           |
|                         |                     |                                    |                                  | **Cons**: Limited scope, |
|                         |                     |                                    |                                  | no diagrams.            |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| DDD (Event Storming)    | Events, commands,   | Event-based, no flows or rationale.| Low: Too abstract for simple     | **Pros**: Domain-focused,|
|                         | JSON/YAML.          | JSON aligns.                       | vibecoding projects.             | parseable.               |
|                         |                     |                                    |                                  | **Cons**: Lacks flows,   |
|                         |                     |                                    |                                  | diagrams.               |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| CASE Tools              | UML-style,          | Adds repositories, tracing. Lacks  | Low: Heavy for small teams.      | **Pros**: Automated,     |
|                         | XML/JSON, diagrams. | needs, rationale. UML vs. Mermaid. | Tool overhead unsuitable.        | visual.                  |
|                         |                     |                                    |                                  | **Cons**: Complex, tool- |
|                         |                     |                                    |                                  | dependent.               |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+
| USECASES.jsonc          | Actors, flows,      | N/A. Unique in Mermaid, needs,     | Perfect: Tailored for vibecoding,| **Pros**: Lightweight,   |
|                         | rationale, Mermaid, | rationale. JSONC for simplicity.   | supports AI and human needs.      | AI-friendly, visual.     |
|                         | JSONC.              |                                    |                                  | **Cons**: Needs converter|
|                         |                     |                                    |                                  | for docs.               |
+-------------------------+---------------------+------------------------------------+----------------------------------+--------------------------+

Why USECASES.jsonc?
-------------------

``USECASES.jsonc`` is designed for vibecoding’s unique needs:

- **Lightweight**: JSONC is simpler than XML or Gherkin, editable without CASE tools.
- **AI-Friendly**: Structured for AI parsing (e.g., Cursor, LLMs), with rationale fields (needs, choices) for context.
- **Visual**: Mermaid diagrams are lightweight, unlike UML’s tool dependency.
- **Flexible**: Adapts to CLI, web, or scripts, unlike DDD’s domain focus or BDD’s test focus.
- **Balanced**: More detailed than Agile stories, leaner than UML/CASE, inspired by UCDD.

Use it for rapid prototyping, AI-driven development, and small-team collaboration.