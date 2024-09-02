## packages

1. Definition: A package in Go is a collection of source files in the same directory that are compiled together. It's a way to group related Go source files to form a single unit.

2. Purpose: Packages serve several important purposes:

-   Code organization
-   Code reusability
-   Encapsulation
-   Namespace management

3. Main Package: Every Go program must have a main package with a main() function, which serves as the entry point of the executable program. if you wanna seperete some func or varible in other file in root or other package you should start main if you are in root or other package name in directory that package there is

4. Creating Packages: To create a package, simply create a new directory and add Go files with the package declaration at the top:

```go
package mypackage

// Functions, types, variables here

```

5. Importing Packages: Use the import keyword to use code from other packages:

```go
package main

import (
    "fmt"
    "myproject/pkg/mypackage"
)

func main() {
    fmt.Println(mypackage.SomeFunction())
}

```

6. Exported Names: In Go, a name is exported (visible outside its package) if it begins with a capital letter:

```go
package mypackage

func ExportedFunction() {} // Visible outside the package
func unexportedFunction() {} // Only visible within the package

```

7. Init Function: Each package can have an init() function that's called automatically when the package is initialized:

```go
package mypackage

func init() {
    // Initialization code
}

```

8. Blank Identifier: You can use the blank identifier (\_) to import a package solely for its side effects (like init()):

```go
import (
    _ "myproject/pkg/mypackage"
)

```

9. Dot Imports: You can use a dot (.) to import all exported names directly into the current namespace, but this is generally discouraged:

```go
import . "fmt"

func main() {
    Println("Hello") // Instead of fmt.Println
}
```

10. Aliasing Imports: You can provide an alternative name for an imported package:

```go
import (
    f "fmt"
)

func main() {
    f.Println("Hello")
}

```

11. Package Documentation: Go encourages documenting packages. Comments before the package clause are used as package documentation:

```go
// Package mypackage provides utility functions for...
package mypackage

```

### Why We Use Packages:

-   Organization: Packages help in logically organizing code, making large projects manageable.

-   Reusability: Code in packages can be easily reused across different projects.

-   Encapsulation: Packages provide a level of encapsulation, allowing you to hide implementation details.

-   Namespace Management: Packages help avoid naming conflicts by providing separate namespaces.

-   Compilation Speed: Packages can be compiled separately, which can speed up compilation of large projects.

-   Dependency Management: Packages make it easier to manage and version external dependencies.

-   Testing: Go's package system integrates well with its testing framework, encouraging good testing practices.

-   Documentation: The package system in Go encourages and facilitates good documentation practices.

-   Modularity: Packages promote modular design, making code more maintainable and easier to understand.

-   Collaboration: Packages make it easier for teams to work on different parts of a project independently.

## Third-party packages

Definition: Third-party packages are libraries or modules created by the Go community or other developers, not included in the Go standard library.

Purpose: They provide additional functionality, tools, and abstractions to solve common problems or implement specific features.

Finding Third-Party Packages:

-   pkg.go.dev: Official Go package discovery site
-   GitHub: Many Go packages are hosted on GitHub
-   Awesome Go: A curated list of Go frameworks, libraries, and software

Installing Third-Party Packages: With Go Modules (recommended):

```bash
go get github.com/author/package
```

Importing Third-Party Packages:

```go
package main

import (
    "github.com/author/package"
)

func main() {
    package.SomeFunction()
}
```

Go Modules: Go Modules is the official dependency management system for Go. It allows you to specify and version your dependencies:

```go
module myproject

go 1.16

require (
    github.com/author/package v1.2.3
)
```

Versioning: Go Modules use semantic versioning. You can specify version constraints in your go.mod file:

```go
require (
    github.com/author/package v1.2.3
    github.com/another/lib v2.0.0+incompatible
)

```

Updating Dependencies: To update a package to its latest version:

```bash
go get -u github.com/author/package
```

Security Considerations:

-   Always verify the source and reputation of third-party packages
-   Check for recent updates and active maintenance
-   Be aware of potential vulnerabilities

License Compliance: Ensure that the licenses of third-party packages are compatible with your project's license.

Replacing Dependencies: Go Modules allow you to replace a dependency with a local version or a fork:

```go
replace github.com/author/package => ../my-local-package
```

Cleaning Up Unused Dependencies: To remove unused dependencies:

```bash
go mod tidy
```

### Popular Third-Party Packages: Here are some widely used third-party packages:

-   gorilla/mux: A powerful HTTP router and URL matcher
-   gorm: The fantastic ORM library for Go
-   cobra: A library for creating powerful modern CLI applications
-   logrus: Structured, pluggable logging for Go
