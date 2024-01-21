import { vtScanUrl, vtUrlAnalysis } from '../utils/virus-total';
import * as logger from '../utils/logger';

const extractUrlsFromString = (str) => {
  const regex = /https?:\/\/[^\s/$.?#].[^\s]*/gi;
  const matches = str.match(regex);
  return [...new Set(matches)];
};

const urlSafetyCheck = async (url) => {
  let data = await vtScanUrl(url);
  data = await vtUrlAnalysis(data.data.id);
  const stats = data.data.attributes.stats;
  logger.debug(JSON.stringify(stats));
  return !(stats.malicious > 0 || stats.suspicious > 0);
};

export default {
  extractUrlsFromString,
  urlSafetyCheck,
};
