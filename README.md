# meta-lang-go
Go code patterns for AI. Stores reusable patterns (e.g., cmd-handler, msg-handler). Each pattern includes <pattern>.json, code.go, tests.go. AI uses meta-lang-hub/index-all.json to locate patterns, replaces code with indexes ([go-cmd-handler]) for ≥2 uses, adds comments (// [index]) for single uses.
