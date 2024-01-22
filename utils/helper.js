const extractUrlsFromString = (str) => {
  const regex = /https?:\/\/[^\s/$.?#].[^\s]*/gi;
  const matches = str.match(regex);
  return [...new Set(matches)];
};

export default {
  extractUrlsFromString,
};
