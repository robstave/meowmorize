const markdownInstructions = `
# Markdown format

The markdown format is a pretty simple combination of comments and headers that chats can easily produce.
The file should be saved as .txt or .md


## Markdown Create Prompt V2.0

Create a set of educational flashcards in Markdown format using the following structure.

Each card must be wrapped in HTML comments:

\`\`\`
<!-- Card Start -->

[card content]

<!-- Card End -->
\`\`\`

For each card:

1. **Front side:** Begins with \`### Front\` followed by the front text in Markdown.
2. **Back side:** Begins with \`### Back\` followed by the back text.
3. Use standard Markdown formatting (bold, images, etc.) within the content.

**Example card format:**

\`\`\`
<!-- Card Start -->

### Front

Question here

### Back

Answer here

<!-- Card End -->
\`\`\`

Please create the flashcards directly in this format.
`;

export default markdownInstructions;
