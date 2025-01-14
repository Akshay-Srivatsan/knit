# Knit 🧶

![Test Workflow](https://github.com/zyedidia/knit/actions/workflows/test.yaml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/zyedidia/knit.svg)](https://pkg.go.dev/github.com/zyedidia/knit)
[![Go Report Card](https://goreportcard.com/badge/github.com/zyedidia/knit)](https://goreportcard.com/report/github.com/zyedidia/knit)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zyedidia/knit/blob/master/LICENSE)

Knit is a build tool inspired by Make and Plan9 Mk. You define rules with a
Make-like embedded syntax within a Lua program. Rules can be passed around as
Lua objects, and generated by Lua code. You can use the Lua module system to
make reusable modules for building any kind of source code. Knit combines the
readability of a Make-style rules language will the power and expressiveness of
Lua. If you are familiar with Make, you can learn Knit very quickly.

Knit tracks more of your build to give you better incremental builds. For
example, Knit automatically adds an implicit dependency on a rule's recipe, so
if you change a recipe (either directly or through a variable change), Knit
will automatically re-run all rules that were affected.

Knit has support for namespaced sub-builds that execute relative to their
directory, but Knit avoids build fragmentation because sub-builds don't rely on
spawning build sub-processes. No more `make -C` to do sub-builds! Everything is
tracked by the root Knitfile, but you can still make directory-specific rules.

Knit's rules language is heavily inspired by [Plan9
Mk](https://9p.io/sys/doc/mk.html). In some ways, Knit can be considered a
modern version of Mk with a Lua meta-programming system built on top of it
(there are some differences compared to Mk).

Why make yet another build system? Because it's fun and useful to me! Maybe it
will be useful to you too. Everyone hates something about their build system so
if you have feedback or a request, let me know! The project is new enough that
your feedback may be seriously taken into account.

I have written an article with more details about Knit
[here](https://zyedidia.github.io/blog/posts/3-knit-better-make/).

# Features

* Knit uses Lua for customization. This makes it possible to write reusable
  build libraries, and in general makes it easier to write powerful and
  expressive builds.
* Knit has built-in syntax for a rules language inspired by Make and Plan9 Mk.
  This makes it very familiar to anyone who has used Make/Mk.
* Knit has direct support for sub-builds (compared to Make, which usually
  involves spawning a separate make sub-process to perform a sub-build).
* Knit can hash files to determine if they are out-of-date, rather than just
  relying on file modification times.
    * Knit additionally uses hashes for "dynamic task elision": if Knit can
      dynamically determine that a prerequisite that was rebuilt actually
      changed nothing, it won't re-run the dependent build step, allowing for
      even better incremental builds compared to timestamp-based approaches
      (Make, Ninja, etc.).
* Knit tracks recipe changes, so if you update a variable (in the Knitfile or
  at the command-line), any dependent rules will be automatically rebuilt.
* Knit supports `%` meta-rules and regular expression meta-rules.
* Knit uses rule attributes instead of using targets such as `.SECONDARY` to
  indicate special processing.
* Knit supports virtual rules that are independent of the file system.
* Knit uses sane variable names like `$input`, `$output`, and `$match` instead
  of Make's `$^`, `$@`, and `$*`.
* Knit supports rules with multiple outputs, and treats them like Make's group
  targets by default.
* Knit supports sub-tools that implement various build utilities including:
    * Generating a graph visualization using graphviz (dot).
    * Showing build status information (whether targets are out-of-date and
      why).
    * Exporting a compile commands database for use with a language server.
    * Automatically cleaning all build outputs.
    * Converting your build into a shell script, Makefile, or Ninja file.
* Knit will search up the directory hierarchy for a Knitfile, allowing you
  to run your build from anywhere in your project.
* Knit supports parallel builds and uses all cores by default.
* Cross-platform support (Windows support is still experimental).
    * Knit uses a shell to execute commands. By default, Knit searches for `sh`
      on your system and uses that. If it cannot find `sh`, it uses an internal
      (cross-platform) shell.

# Example Knitfile

Here is a very basic Knitfile for building a simple C project.

```lua
return b{
    $ hello: hello.o
        cc -O2 $input -o $output
    $ %.o: %.c
        cc -O2 -c $input -o $output
}
```

The syntax for rules is nearly the same as Make, and Knit supports `%`
meta-rules just like Make. However, rather than using a custom language
to configure the build, Knit uses Lua.

Here is a more complex example Knitfile used for building a simple C project.
This time the Knitfile supports various configurations (changing `cc` and
enabling debug flags), and automatically detects the source files.

```lua
local knit = require("knit")

local conf = {
    cc = cli.cc or "gcc",
    debug = tobool(cli.debug) or false,
}

local cflags := -Wall

if conf.debug then
    cflags := $cflags -Og -g
else
    cflags := $cflags -O2
end

local src = knit.glob("*.c")
local obj = knit.extrepl(src, ".c", ".o")
local prog := hello

return b{
    $ build:VB: $prog

    $ $prog: $obj
        $(conf.cc) $cflags $input -o $output
    $ %.o:D[%.d]: %.c
        $(conf.cc) $cflags -MMD -c $input -o $output
}
```

Running `knit hello` would build all the necessary `.o` files and then link
them together. Running `knit hello debug=1` would change the flags and re-run
the affected rules. Running `knit build` will build `hello` (effectively an
alias for `knit hello`).  The `VB` attributes on the build rule means that it
is virtual (not referring to a file on the system), and should always be built
(out-of-date).

Running `knit -t clean` will run a sub-tool that automatically removes all
generated files.

Header dependencies are automatically handled by using the `-MMD` compiler flag
with the `D[%.d]` attribute. To explicitly name the dependency file (e.g., to
put it in a `.dep` folder), you could instead use:

```
$ %.o:D[.dep/%.dep]: %.c
    $(conf.cc) $cflags -MMD -MF $dep -c $input -o $output
```

Note that Knitfiles are Lua programs with some modified syntax: special syntax
using `$` for defining rules, and special syntax using `:=` for defining raw
strings (no quotes) with interpolation.

See the [docs](./docs/knit.md) for more information.

See [examples](./examples) for a few examples, and see this repository's
Knitfile and the tests for even more examples.

# Installation

Prebuilt binaries are available from the [release page](https://github.com/zyedidia/knit/releases).

You can install one automatically using [eget](https://github.com/zyedidia/eget).

```
eget zyedidia/knit
```

Or you can build from source (requires Go 1.19):

```
go install github.com/zyedidia/knit/cmd/knit@latest
```

# Experimental or future possible features

* Ninja to Knit converter (for compatibility with cmake, and for benchmarking).
  See [knitja](https://github.com/zyedidia/knitja) for the converter tool.
* Performance optimizations.
    * Knit can already be used to build large projects, such as CVC5 (using the
      knitja converter). For massive builds though, like LLVM, Knit suffers
      from some performance problems that could be improved.
* Better support for dynamic dependencies. Currently it is possible to handle
  dynamic dependencies by generating rules, but I would like to explore the
  possibility of a more clean and cohesive solution.
    * Ptrace enabled automatic dependency discovery (Linux-only feature).
      See the [xkvt](https://github.com/zyedidia/xkvt) project for some
      experiments on this front.
* Global build file cache (similar to `ccache`, but for every command that is
  executed).
* A restrictive mode for build sandboxing.

# Feedback

It is always useful useful to get feedback from others to improve Knit. If you
have feedback, or questions about how to use it, please open a discussion. It
would be great to discuss the good and bad parts of the current design, and how
it can be improved.

# Usage

```
Usage of knit:
  knit [TARGETS] [ARGS]

Options:
  -B, --always-build        unconditionally build all targets
      --cache string        directory for caching internal build information (default ".")
      --cpuprofile string   write cpu profile to 'file'
  -D, --debug               print debug information
  -C, --directory string    run command from directory
  -n, --dry-run             print commands without actually executing
  -f, --file string         knitfile to use (default "knitfile")
      --hash                hash files to determine if they are out-of-date (default true)
  -h, --help                show this help message
      --keep-going          keep going even if recipes fail
  -q, --quiet               don't print commands
      --shell string        shell to use when executing commands (default "sh")
  -s, --style string        printer style to use (basic, steps, progress) (default "basic")
  -j, --threads int         number of cores to use (default 8)
  -t, --tool string         subtool to invoke (use '-t list' to list subtools); further flags are passed to the subtool
  -u, --updated strings     treat files as updated
  -v, --version             show version information
```

Available sub-tools (`knit -t list`):

```
list - list all available tools
graph - print build graph in specified format: text, tree, dot, pdf
clean - remove all files produced by the build
targets - list all targets (pass 'virtual' for just virtual targets)
compdb - output a compile commands database
commands - output the build commands (formats: knit, json, make, ninja, shell)
status - output dependency status information
```

# Contributing

If you find a bug or have a feature request please open an issue for
discussion. I am sometimes prone to being unresponsive to pull requests, so I
apologize in advance. Please ping me if I forget to respond. If you have a
feature you would like to implement, please double check with me about the
feature before investing lots of time into implementing it.

If you have a question or feedback about the current design, please open a
discussion.
