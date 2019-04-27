# English Lang

Yet another programming language

Experimental & in very [pre-design](./ast/README.md) status.

## Status

* [ ] CI setup
* [ ] English AST Paser **In Progress**
* [ ] The first `Hello World` program **In Progress**
* [ ] The first `Loop` program
* [ ] English Command Line Tool
* [ ] Core `article` - Standard Input/Output
* [ ] Core `article` - File System
* [ ] Core `article` - Go Routine

## Description

This target of this project to running english articles/sentences on computer, and output meaningful result.

This project can **NOT** reply your questions like an `AI` or `alpha go`, it just a **programming language**, and all variables and data processing are controlled by writers.  

## Interpreted Language

* Plan to parse & interprete the `english source code` on `Golang` VM.
* If the progressing of `english ast parser` sub-project well, could compiler the `english source code` to `LLVM IR`, then compiler to binary.

## Internal

* Use [prose](https://github.com/jdkato/prose) nlp tool to tokenizing word, segmente sentences,and tag `part-of-speech`.
* Parse `english senteces` to AST by `part-of-speech` senquence.
* Interprete the AST on `golang VM`
