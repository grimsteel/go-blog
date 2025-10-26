A tutorial for programming in Go for C programmers: build a website in Go without any prior Go experience

This blog article assumes a good level of familiarity with general programming concepts as well as web development architecture. (HTTP, HTML, etc.)

### Why use Go?

Go is a performant and memory safe modern programming language. It contains native support for lightweight threads, called "Goroutines", making it especially suited for web services (Learn).

The wealth of libraries available within Go itself mean external libraries do not need to be used to create performant and scalable websites, unlike many other languages. (Go Web Programming)

Furthermore, Go is already being adopted by numerous big-tech companies with an interest in web development despite its comparatively young age (Go Web Programming)

In addition, the static and strongly typed system will appear familiar to those experienced with C-like languages, making it an easy language to get started with for experienced programmers (Will).

### Part 1: An Introduction to Go

#### Getting Started

You’ll need to install the Go compiler first, available at [https://go.dev](https://go.dev) or in your favorite package manager. As with other languages, I’d recommend creating a new folder for each project you create.

Go also has a free online interactive playground: [https://go.dev/play/](https://go.dev/play/). This is useful for testing and sharing code snippets

#### Hello, World

Once you have the Go compiler installed, make a new folder for your project and run `go mod init` inside it. 

Enter the following code in a file called `hello.go`, and run `go run hello.go` to compile and run it.

```go
import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

How does this code work? Let’s break it down:

- It starts with an import statement: this imports the “fmt” package, which is effectively equivalent to “stdio” in C.  
- It has a main function, which contains a print statement.

Compare this to a Hello, world C program:

```c
#include <stdio>

int main() {
	printf("Hello, world!\n");
	return 0;
}
```

Apart from differences in the signature of the main method, they look very similar\!

#### Go Syntax Introduction

Go’s syntax is relatively simple. The following section is a quick introduction to Go’s syntax. I’ve pointed out some surprising or unique attributes of it where applicable. 

For a more thorough introduction to syntax, check out the excellent GitHub Gist made by Bryan Will, or the official \[Tour of Go\]

Like Rust, semicolon line terminators are *optional* in Go, in most cases. Leave them out unless your program doesn’t compile\!

Go uses C-style comments. (// and /\* \*/)

C Type Comparison: (Will)

Many integer types within Go are almost the same as the corresponding ones in stdint. 

| C type | stdint type | Go type |
| :---- | :---- | :---- |
| char | int8\_t | int8 |
| unsigned char | uint8\_t | uint8 *or* byte |
| short | int16\_t | int16 |
| int | int32\_t | int32 *or* rune\* |
| long | int64\_t | int64 |
|  | size\_t | int |
|  | uintptr\_t | uintptr |
| float |  | float32 |
| double |  | float64 |
| bool |  | bool |
| char\* |  | string |

\*What is a rune? A rune is a single Unicode codepoint. These become relevant when finding the length of a string with non-ASCII characters, but we won’t be using runes in this guide.

Like other modern languages, Go does have a built-in string type.

Go’s operators are effectively the same as C. It includes the standard arithmetic, bitwise, and relational operators you’re likely familiar with.

#### Variable Declaration

Variables can either be declared with the var keyword or the walrus operator. Multiple variables — even those of *different types —* can be declared and assigned in a single statement. (Will)

Go infers types from the right-hand side of the expression, but an explicit type annotation can be added to the declaration statement. In this case, all variables in the declaration must have the same type. (Will)

```go
// the two following statements are equivalent
var a, b = "foo", 42
// the walrus operator will declare any variables not declared already
a, b := "foo", 42

// assignment
a, b = "bar", 41

// typed declaration
var c, d int8 = 6, 7
```

Unlike other languages, it is a **compile-time** error in Go to declare a variable and not use it. This threw me off at first\! I’ll often declare variables but not actually write code that uses them until well into the project. Enforcing usage is definitely an interesting choice.

#### Arrays

Arrays have fixed size (known at compile-time) and known type. Go actually allows you to compare two arrays by *value* using the \== operator, as long as they have the same size and type\! (Will)

The array syntax may look backwards from a C-perspective (just wait until we get to pointers\!)

Array literals are also possible with curly braces, like C.

```go
var foo [67]float64
// explicitly specify variable type
var bar [3]int = [3]int {1, 2, 3}
// infer size and variable type
var foobar = [...]string { "a", "b", "c" }
```

#### Casting

Go has no implicit type coercion, even between integer types and float types. To cast a variable to a type, call the type like a function. (Will)

```go
var a = int8(123123123)
```

#### Functions

Functions are declared with the func keyword. Curly braces are required, and the return type appears within parentheses after the parameters. (Learn)

Multiple return values are possible, but they must be destructured by the caller. (The syntax here is similar to that of JavaScript or Rust)

```go
func foo(a bool, b bool) (bool) {
	return a && b
}
// multiple return values
func bar() (bool, bool) {
	return true, false
}
func foobar() (bool) {
	// destructure the two return values
	a, b := bar()
	return foo(a, b)
}
```

As you may have expected, your program’s entrypoint function should be called `main`

#### Packages and Imports

Go programs are structured into packages which can be imported by other packages. Go enforces directory structure with packages (Learn): all source files in a given directory must be a part of the same package. (By convention, the directory name and the package name should match, but this is *not* enforced (Will)).

What’s interesting about Go is that import statements for non-standard library packages are “absolute” and specify the exact source of a package, like GitHub. This is a stark difference from other languages which typically have a standard package repository and list dependencies in a standalone file, like package.json or Cargo.toml

```go
package main

import "fmt"

import "github.com/user/package"
```

If you are making a command-line binary, your package must be called “main”. However, one of the first roadblocks I encountered with this involved naming the module “main” in go.mod. Don’t do this\! Name it something else instead This can cause issues later on with testing. (Berger)

One common package in the standard library is “fmt”. This package allows you to print, among other things.

The fmt library contains fmt.Println, fmt.Print, and fmt.Printf. These do the same things as their counterparts in languages like Java.

I won’t cover making your own libraries in this guide, but the video “Learn GO Fast: Full Tutorial” covers it well if you’re interested.

#### Structs and Pointers

Public members accessible outside of the struct **must begin with an uppercase character**. This is unlike other languages which use a visibility keyword for this purpose.

Conceptually, structs in Go are identical to C structs. A struct is defined with the `struct` keyword. It can be named as a type by prepending the definition with `type TypeName`. (Will) (This actually applies to any custom type in Go \- think of it like a `typedef` from C++).

```go
type BlogArticle struct {
	Url string
	Length int
	Author string
}

// structs can later be modified using dot syntax
var article BlogArticle
article.Url = "https://example.com"
```

Pointer types use the same \* and & syntax from C. This is especially useful for structs which have a non-negligible cost associated with copying (Will)

```go
var a int = 3
var b int* = &a

articlePtr := &article
```

#### Loops

Go has only one loop construct: the for loop. It can be used as a traditional c-style for loop, a while loop, for a for-each loop over an array.  (Learn)

The syntax for the C-style loop and while loop should look familiar \- it’s the exact same as C or Java, except without the parentheses.

Like C, parts of the for loop can be omitted. Omitting everything as well as the semicolons yields an infinite loop.

```go
// C-style
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// while
j := 50
for j >= 0 {
	fmt.Println(j)
	j--
}

// infinite loop (while true)
for {
	fmt.Println("yes")
}
```

#### Fizz Buzz

Putting it all together: Here’s how you would write Fizz Buzz in Go. You can definitely write it in a much simpler way, but I’ve chosen to use as many different Go features as possible.

```go
import "fmt"

func fizzBuzz(end uint) {
	for i := range(end) {
		fizz, buzz := i % 3 == 0, i % 15 == 0
		switch {
			case fizz && buzz:
				fmt.Println("FizzBuzz");
			case fizz:
				fmt.Println("Fizz");
			case buzz:				fmt.Println("Buzz");
			default:
				fmt.Println(i);
}
}
}
```

### Part 2: Making a Website in Go

In this section, I’ll introduce common libraries and packages used when creating web servers in Go. We’ll build a simple blog website capable of rendering rich text content and handling user comments. In fact, this is how the very article you are currently reading was published\! (not yet)

The content for the articles will be contained in Markdown files, and metadata will be contained in a single JSON file. The aim of this short project is to demonstrate a wide range of Go’s capabilities.

#### Starting a Web Server

We’re going to be using the `net/http` library. To start a web server, create an instance of the `http.Server` struct and call `ListenAndServe`:

```go
package main

import "net/http"

func main() {
	server := &http.Server{
   	 	Addr:           "0.0.0.0:8080",
    	  	Handler:       /* TODO */,
  	}
  	server.ListenAndServe()
}
```

The `Handler` member should be a function that satisfies `http.HandlerFunction`. 

#### Routing and the Mux

Go has a wealth of libraries available for your use, and there are many such libraries which abstract away a lot of the boilerplate needed for creating a powerful router, but I want to focus on what Go has to offer in this article.

That being said, here are some libraries you might want to look into yourself to make web development easier:

- Chi (JetBrains)
- Go router (book)

Go includes a built-in router of sorts, `http.ServeMux`. This allows you to assign a function to different routes. To use it, call `http.NewServeMux` and pass the instance as your handler function.  
Every website also needs to be able to serve static files. `http.FileServer` can accomplish that: pass an instance of `http.Dir` into it. You’ll have to wrap it in `http.StripPrefix` when mounting it into the mux

```go
package main

import (
  "net/http"
)

func main() {
  staticPath := "/static/"
  mux := http.NewServeMux()
  mux.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(http.Dir("./static")))
  mux.HandleFunc("/", index)

  server := &http.Server{
    Addr:     "0.0.0.0:8080",
    Handler:  mux,
  }
  server.ListenAndServe()
}
```

(code snippet adapted from Go Web Programming)

I’ll be using [water.css](https://watercss.kognise.dev/) to add some basic styling to the website, so go ahead and download that into the static folder. This is what the directory structure looks like right now:

```
project
- main.go
- go.mod
- static
  - water.css
- templates
```

#### Templates

Go has built in support for HTML templating. The syntax is reminiscent of Jinja from Python. 

Let’s start by adding a post listing page. Create a folder called `templates`, and add a file called `post.html` into this folder. 

#### Template Inheritance

Components uising templates (Components...)

Modify base

#### Aside: Components using templates\!** (Component

#### Adding Libraries

Finally, let’s finish the most important part of the website: the blog article page itself\! We’ll need to use a Markdown renderer, like `gomarkdown/markdown`. Like I mentioned earlier, adding a library to a Go project doesn’t involve adding it to a single dependency file. Instead, just import it in any source file and run `go mod sum` to regenerate the lock file.

```go
import "github.com/gomarkdown/markdown"

func renderPost() {
	
}
```

#### Form Responses

### References

1. Bhattacharyea, Aniket “Build a Blog with Go Templates.” *The GoLand Blog,* 8 November 2022, [blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/](http://blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/). JetBrains.

2. Chang, Sau Sheong. *Go Web Programming*. Manning Publications, 2016, *Go (O’Reilly Learning),* [learning.oreilly.com](https://learning.oreilly.com/library/view/go-web-programming/9781617292569/)

3. *Components with HTML Templates in Go\!? \~ FULL STACK  Golang.* samvcodes, 3 Dec 2023, *YouTube*. Web. 6 Oct 2025, [www.youtube.com/watch?v=nX1syyFnUS4](http://www.youtube.com/watch?v=nX1syyFnUS4)

4. Davidson, Kristin “How To Use Dates and Times in Go” How to Code in Go, 2 February 2022, digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go. DigitalOcean

5. *Learn GO Fast: Full Tutorial*. Alex Mux. 4 Sep 2023\. *YouTube.* Web. 6 Oct 2025, [www.youtube.com/watch?v=8uiZC0l4Ajw](http://www.youtube.com/watch?v=8uiZC0l4Ajw).

6. Will, Bryan. “Go language overview for experienced programmers.” *GitHub Gist*, Oct. 2016, [gist.github.com/BrianWill/5be5a96c86cb4f7ce24034b50cde24fd](http://gist.github.com/BrianWill/5be5a96c86cb4f7ce24034b50cde24fd).


