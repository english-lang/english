# English Lang

Yet another programming language

Experimental & in very [pre-design](./ast/README.md) status.

## **Pause**

The accuracy of lib `prose` is not enough, look up new lib or model.

`Part-of-speech` tagging is the core of this project.

## Status

* [ ] CI setup
* [ ] English AST Parser **In Progress**
* [ ] The first `Hello World` program **In Progress**
* [ ] The first `Loop` program
* [ ] English Command Line Tool
* [ ] Core `article` - Standard Input/Output
* [ ] Core `article` - File System
* [ ] Core `article` - Go Routine

## Description

This target of this project to running english articles/sentences on computer, and output meaningful result.

This project can **NOT** reply your questions like an `AI` or `alpha go`, it just a **programming language**, and all variables and data processing are controlled by writers.  

## Runtime

* Plan to parse & interpret the `english source code` on `Golang` HLVM (`High Level VM`), and its easy to impl the native api bindings (`golang` have impl that).

* Another option is compile the `english source code` to `LLVM IR`, then compiler to binary, but the problem is hard to build the `std library`.

## Internal

* Use [prose](https://github.com/jdkato/prose) nlp tool to tokenizing word, segmente sentences,and tag `part-of-speech`.
* Parse `english sentences` to AST by `part-of-speech` sequence.
* Interpret the AST on `golang VM`

## Problems

* **Hard** to define string literal in english language.
