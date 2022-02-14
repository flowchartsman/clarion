# Clarion - A Monochrome Theme Inspired By 🧑‍🔬
Clarion is a mostly-monochromatic, minimally-highlighted colorscheme, clearing
away the rainbow madness and allowing you to concentrate on what matters the
most: your code.

![Clarion Default Preview](img/clarion-peach.jpg?raw=true)
![Clarion Orange Preview](img/clarion-orange.jpg?raw=true)
![Clarion Yellow Preview](img/clarion-yellow.jpg?raw=true)

## Guiding Principles

### Readability is Paramount
Programmers spend the majority of their careers looking at text. Your eyes are an important resource, so a good colorscheme should be as readable as possible, minimizing eyestrain and maximizing comprehension.

Clarion is inspired by research on readability and color, and seeks to put this research to the test. See [the specification](SPEC.md) for more information.

### Minimal Color Highlighting
If everything is important, nothing is! You're here to code, aren't you? Most syntax highlighting is simply choosing random colors that seem to work well together. It's an art not a science, and Clarion aims to skew more towards the science side, only highlighting those things which can benefit from it.

As semantic highlighting and advanced language server features become more prevelant, Clarion will embrace this preferentially insofar as it is not distracting.

### Spec-Driven
A specification keeps things centralized and consistent. Inspired by the [Dracula](https://draculatheme.com/) Theme's [specification][https://spec.draculatheme.com/]. Clarion is driven by a [similar spec](SPEC.md), and in fact uses this doc to generate the theme itself, making the spec the source of truth.

### Conceptual Palette
Rather than simply picking values from a color palette and arbitrarily assigning them to syntax or UI elements, Clarion seeks to define a "conceptual palette", where each color has a specific meaning that applies consistently across contexts.

If you see `problem`![problem swatch](https://via.placeholder.com/15/b50000.png?text=+), there is a serious problem of some kind within that context. Similarly, if you see `new`![new swatch](https://via.placeholder.com/15/4b6319.png?text=+), something has been added or is otherwise "new". Correspondingly, if something doesn't have a spceial meaning, it will not have a color. 

## Status
Clarion is an untested proof-of-concept, and very much a work in progress. The specification is very bare-bones and its format and content will likely change as more values are codified into it. These include such things that currently only exist in templates or the spec generation code, such as editor-specific styling identifiers.

The current stable target is a published VSCoce colorscheme in the Visual Studio marketplace. This alone is a large undertaking, since there are hundreds of different color directives in VSCode, many of which are derivied or are simple alpha deltas, so care must be taken to exhaustively specify as many of these as possible to avoid "surprise" colors that are inconsistent with the goal of a very limited palette.

For this reason, creating themes for other editors is beyond the current scope and will rely on more tooling for spec extraction and templating. Community contributions are very much desired, and Go template functions already exist to render hex values as well as 256-color terminal approximations.

It is still an open question as to whether the "conceptual palette" is a meaningful abstraction and how well it conforms to what science tells us about color and reading comprehension.

While the primary source work that Clarion is based off of targeted Dyslexia, more research on visual ergonomics should be explored, and other disabilities related to reading comprehension, such as colorblindness, should be accomodated as much as possible.

## Development/Contributing
TBD.
