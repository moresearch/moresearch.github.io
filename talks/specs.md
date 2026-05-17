# Talks template (strict two-column)

This spec defines the strict two-column Hackspree talks viewer.

Design & behavior
- Desktop: two columns (340px left identity rail matching homepage; right column is the PDF viewer and visually dominant).
- Left rail replicates the homepage identity: logo (real <img>), identity block (name/description), social links, and talk actions.
- Right column contains talk heading and PDF iframe that fills available height.

Implementation rules
- Template: talks/template.html. Generated pages replace {{ID}} with the zero-padded talk id and copy the PDF to /talks/NNN/NNN.pdf.
- Logo must be a real <img> and use the homepage asset (https://hackspree.com/logo.png) unless the asset is copied into the blog output.
- Use Font Awesome for icons only. Link to the official stylesheet (CDN allowed). Do not set Font Awesome as a page font.
- Grid: `.talk-shell { grid-template-columns: 340px minmax(0,1fr); }` and both columns must use `min-width:0` to avoid overflow.
- PDF sizing: `.pdf-frame { height: calc(100dvh - 120px); }` and iframe fills container.

Validation checklist
- Logo image returns 200
- PDF returns 200
- No horizontal overflow at desktop and mobile
- Font Awesome icons render for social-links and actions
- Talk actions work (Back, Download, Fullscreen)
