
# Importing from Markdown

I like to use Obsidian for most of my note taking. I would like a way to import cards directly from my notes or otherwise.
This is the format that the app will understand

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

The json format for a card will be
```json
      {
        "id": "e0c32c1c-b36f-4e10-9f47-b8e88c8ff383",
        "deck_id": "123e4567-e89b-12d3-a456-426614174000",
        "link":"https://www.amazon.com/map-paris/s?k=map+of+paris",
        "front": {
          "text": "Capital of _France_"
        },
        "back": {
          "text": "Paris"
        }
      }
```







