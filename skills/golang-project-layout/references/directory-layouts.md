# Directory Layouts

## Universal Layout (Most Projects)

```
project/
├── cmd/                    # Entry points - ONE subdirectory per main package
│   ├── server/            # Main application #1
│   │   └── main.go
│   ├── client/            # Main application #2
│   │   └── main.go
│   └── migrate/           # Main application #3
│       └── main.go
│   └── cli/               # Main application #4
│       └── main.go
│   └── worker/            # Main application #5
│       └── main.go
├── internal/              # Private application code (`internal/` MUST be used for non-exported packages)
│   ├── app/              # Application initialization
│   ├── config/           # Configuration loading
│   ├── handler/          # HTTP/request handlers
│   ├── model/            # Data models/domain
│   └── service/          # Business logic
├── logger/                # Public package (optional - only if useful to other projects)
│   └── logger.go
├── api/                   # API definitions (optional)
│   └── openapi.yaml
├── configs/               # Configuration files (optional)
│   └── config.yaml
├── scripts/               # Build/deployment scripts (optional)
├── build/                 # goyek build automation
│   ├── go.mod             # Separate module for build tooling
│   ├── go.sum
│   ├── main.go
│   ├── build.go           # Artifact creation tasks
│   └── check.go           # format/lint/test tasks
├── dist/                  # Generated artifacts (gitignored)
│   ├── bin/               # Built binaries
│   └── test/              # Coverage and other test artifacts
├── build.sh               # Root wrapper: cd build && go run . "$@"
├── go.mod
├── go.sum
├── .gitignore             # Git ignore patterns
├── .golangci.yml          # Linter configuration
├── LICENSE                # License file
└── README.md
```

## Small Projects (Single Binary)

For simple tools, keep it minimal:

```
my-tool/
├── cmd/
│   └── my-tool/
│       └── main.go        # Single main package
├── internal/
│   └── core.go            # Application logic
├── build/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── build.go           # Artifact creation tasks
│   └── check.go           # format/lint/test tasks
├── dist/
│   ├── bin/
│   └── test/
├── build.sh               # Root goyek wrapper
├── go.mod
├── .gitignore             # Git ignore patterns
├── .golangci.yml          # Linter configuration (optional)
├── LICENSE                # License file (recommended)
└── README.md
```

## Libraries (Reusable Code)

```
my-library/
├── example/               # Example
├── logger/                # Public package
│   ├── logger.go
│   └── logger_test.go
├── internal/
│   └── impl/              # Private implementation details
│       └── core.go
├── build/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── build.go           # Artifact creation tasks
│   └── check.go           # format/lint/test tasks
├── dist/
│   ├── bin/
│   └── test/
├── build.sh               # Root goyek wrapper
├── go.mod
├── go.sum
├── .gitignore             # Git ignore patterns
├── .golangci.yml          # Linter configuration
├── LICENSE                # License file
└── README.md
```

**Key points for libraries:**

- Put public API in root-level directories (e.g., `logger/`)
- Use `internal/` for private implementation
- Don't use `cmd/` (unless you have example binaries)

## The cmd/ Directory Convention

**CRITICAL**: All `main` packages must reside in `cmd/`. `cmd/` MUST contain only `main.go` with minimal logic — parse flags, wire dependencies, call `Run()`. NEVER put business logic in `cmd/` — project-specific code belongs in `internal/`; reusable public packages belong in descriptive root-level directories like `logger/`.

### Single Application

```
cmd/
└── myapp/
    └── main.go    // package main
```

### Multiple Applications

When you need multiple binaries (e.g., server, CLI tool, migration utility):

```
cmd/
├── server/
│   └── main.go        // Runs the API server
├── client/
│   └── main.go        // CLI client tool
├── worker/
│   └── main.go        // Background worker
└── migrate/
    └── main.go        // Database migration utility
```

Each `main.go`:

- Declares `package main`
- Has its own `func main()`
- Should be built through the root build wrapper: `./build.sh build`

**Building all binaries:**

```bash
./build.sh build          # Build all main packages into dist/bin/
```

## Common Mistakes to Avoid

### Don't Do This

```
myproject/
├── src/              # Go doesn't use /src (Java pattern)
├── main.go           # Don't put main at root
├── utils/            # Generic package name
├── helpers/          # Generic package name
└── common/           # Generic package name
```

### Do This Instead

```
myproject/
├── cmd/
│   └── myapp/
│       └── main.go   # Main in cmd/
├── internal/
│   ├── util/         # Specific utility names
│   ├── format/       # Or domain-specific names
│   └── service/      # Project-specific code and implementation details
└── logger/           # Optional public package if useful to other projects
```

## Public Package Convention

Do not create a catch-all `pkg/` directory for exportable code. If a package is part of the public API and useful to other projects, put it at the repository root using a specific package name:

```
myproject/
├── logger/
│   ├── logger.go
│   └── logger_test.go
├── retry/
│   ├── retry.go
│   └── retry_test.go
└── internal/
    └── service/
        └── service.go
```

Use `internal/` for code that is only meant for this project, even if it is shared across multiple binaries inside the repository.

## goyek Build Automation

Use goyek for Go build automation. Keep build tasks as Go code in `build/` and expose a root `build.sh` wrapper so developers and CI can run the same commands from the repository root.

```
myproject/
├── build/
│   ├── go.mod             # Separate module for build tooling
│   ├── go.sum
│   ├── main.go            # goyek entry point; creates dist/bin and dist/test
│   ├── build.go           # clean/build/artifact creation tasks
│   └── check.go           # format/lint/test tasks
├── dist/
│   ├── bin/               # Built binaries
│   └── test/              # Coverage and other test artifacts
├── build.sh               # Runs the build module from the repository root
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    └── app/
        └── app.go
```

Common usage:

```bash
./build.sh -h
./build.sh check
./build.sh test
./build.sh lint
./build.sh build            # Builds commands into dist/bin/
./build.sh all
```

Generated binaries go in `dist/bin/`. Generated test artifacts, such as coverage profiles and reports, go in `dist/test/`. The `clean` task removes `dist/`.
