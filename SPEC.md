# Clarion Spec
This is the specification for Clarion. This file is parsed to generate the color scheme, and so represents the current state of the theme.

It is based on an interpretation of the best research available to the author at the time of this writing.

## Background Colors
Taken from [Good Background Colors for Readers: A Study of People with and without Dyslexia][1]. Peach was listed highest in both test groups, and is thus the default, though both of the others came quite close. Should further research necessitate changes, ordering and colors may change, too.

### Peach
* Swatch: <span style="background-color:#edd1b0;border:2px solid black
">&emsp;</span>
* Hex: #edd1b0

### Orange
* Swatch: <span style="background-color:#eddd6e;border:2px solid black
">&emsp;</span>
* Hex: #eddd6e

### Yellow
* Swatch: <span style="background-color:#f8fd89;border:2px solid black
">&emsp;</span>
* Hex: #f8fd89

## Foreground Colors

Black was chosen as the foreground color for its ubiqity and contrast and derived from the experimental methods in [[1]].

### Black
* Swatch: <span style="background-color:#000000;border:2px solid black
">&emsp;</span>
* Hex: #000000

## Conceptual Colors
Palette derived from [Colorgorical][2] implementation [here][3], with all sliders maxed out.

While the distribution of colors in the CIELAB space is derived from research, the Colorgorical implementation is not deterministic, and the current color palette was derived through trial-and error.

So too, the "conceptual palette" is fiat from the author's (Andy Walker) own ideas about what concepts are important.

Therefore personal bias towards "notable concepts" and "good concept:color concordance" played a significant role in the selection of conceptual colors, making this by far the weakest part of the spec with respect to actual science.

If more concrete sources to inform either the chromatic or conceptual palettes are found, they will be incorporated without hesitation.

### Problem
* Swatch: <span style="background-color:#c31a31;border:2px solid black
">&emsp;</span>
* Hex: #c31a31

#### Rationale
TBD

#### Usage
* errors
* fatal problems

### Notice
* Swatch: <span style="background-color:#d47a37;border:2px solid black
">&emsp;</span>
* Hex: #d47a37

#### Rationale
TBD

#### Usage
* notable warnings
* non-fatal problems
* failed tests

### Affirm
* Swatch: <span style="background-color:#769d31;border:2px solid black
">&emsp;</span>
* Hex: #769d31

#### Rationale
TBD

Note: not for passing tests. Failure is notable. Success is not.

#### Usage
* successes when contrast with failure is notable

### Meta
* Swatch: <span style="background-color:#6e5fae;border:2px solid black
">&emsp;</span>
* Hex: #6e5fae

#### Rationale
TBD

#### Usage
* Comments and non-code annotations

### Modified
* Swatch: <span style="background-color:#b76f7c;border:2px solid black
">&emsp;</span>
* Hex: #b76f7c

#### Rationale
TBD

#### Usage
* Diff modifications

### Added
* Swatch: <span style="background-color:#1a8298;border:2px solid black
">&emsp;</span>
* Hex: :#1a8298

#### Rationale
TBD

#### Usage
* Diff additions

### Removed
* Swatch: <span style="background-color:#cccccc;border:2px solid black
">&emsp;</span>
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