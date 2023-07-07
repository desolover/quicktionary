### quicktionary

Golang-lib for parsing russian-language-words from [Wiktionary](https://ru.wiktionary.org/).

#### Packages
* `quicktionary` — for parsing [XML-dumps](https://dumps.wikimedia.org/ruwiktionary/) (stored in "pages-articles-multistream" files)
* `speech` — for decoding speech-parts templates.

Code example, converting speech-parts templates from "pages-articles-multistream" XML to internal form:
```golang
file, err := os.Open(`./ruwiktionary-20210801-pages-articles-multistream.xml`)
if err != nil {
    return err
}
defer file.Close()

parsedTemplates, err := quicktionary.ParseTemplatesSource(file)
if err != nil {
    return err
}

templatesFile, err := os.Create(`./templates.json`)
if err != nil {
    return err
}
defer templatesFile.Close()

if err = json.NewEncoder(templatesFile).Encode(parsedTemplates); err != nil {
    return err
}
```

Code example, parsing words:
```golang
file, err := os.Open(`./ruwiktionary-20210801-pages-articles-multistream.xml`)
if err != nil {
    return err
}
defer file.Close()

templatesFile, err := os.Open(`./templates.json`)
if err != nil {
    return err
}
defer templatesFile.Close()

templates, err := speech.LoadTemplates(templatesFile)
if err != nil {
    return err
}

parser := NewParser(templates)
words, err := parser.Do(file)
if err != nil {
    return err
}
```