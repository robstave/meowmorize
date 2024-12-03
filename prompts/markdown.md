
# Importing from Markdown

I like to use Obsidian for most of my note taking.  Its nice to write flash card notes in.
I also have written an app that can import flashcards from Markdown.

This is the format that the app uses.

``` 
<!-- Card Start -->

### Front

What is the **largest planet** in our solar system?


### Back

The largest planet is **Jupiter**.

![Jupiter](https://example.com/jupiter.jpg)

<!-- Card End -->

```
The format for each card starts with the HTML comment for start and end.
All markdown after `### Front ` will be the front card and `### Back` for the back card.

Ignore any text in the same line as Front.  for example, if its `### Front Capitol Paris`  Treat it as `### Front` Just look for the first part.

Before the end comment...if there is `<!--- Card Link --->` that is a link.

Your job is to create the markdown so I can import it.  Do not show me how to do it, just create the markdown.







