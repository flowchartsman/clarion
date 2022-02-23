# Clarion - A Monochrome Theme Inspired By 🧑‍🔬
![Clarion Logo](img/logo.png?raw=true)
Clarion is a mostly-monochromatic, minimally-highlighted colorscheme, clearing
away the rainbow madness and allowing you to concentrate on what matters the
most: your code.

![Clarion White Preview](img/Clarion-White.jpg?raw=true)
![Clarion Orange Preview](img/Clarion-Orange.jpg?raw=true)
![Clarion Peach Preview](img/Clarion-Peach.jpg?raw=true)

## Guiding Principles

### Readability is Paramount
Programmers spend the majority of their careers looking at text. Your eyes are an important resource, so a good theme should be as readable as possible, minimizing eyestrain and maximizing comprehension. Research overwhelmingly suggests that no single background color is better or worse for readability, so long as an good contrast ratio is maintained with monochromatic text. 

*See [the specification](SPEC.md) for more information.*

### Minimal Syntax Highlighting
If everything is important, nothing is! Color is an important tool for conveying important information, but the more it's used, the less meaningful it becomes. There are only a finite amount of colors you can distinguish easily, and the more syntax elements that get a color, the greater the chance of overlap, and the less it will mean. Clarion tries to avoid using color as much as possible, concentrating instead on carefully-chosen font weights for certain landmark elements to help you orient yourself in your ccode.

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
