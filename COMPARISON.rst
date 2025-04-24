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

Framework Comparison Table
==========================

.. list-table::
   :header-rows: 1
   :widths: 20 20 30 25 25

   * - Framework
     - Key Features
     - Differences from USECASES.jsonc
     - Fit for Vibecoding
     - Pros/Cons
   * - UML/UP
     - | ID, actors, flows,
       | preconditions, UML
       | diagrams. Text/XML.
     - | Adds alternative flows,
       | extensions, UML tools. Lacks
       | needs, rationale, Mermaid.
     - | Moderate: Detailed but heavy
       | for small teams. Tool dependency
       | clashes with rapid iteration.
     - | **Pros**: Comprehensive, visual.
       | **Cons**: Complex, tool-dependent.
   * - Agile (User Stories)
     - | As a [role], I want
       | [action]. Markdown,
       | JSON.
     - | Simpler, no flows or diagrams.
       | Lacks rationale, structure.
     - | High: Lightweight, but lacks
       | depth for AI parsing.
     - | **Pros**: Simple, iterative.
       | **Cons**: Too minimal, no diagrams.
   * - UCDD
     - | Actors, goal, main
       | scenario, text/JSON.
     - | Similar, but lacks Mermaid,
       | needs, rationale. Has extensions.
     - | High: Narrative structure suits
       | vibecoding. No diagrams.
     - | **Pros**: Detailed, lightweight.
       | **Cons**: Lacks visual aids.
   * - BDD
     - | Given-When-Then,
       | Gherkin/JSON.
     - | Concise, no actors or rationale.
       | Gherkin vs. JSONC. No diagrams.
     - | Moderate: Scenarios useful, but
       | Gherkin unfamiliar to vibecoding.
     - | **Pros**: Parseable, test-friendly.
       | **Cons**: Limited scope, no diagrams.
   * - DDD (Event Storming)
     - | Events, commands,
       | JSON/YAML.
     - | Event-based, no flows or rationale.
       | JSON aligns.
     - | Low: Too abstract for simple
       | vibecoding projects.
     - | **Pros**: Domain-focused, parseable.
       | **Cons**: Lacks flows, diagrams.
   * - CASE Tools
     - | UML-style,
       | XML/JSON, diagrams.
     - | Adds repositories, tracing. Lacks
       | needs, rationale. UML vs. Mermaid.
     - | Low: Heavy for small teams.
       | Tool overhead unsuitable.
     - | **Pros**: Automated, visual.
       | **Cons**: Complex, tool-dependent.
   * - USECASES.jsonc
     - | Actors, flows,
       | rationale, Mermaid,
       | JSONC.
     - | N/A. Unique in Mermaid, needs,
       | rationale. JSONC for simplicity.
     - | Perfect: Tailored for vibecoding,
       | supports AI and human needs.
     - | **Pros**: Lightweight, AI-friendly, visual.
       | **Cons**: Needs converter for docs.

Why USECASES.jsonc?
-------------------

``USECASES.jsonc`` is designed for vibecoding’s unique needs:

- **Lightweight**: JSONC is simpler than XML or Gherkin, editable without CASE tools.
- **AI-Friendly**: Structured for AI parsing (e.g., Cursor, LLMs), with rationale fields (needs, choices) for context.
- **Visual**: Mermaid diagrams are lightweight, unlike UML’s tool dependency.
- **Flexible**: Adapts to CLI, web, or scripts, unlike DDD’s domain focus or BDD’s test focus.
- **Balanced**: More detailed than Agile stories, leaner than UML/CASE, inspired by UCDD.

Use it for rapid prototyping, AI-driven development, and small-team collaboration.
