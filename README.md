# monkey
Learn `writing an interpreter in go` and `writing a compiler in go`.
Implement a monkey programming language.

## Features
* Monkey is a REPL(Read Eval Print Loop) programming language, e.g. Python.
* The parser we write is a recursive descent parser.
> There are two main strategies when parsing a programming language: top-down parsing or bottom-up parsing.
> “Recursive descent parsing” is a variant of top-down parsing.

## Related concepts
* **lexical analysis, lexing or tokenization:** is the process of converting a sequence of characters into a sequence of lexical tokens.
* **parser:** A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree(AST) or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process. The parser is often preceded by a separate lexical analyser, which creates tokens from the sequence of input characters.