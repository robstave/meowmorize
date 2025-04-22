# Educational Flashcard Creation Guide

## Card Structure Requirements

Each flashcard must follow this exact structure:

<!-- Card Start -->

### Front
[Question or concept to learn]

### Back
[Answer or explanation]

<!-- Card End -->

## Formatting Guidelines

1. **Headers**: Always use `###` (H3) for "Front" and "Back" headers, never `##` or `#`
2. **Content Organization**:
   - Keep front-side content concise (1-3 sentences ideal)
   - Back-side can be more detailed but focused on key points
   - Use bullet points for lists of related items

3. **Rich Formatting Options**:
   - **Bold**: Use `**text**` for important terms or key concepts
   - **Italic**: Use `*text*` for emphasis
   - **Lists**: Use bullet points (`- item`) or numbered lists (`1. item`)
   - **Code**: Use backticks for inline `code` or triple backticks for code blocks
   - **Tables**: Use Markdown tables for structured data
   - **Images**: Include with `![alt text](image_url)` when relevant

## Card Types

### Standard Knowledge Card
<!-- Card Start -->
### Front
What is Object-Oriented Programming?

### Back
**Object-Oriented Programming (OOP)** is a programming paradigm based on the concept of "objects" that contain:

- Data in the form of properties/attributes
- Code in the form of methods
- Objects can interact with one another through their methods

Key principles include: encapsulation, inheritance, polymorphism, and abstraction.
<!-- Card End -->

### Multiple Choice Card
<!-- Card Start -->
### Front
Which AWS service is best suited for managing the entire machine learning lifecycle?  

A) Amazon Comprehend  
B) Amazon SageMaker  
C) Amazon Polly  
D) Amazon Translate

### Back
**Correct Answer**: B

**Explanation**: Amazon SageMaker is designed to manage the entire machine learning lifecycle, including data labeling, model training, deployment, and monitoring.

*Related services*: 
- Amazon Comprehend (natural language processing)
- Amazon Polly (text-to-speech)
- Amazon Translate (language translation)
<!-- Card End -->

### Code Example Card
<!-- Card Start -->
### Front
Write a Python function to check if a string is a palindrome.

### Back

```python
def is_palindrome(text):
    # Remove spaces and lowercase
    clean_text = text.lower().replace(" ", "")
    # Check if text equals its reverse
    return clean_text == clean_text[::-1]
    
# Example usage
print(is_palindrome("radar"))  # True
print(is_palindrome("hello"))  # False
```

<!-- Card End -->


Additional Guidelines
Card Length: Keep cards focused - if explanation exceeds 10 lines, consider breaking into multiple cards
Citations: Include source references when appropriate
Related Concepts: Link related ideas at the end of the back side
Difficulty Levels: Consider marking cards as [Basic], [Intermediate], or [Advanced]
Diagrams: Use ASCII/text diagrams when visual representation would help
Remember that effective flashcards facilitate active recall rather than passive reading.