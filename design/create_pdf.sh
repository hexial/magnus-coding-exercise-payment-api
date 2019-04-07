#!/bin/sh
pandoc \
    --standalone \
    --css styles/pandoc.css \
    --lua-filter=pdc-links-target-blank.lua \
    -V geometry:"top=2.54cm, bottom=2.54cm, left=3.81cm, right=3.81cm" \
    --from markdown \
    --variable papersize=A4 \
    --pdf-engine=xelatex \
    --output design.pdf DESIGN.md; \
