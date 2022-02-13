# Clarion Spec
This is the specification for Clarion. This file is parsed to generate the color scheme, and so represents the current state of the theme.

It is based on an interpretation of the best research available to the author at the time of this writing.

## Background Colors
Taken from the paper **Good Background Colors for Readers: A Study of People with and without Dyslexia**. Peach was listed highest in both test groups, and is thus the default, though both of the others came quite close. Should further research necessitate changes, ordering and colors may change, too.[[1]]

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
Palette derived from colorsafe.co with peach as the base color.
* Restrictions: WCAG Standard - AA
* Font Size: 18px
* Font Weight: 400

 
### Outstanding Issues
While the distribution of the Conceptual Palette colors in the CIELAB space is derived from research, the Colorgorical implementation used is not deterministic, and the current color palette was derived through trial-and error and choices are fiat from the author's (Andy Walker) own ideas about what concepts are important. Therefore, personal bias towards "notable concepts" and "good concept:color concordance" played a significant role in the selection of the current conceptual palette, making this by far the weakest part of the spec with respect to actual science.

Additionally, more testing should be done to ensure that the conceptual palette colors are readable against Clarion base colors. This should be as deterministic as possible.

Finally, it is a well-known phenomenon that different colors appear differently against varying color backgrounds. Because conceptual colors are an important principle of Clarion's design, efforts should be undertaken to ensure that they appear the same in theme variations to maintain their utility.[[2]][[3]]

If more concrete sources to inform or programmatically derive the chromatic or conceptual palettes are found, they will be incorporated without hesitation.

Colors were formerly based off of **Colorgorical: Creating discriminable and preferable color palettes for information visualization** paper, using the online implementation with all sliders maxed out.[[4]][[5]], and this work should be revisited if the conceptual paletted needs additions or to be systemitized further.

### Problem
* Swatch: ![problem swatch](https://via.placeholder.com/15/b50000.png?text=+)
* Sample: ![problem sample](https://via.placeholder.com/150x50/edd1b0/b50000.png?text=Problem)
* Hex: #b50000

#### Rationale
TBD

#### Example Usage
* errors
* fatal problems

### Notice
* Swatch: ![notice swatch](https://via.placeholder.com/15/804600.png?text=+)
* Sample: ![notice sample](https://via.placeholder.com/150x50/edd1b0/804600.png?text=Notice)
* Hex: #804600

#### Rationale
TBD

#### Example Usage
* notable warnings
* failed tests

### New
* Swatch: ![new swatch](https://via.placeholder.com/15/4b6319.png?text=+)
* Sample: ![new sample](https://via.placeholder.com/150x50/edd1b0/4b6319.png?text=New)
* Hex: #4b6319

#### Rationale
TBD

#### Example Usage
* Diff additions
* Untracked files

### Modified
* Swatch: ![modified swatch](https://via.placeholder.com/15/3a599b.png?text=+)
* Sample: ![modified sample](https://via.placeholder.com/150x50/edd1b0/3a5998.png?text=Modified)
* Hex: #3a599b

#### Rationale
TBD

#### Example Usage
* Diff modifications
* modified files

### Excluded
* Swatch: ![excluded swatch](https://via.placeholder.com/15/555555.png?text=+)
* Sample: ![excluded sample](https://via.placeholder.com/150x50/edd1b0/555555.png?text=Excluded)
* Hex: #555555

#### Rationale
TBD

#### Example Usage
* Comments
* Ignored files
* Diff removals
* Skipped tests

## Color Permutations
* ΔETarget: 3.0
* LStep: 0.01
* Variations: 6

To generate lighter or darker variants, colors are translated along the L (lightniess) axis of the CIELAB color space until they are "noticeably different", as measured by the CIE Distance metric ΔE* derived using CIEDE2000.[[6]][[7]]

A ΔE* of 1.0 is described as a "just noticable difference" (JND), so Clarion opts for a higher target **ΔETarget** in an attempt to provide greater distinctions. Currently this is done by increasing or decreasing the L value by **LStep** until ΔE* to the prior color meets or exceeds **ΔETarget**.

Background colors get **Variations**/2 lighter variants and **Variations**/2 darker ones. Foreground colors get **Variations** lighter variants.

## Caveats
I (Andy Walker) am not an expert in ergonomics, reading comprehension, color theory or related disciplines. I'm just a programmer interpreting the work of actual experts.

I welcome all forms of feedback, informed correction and collaboration with actual experts in the relevant fields, and will readily incorporate suggested changes to this specification, and grant primary credit for any contributions. All I care about is having the most informed colorscheme possible to ease the life of my fellow programmers and anyone who needs to stare at text all day.

## Links and Citations
[[1]] Rello, L. & Bigham, J. P. (2017). Good Background Colors for Readers: A Study of People with and without Dyslexia. Proceedings of the 19th International ACM SIGACCESS Conference on Computers and Accessibility, 72-89.
[(alternate source)](https://www.cs.cmu.edu/~jbigham/pubs/pdfs/2017/colors.pdf)

[[2]] Simultaneous and Successive Contrast - Color Usage Research Lab, Nasa Ames Research Center

[[3]] Color Discrimination and Identification - Color Usage Research Lab, Nasa Ames Research Center

[[4]] Gramazio, C. C. et al. (2016). Colorgorical: Creating discriminable and preferable color palettes for information visualization. IEEE Transactions on Visualization and Computer Graphics, 23(1), 521-530. [(alternate source)](http://vrl.cs.brown.edu/color/pdf/colorgorical.pdf)

[[5]] http<area>://vrl.cs.brown.edu/color

[[6]] Wikiepedia Page on Color difference - Section on ΔE*
 
[[7]] Wikiepedia Page on Color difference - Section on the CIEDE2000 formula for calculating ΔE*


[1]: https://doi.org/10.1145/3132525.3132546
[2]: https://colorusage.arc.nasa.gov/Simult_and_succ_cont.php
[3]: https://colorusage.arc.nasa.gov/discrim.php
[4]: https://doi.org/10.1109/TVCG.2016.2598918
[5]: http://vrl.cs.brown.edu/color
[6]: https://en.wikipedia.org/wiki/Color_difference#CIELAB_%CE%94E*
[7]: https://en.wikipedia.org/wiki/Color_difference#CIEDE2000
