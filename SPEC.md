# Clarion Spec
This is the specification for Clarion. This file is parsed to generate the color scheme, and so represents the current state of the theme.

It is based on an interpretation of the best research available to the author at the time of this writing.

## Background Colors
Taken from the paper **Good Background Colors for Readers: A Study of People with and without Dyslexia**[[1]]. Peach was listed highest in both test groups, and is thus the default, though both of the others came quite close. Should further research necessitate changes, ordering and colors may change, too.

### Peach
* Swatch: ![#edd1b0](https://via.placeholder.com/15/edd1b0/000000?text=+)
* Hex: #edd1b0

### Orange
* Swatch: ![#eddd6e](https://via.placeholder.com/15/eddd6e/000000?text=+)
* Hex: #eddd6e

### Yellow
* Swatch: ![#f8fd89](https://via.placeholder.com/15/f8fd89/000000?text=+)
* Hex: #f8fd89

## Foreground Colors

Black was chosen as the foreground color for its ubiqity and contrast and derived from the experimental methods in [[1]].

### Black
* Swatch: ![#000000](https://via.placeholder.com/15/000000/000000?text=+)
* Hex: #000000

## Conceptual Colors
Palette derived from the work in **Colorgorical: Creating discriminable and preferable color palettes for information visualization** paper, using the online implementation with all sliders maxed out.[[2]][[3]]

While the distribution of colors in the CIELAB space is derived from research, the Colorgorical implementation is not deterministic, and the current color palette was derived through trial-and error.

So too, the "conceptual palette" is fiat from the author's (Andy Walker) own ideas about what concepts are important.

Therefore personal bias towards "notable concepts" and "good concept:color concordance" played a significant role in the selection of conceptual colors, making this by far the weakest part of the spec with respect to actual science.

If more concrete sources to inform either the chromatic or conceptual palettes are found, they will be incorporated without hesitation.

### Problem
* Swatch: ![#c31a31](https://via.placeholder.com/15/c31a31/000000?text=+)
* Hex: #c31a31

#### Rationale
TBD

#### Usage
* errors
* fatal problems

### Notice
* Swatch: ![#d47a3](https://via.placeholder.com/15/d47a3/000000?text=+)
* Hex: #d47a37

#### Rationale
TBD

#### Usage
* notable warnings
* non-fatal problems
* failed tests

### Affirm
* Swatch: ![#769d31](https://via.placeholder.com/15/769d31/000000?text=+)
* Hex: #769d31

#### Rationale
TBD

Note: not for passing tests. Failure is notable. Success is not.

#### Usage
* successes when contrast with failure is notable

### Meta
* Swatch: ![#6e5fae](https://via.placeholder.com/15/6e5fae/000000?text=+)
* Hex: #6e5fae

#### Rationale
TBD

#### Usage
* Comments and non-code annotations

### Modified
* Swatch: ![#b76f7c](https://via.placeholder.com/15/b76f7c/000000?text=+)
* Hex: #b76f7c

#### Rationale
TBD

#### Usage
* Diff modifications

### Added
* Swatch: ![#1a8298](https://via.placeholder.com/15/1a8298/000000?text=+)
* Hex: :#1a8298

#### Rationale
TBD

#### Usage
* Diff additions

### Removed
* Swatch: ![#cccccc](https://via.placeholder.com/15/cccccc/000000?text=+)
* Hex: :#cccccc

#### Rationale
TBD

#### Usage
* Diff removals
* Skipped tests

## Caveat

I (Andy Walker) am not an expert in ergonomics, reading comprehension, color theory or related disciplines. I'm just a programmer interpreting the work of actual experts.

I welcome all forms of feedback, informed correction and collaboration with actual experts in the relevant fields, and will readily incorporate suggested changes to this specification, and grant primary credit for any contributions. All I care about is having the most informed colorscheme possible to ease the life of my fellow programmers and anyone who needs to stare at text all day.

## Links and Citations
[[1]] Rello, L. & Bigham, J. P. (2017). Good Background Colors for Readers: A Study of People with and without Dyslexia. Proceedings of the 19th International ACM SIGACCESS Conference on Computers and Accessibility, 72-89.
[(alternate source)](https://www.cs.cmu.edu/~jbigham/pubs/pdfs/2017/colors.pdf)

[[2]] Gramazio, C. C. et al. (2016). Colorgorical: Creating discriminable and preferable color palettes for information visualization. IEEE Transactions on Visualization and Computer Graphics, 23(1), 521-530. [(alternate source)](http://vrl.cs.brown.edu/color/pdf/colorgorical.pdf)

[[3]] http<area>://vrl.cs.brown.edu/color
 
[1]: https://doi.org/10.1145/3132525.3132546
[2]: https://doi.org/10.1109/TVCG.2016.2598918
[3]: http://vrl.cs.brown.edu/color
[4]: https://spec.draculatheme.com/