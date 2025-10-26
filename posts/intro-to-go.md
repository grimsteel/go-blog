A tutorial for programming in Go for C programmers: build a website in Go without any prior Go experience

This blog article assumes a good level of familiarity with general programming concepts as well as web development architecture. (HTTP, HTML, etc.)

### Why use Go?

Go is a performant and memory safe modern programming language. It contains native support for lightweight threads, called "Goroutines", making it especially suited for web services[^learn].

The wealth of libraries available within Go itself mean external libraries do not need to be used to create performant and scalable websites, unlike many other languages[^web].

Furthermore, Go is already being adopted by numerous big-tech companies with an interest in web development despite its comparatively young age[^web].

In addition, the static and strongly typed system will appear familiar to those experienced with C-like languages, making it an easy language to get started with for experienced programmers[^gist].

### Part 1: An Introduction to Go

#### Getting Started

You’ll need to install the Go compiler first, available at [https://go.dev](https://go.dev) or in your favorite package manager. As with other languages, I’d recommend creating a new folder for each project you create.

Go also has a free online interactive playground: [https://go.dev/play/](https://go.dev/play/). This is useful for testing and sharing code snippets[^learn].

#### Hello, World

Once you have the Go compiler installed, make a new folder for your project and run `go mod init` inside it. 

Enter the following code in a file called `hello.go`, and run `go run hello.go` to compile and run it[^learn].

```go
import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

How does this code work? Let’s break it down:

- It starts with an import statement: this imports the “fmt” package[^learn], which is effectively equivalent to “stdio” in C.
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

For a more thorough introduction to syntax, check out the excellent [GitHub Gist made by Bryan Will](https://gist.github.com/BrianWill/5be5a96c86cb4f7ce24034b50cde24fd), or the official [Tour of Go](https://go.dev/tour/welcome/1).

Like Rust, semicolon line terminators are *optional* in Go, in most cases. Leave them out unless your program doesn’t compile\! [^gist]

Go uses C-style comments. (// and /\* \*/)

C Type Comparison: [^gist]

Many integer types within Go are almost the same as the corresponding ones in stdint. 

| C type | stdint type | Go type |
| :---- | :---- | :---- |
| char | int8\_t | int8 |
| unsigned char | uint8\_t | uint8 *or* byte |
| short | int16\_t | int16 |
| int | int32\_t | int32 |
| long | int64\_t | int64 |
|  | size\_t | int |
|  | uintptr\_t | uintptr |
| float |  | float32 |
| double |  | float64 |
| bool |  | bool |
| char\* |  | string |

Like other modern languages, Go does have a built-in string type.

Go’s operators are effectively the same as C. It includes the standard arithmetic, bitwise, and relational operators you’re likely familiar with.

#### Variable Declaration

Variables can either be declared with the var keyword or the walrus operator. Multiple variables — even those of *different types —* can be declared and assigned in a single statement[^gist].

Go infers types from the right-hand side of the expression, but an explicit type annotation can be added to the declaration statement. In this case, all variables in the declaration must have the same type[^gist].

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

Unlike other languages, it is a **compile-time** error in Go to declare a variable and not use it. This threw me off at first\! I’ll often declare variables but not actually write code that uses them until well into the project. Enforcing usage is definitely an interesting choice[^learn].

#### Arrays

Arrays have fixed size (known at compile-time) and known type. Go actually allows you to compare two arrays by *value* using the == operator, as long as they have the same size and type[^gist]!

The array syntax may look backwards from a C-perspective: the length is before the type.

Array literals are also possible with curly braces, like C.

```go
var foo [67]float64
// explicitly specify variable type
var bar [3]int = [3]int {1, 2, 3}
// infer size and variable type
var foobar = [...]string { "a", "b", "c" }
```

#### Functions

Functions are declared with the func keyword. Curly braces are required, and the return type appears within parentheses after the parameters[^learn].

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

Go programs are structured into packages which can be imported by other packages. Go enforces directory structure with packages[^learn]: all source files in a given directory must be a part of the same package. (By convention, the directory name and the package name should match for libraries, but this is *not* enforced![^gist]).

What’s interesting about Go is that import statements for non-standard library packages are “absolute” and specify the exact source of a package, like GitHub. This is a stark difference from other languages which typically have a standard package repository and list dependencies in a standalone file, like package.json or Cargo.toml.

```go
package main

import "fmt"

import "github.com/user/package"
```

If you are making a command-line binary, your package must be called “main”. However, one of the first roadblocks I encountered with this involved naming the _module_ “main” in go.mod. Don’t do this\! Name it something else instead This can cause issues later on with testing[^learn].

One common package in the standard library is “fmt”. This package allows you to print, among other things.

The fmt library contains fmt.Println, fmt.Print, and fmt.Printf. These do the same things as their counterparts in languages like Java[^learn].

I won’t cover making your own libraries in this guide, but the video [“Learn GO Fast: Full Tutorial”](www.youtube.com/watch?v=8uiZC0l4Ajw) covers the concepts well if you’re interested.

#### Structs and Pointers

Public members accessible outside of the struct **must begin with an uppercase character**. This is unlike other languages which use a visibility keyword for this purpose[^gist].

Conceptually, structs in Go are identical to C structs. A struct is defined with the `struct` keyword. It can be named as a type by prepending the definition with `type TypeName`[^gist]. (This actually applies to any custom type in Go \- think of it like a `typedef` from C++).

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

Pointer types use the same \* and & syntax from C. This is especially useful for structs which have a non-negligible cost associated with copying[^gist].

```go
var a int = 3
var b int* = &a

articlePtr := &article
```

#### Loops

Go has only one loop construct: the for loop. It can be used as a traditional c-style for loop, a while loop, for a for-each loop over an array[^learn].

The syntax for the C-style loop and while loop should look familiar \- it’s the exact same as C or Java, except without the parentheses.

Like C, parts of the for loop can be omitted. Omitting everything as well as the semicolons yields an infinite loop[^gist].

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


#### Conditionals

`if` statements do not have parentheses, but the curly braces around the block are required, unlike C-like languages.

```go
if a == 3 {
    
} else if a == 4 {

}
```

`switch` statements without an argument essentially act like long if-else chains. The Fizz Buzz example below has one of these.

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
      case buzz:
        fmt.Println("Buzz");
      default:
	    fmt.Println(i);
    }
  }
}
```

We use the walrus operator to evaluate whether a specific number is fizz, buzz, or both. We then use an argument-less switch.

### Part 2: Making a Website in Go

In this section, I’ll introduce common libraries and packages used when creating web servers in Go. We’ll build a simple blog website capable of rendering rich text content and handling user comments. In fact, this is how the very article you are currently reading was published!

The content for the articles will be contained in Markdown files, and metadata will be contained in a single JSON file. The aim of this short project is to demonstrate a wide range of Go’s capabilities.

#### Project Setup

Make a folder for your project and initialize a Go module. Remember to call the _module_ something other than `main`.

```sh
$ mkdir go-blog && cd go-blog
$ go mod init go-blog
```

#### Error Checking

Many library methods in Go don't actually "panic", or throw an error when they fail. Instead, they just return an `error` object. For our purposes, we want to panic if something happens[^jetbrains], so having a convenience function for this is nice.

Create a file called `util.go` and add this convenience method to panic if there is a non-nil error:

```go
package main

func check(e error) {
	if e != nil {
		panic(e)
	}
}
```

#### Starting a Web Server

We’re going to be using the `net/http` library. To start a web server, create an instance of the `http.Server` struct and call `ListenAndServe`[^web]:

The `Handler` field is a function that handles each HTTP request. It takes a `ResponseWriter` and a `Request` object as parameters. We can write to the response with `Fprintf`.

You might have noticed that the function syntax is a little different here: it's called a **closure**, or an anonymous function which can inherit the scope of the parent. The syntax is actually identical except for the omission of the function name[^gist].

```go
package main // the _package_ can be called main

import "net/http"
import "fmt"

func main() {
  server := &http.Server{
    Addr:           "0.0.0.0:8080",
    Handler:       func (w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, world!")
    },
  }
  check(server.ListenAndServe())
}
```

The `Handler` member should be a function that satisfies `http.HandlerFunction`. 

#### Routing and the Mux

Go has a wealth of libraries available for your use, and there are many such libraries which abstract away a lot of the boilerplate needed for creating a powerful router, but I want to focus on what Go has to offer in this article.

That being said, here are some libraries you might want to look into yourself to make web development easier:

- Chi[^jetbrains]: there's a [blog article from JetBrains](blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/) which explains this well
- Go Router[^web]

Go includes a built-in router of sorts, `http.ServeMux`. This allows you to assign a function to different routes. To use it, call `http.NewServeMux` and pass the instance as your handler function.

Every website also needs to be able to serve static files. `http.FileServer` can accomplish that[^jetbrains]: pass an instance of `http.Dir` into it. You’ll have to wrap it in `http.StripPrefix` when mounting it into the mux. This is because the file paths _within_ the static directory do not have a `/static/` prefix.

We'll mount our original handling function under `/` using the `HandleFunc` method.

```go
package main

import (
  "net/http"
)

func main() {
  staticPath := "/static/"
  mux := http.NewServeMux()
  mux.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(http.Dir("./static")))
  mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
  })

  server := &http.Server{
    Addr:     "0.0.0.0:8080",
    Handler:  mux,
  }
  check(server.ListenAndServe())
}
```

I’ll be using [water.css](https://watercss.kognise.dev/) to add some basic styling to the website, so go ahead and download that into the static folder. This is what the directory structure looks like right now:

```
project
- main.go
- go.mod
- static
  - water.css
- templates
```


#### JSON Parsing

We'll store the post metadata in a JSON file: create a file called `posts.json` in the root folder for your project:
```json
[{
    "id": "first-post",
    "date": "2025-10-20",
    "filename": "first-post.md",
    "title": "Example Post"
}]
```

Go has excellent support for structured JSON parsing[^web]. Create a new file called `post.go` and setup a `Post` struct:

```go
package main

type Post struct {
	Id string
	Date string
	Filename string
	Title string
}
```

We'll make a function that gets all of the posts in `posts.json`.

```
import (
    "encoding/json"
    "os"
)

func getPostList() ([]Post) {
    postListJson, err := os.ReadFile("posts.json")
    check(err)
    
    // this variable will store the posts array
    var posts []Post
    // parse JSON
    check(json.Unmarshal(postListJson, &posts))

    return posts
}
```

We use the `os.ReadFile` method to read the given file, and we use `json.Unmarshal`[^web] to parse the JSON. The check method will cause the program to panic if parsing fails.

#### Templates

Go has built in support for HTML templating. The syntax is reminiscent of Jinja from Python. 

Let’s start by creating the homepage. Create a folder called `templates`, and add a file called `index.html` into this folder. 

```html
<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Blog</title>
    <link rel="stylesheet" href="/static/water.css" />
  </head>
  <body>
    <h1>Example Blog Blog</h1>
    <hr />

    {{ range $val := . }}
    <div>
      <h2>{{ $val.Title }}</h2>
      <a href="/posts/{{ $val.Id }}">View post</a>
      <p>Published on {{ $val.Date }}</p>
    </div>
    {{ end }}
  </body>
</html>
```

We use `range` to iterate over the posts [^jetbrains]. `$val` will refer to the current post, so we can access all of its public fields using dot notation. 

> What does `$val := .` mean?
>
> This is the template syntax for iterating over an array. In this case, `.` _is_ the posts array, because it is the "root data" passed to the template, as you'll see below

Next, we have to serve this template. Create another function in `util.go` which renders a given template:

```go
import (
	"fmt"
	"html/template"
	"net/http"
)

// jetbrains
func renderTemplate(data any, templateFile string, w http.ResponseWriter) {
	t, err := template.ParseFiles(
		// page template
		fmt.Sprintf("templates/%s.html", templateFile),
	)
	check(err)
	check(t.Execute(w, data))
}
```

We use `Sprintf` to create the template filename string from just the "template id", like `index` for `templates/index.html` for use in `ParseFiles`[^web]. As usual, we use `check` for error checking. The data passed to the template is contained in the `data` parameter.

Finally, tie this into the mux:

```go
// read posts
posts := getPostList()
    
mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    renderTemplate(&posts, "index", w)
})
```

#### Date Handling

Currently, the dates being displayed don't look very nice. Let's format them more casually!

Add a format function to `post.go`:

```go
import "time"


func formatPostDate(post *Post) {
	parsedDate, err := time.Parse(time.DateOnly, post.Date)
	check(err)

	post.Date = parsedDate.Format("Monday, January _2")
}
```

This formatting system might look a little weird from a C perspective[^time]. Rather than using format specifiers like "MM" and "DD", go uses a specific "1-2-3" date: `01/02 03:04:05PM 2006 -0700`[^time].

Therefore, to format the date as `DayOfWeek, Month, Day`, we use the corresponding attributes from Go's format date, resulting in `Monday, January _2`.

If you're confused, don't worry! I was too when I first saw this. I would recommend reading [DigitalOcean's article on datetimes](https://digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go) in Go.

#### Template Inheritance

One of the most essential features in any templating language is the ability to inherit and reuse templates as components[^components].

Create a new `base.html` template:

```html
{{ define "base" }}
<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>{{ template "title" . }} • Blog</title>
    <link rel="stylesheet" href="/static/water.css" />
  </head>
  <body>
    <h1>Example Blog</h1>
    <hr />
    <h2>{{ template "title" . }}</h2>
    
    <!-- pass template data -->
    {{ template "main" . }}
  </body>
</html>
{{ end }}
```

Because we're going to have multiple different sections, we need to explicitly `{{ define }}` this block as `base`. 

When a block has been defined, it can be included using `{{ template }}`. This is how we include `title` and `main`, which will be defined in the individual template files[^components]. This should look familiar for those with Jinja or Liquid experience.

The dot after the included template name is the same as the dot for the post list above: it's the root template data. In this case, it's used to pass all the template data down into the nested blocks.

Let's modify `index.html` to use the base structure:

```html
{{ define "title" }}Home{{ end }}

{{ define "main" }}
{{ range $val := . }}
<div>
  <h2>{{ $val.Title }}</h2>
  <a href="/posts/{{ $val.Id }}">View post</a>
  <p>Published on {{ $val.Date }}</p>
</div>
{{ end }}
{{ end }}
```

The title block is defined with `Home` inside, and the main block is defined with the post listing.

Finally, we need to modify our template rendering function[^web]:

```go
t, err := template.ParseFiles(
    // base template
    "templates/base.html",
    // page template
    fmt.Sprintf("templates/%s.html", templateFile),
)

check(err)
	
// multiple separate templates 
check(t.ExecuteTemplate(w, "base", data))
```

The first change is adding the base template as another argument to `ParseFiles`.

The second is changing the `Execute` call to `ExecuteTemplate` to specify exactly which block we want to render. In this case, the `base` block is our root block.

---

I would also recommend checking out [samvcodes's](http://www.youtube.com/watch?v=nX1syyFnUS4) video on this: they go into much more detail about other uses of templates as well.

#### Adding Libraries

Finally, let’s finish the most important part of the website: the blog article page itself\! We’ll need to use a Markdown renderer, like `gomarkdown/markdown`.

Like I mentioned earlier, adding a library to a Go project doesn’t involve adding it to a single dependency file. Instead, just import it in any source file and run `go mod tidy` to regenerate the lock file.

```go
import "github.com/gomarkdown/markdown"

func renderPost() {
	
}
```

#### Form Responses

### References

[^jetbrains]: Bhattacharyea, Aniket “Build a Blog with Go Templates.” *The GoLand Blog,* 8 November 2022, [blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/](http://blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/). JetBrains.

[^web]: Chang, Sau Sheong. *Go Web Programming*. Manning Publications, 2016, *Go (O’Reilly Learning),* [learning.oreilly.com](https://learning.oreilly.com/library/view/go-web-programming/9781617292569/)

[^components]: *Components with HTML Templates in Go\!? \~ FULL STACK  Golang.* samvcodes, 3 Dec 2023, *YouTube*. Web. 6 Oct 2025, [www.youtube.com/watch?v=nX1syyFnUS4](http://www.youtube.com/watch?v=nX1syyFnUS4)

[^time]: Davidson, Kristin. "How To Use Dates and Times in Go" *How to Code in Go*, 2 February 2022, [digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go](https://digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go). DigitalOcean.

[^learn]: *Learn GO Fast: Full Tutorial*. Alex Mux. 4 Sep 2023\. *YouTube.* Web. 6 Oct 2025, [www.youtube.com/watch?v=8uiZC0l4Ajw](http://www.youtube.com/watch?v=8uiZC0l4Ajw).

[^gist]: Will, Bryan. “Go language overview for experienced programmers.” *GitHub Gist*, Oct. 2016, [gist.github.com/BrianWill/5be5a96c86cb4f7ce24034b50cde24fd](http://gist.github.com/BrianWill/5be5a96c86cb4f7ce24034b50cde24fd).


