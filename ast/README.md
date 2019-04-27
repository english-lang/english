# English AST Tool

## Programming Diagram

* OOP
* FP

## Design

### Process flow

* Read english article source
* Split english paragraphes (each para will have it's own scope/context)
* Segmente english sentence
* Tokenize & tagging english words
* Parse pre-processed AST tree with English Grammer
* Manually detect english **text intent** (by switch code, maybe AI can be involved)
* Mapping pre-processed tree to after-processed AST tree by **text intent** (General programming AST tree)
* Use `golang` intercepte or use `LLVM` compile to binary.

### Word Type Usage

#### Noun

A word (other than a pronoun) used to identify any of a class of people, places, or things ( common noun ) , or to name a particular one of these ( proper noun ).

In `English Lang`, it represents an `identifier`.

#### Verb

A word used to describe an action, state, or occurrence, and forming the main part of the predicate of a sentence, such as hear , become , happen.

In `English Lang`, it represents a `method` of an obejct. 

#### Adverb

A word or phrase that modifies or qualifies an adjective, verb, or other adverb or a word group, expressing a relation of place, time, circumstance, manner, cause, degree, etc. (e.g., gently , quite , then , there ).

In `English Lang`, it represents a modifier of `verb`/`adjective`, and `verb`/`adjective` can get modifiers in its internal.

### Programming Concept

#### Article

A general programming `module`, it can define objects/variables.

#### Paragraph

A general programming `block`, each paragraph end with `\n` or `\r\n`, and variables will be shared in this scope.

#### Sentence

A valid english sentench with english grammer, describe operations.
