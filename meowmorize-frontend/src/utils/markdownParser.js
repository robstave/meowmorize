import { v4 as uuidv4 } from 'uuid';

/**
 * Parses the provided markdown text and extracts card objects.
 * Each card is defined between <!-- Card Start --> and <!-- Card End -->
 * with sections ### Front and ### Back.
 *
 * @param {string} markdownText - The markdown content to parse.
 * @returns {Array} - An array of card objects.
 */
export const parseMarkdownToCards = (markdownText) => {
  const cardRegex = /<!--\s*Card Start\s*-->([\s\S]*?)<!--\s*Card End\s*-->/g;
  const frontRegex = /###\s*Front[\s\S]*?(?=###\s*Back)/i;
  const backRegex = /###\s*Back([\s\S]*?)(?=<!--|$)/i;
  const linkRegex = /<!---\s*Card Link\s*-->\s*(https?:\/\/[^\s]+)/i;

  const cards = [];
  let match;

  while ((match = cardRegex.exec(markdownText)) !== null) {
    const cardContent = match[1];

    // Extract Front
    const frontMatch = frontRegex.exec(cardContent);
    const frontText = frontMatch ? frontMatch[0].replace(/###\s*Front/i, '').trim() : '';

    // Extract Back
    const backMatch = backRegex.exec(cardContent);
    const backText = backMatch ? backMatch[1].trim() : '';

    // Extract Link
    const linkMatch = linkRegex.exec(cardContent);
    const link = linkMatch ? linkMatch[1].trim() : '';

    if (frontText && backText) {
      cards.push({
        id: uuidv4(), // Temporary ID; backend should assign the actual ID
        deck_id: '', // Will be assigned when importing
        link,
        front: {
          text: frontText,
        },
        back: {
          text: backText,
        },
      });
    }
  }

  return cards;
};
