# Educational-Password-Generator
Educational password generator that teaches core cybersecurity concepts while you build and test strong passwords.

# Architecture Overview
## High-Level Design
The project is split into three main pieces:

1. Frontend (Web UI)
* Built with React

2. Backend (Go API)
* REST API written in Go that exposes endpoints for password generation and analysis
* Calculates entropy and approximate crack times using configurable attack models (e.g. online guess rate vs offline hash cracking).

3. Content Layer (Educational Text)
* A structured set of short explanations and tips, stored in JSON/Markdown files.
* The app selects which pieces to show based on the user’s password settings or analysis results.
