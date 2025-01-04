import { formatDistanceToNow } from 'date-fns';

/**
 * Formats a date string into a relative time (e.g., "3 days ago").
 * If the date is null, undefined, or invalid, returns a placeholder ("—").
 *
 * @param {string | null} dateString - The date string to format.
 * @returns {string} - The formatted relative time or a placeholder.
 */
export const formatLastAccessed = (dateString) => {
  if (!dateString) {
    return '—'; // Placeholder for null or undefined dates
  }

  const ZERO_DATE_STRING = '0001-01-01T00:00:00Z';

  if (!dateString || dateString === ZERO_DATE_STRING) {
    return '—'; // Placeholder for null, undefined, or zero dates
  }
  

  const date = new Date(dateString);

  // Check if the date is valid
  if (isNaN(date.getTime())) {
    return '—'; // Placeholder for invalid dates
  }

  return formatDistanceToNow(date, { addSuffix: true });
};
