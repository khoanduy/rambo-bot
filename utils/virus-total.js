import { request } from 'axios';
import { virusTotalApiKey } from '../config.json';

const vtScanUrl = async (url) => {
  const { URLSearchParams } = require('url');

  const encodedParams = new URLSearchParams();
  encodedParams.set('url', url);

  const options = {
    method: 'POST',
    url: 'https://www.virustotal.com/api/v3/urls',
    headers: {
      accept: 'application/json',
      'x-apikey': virusTotalApiKey,
      'content-type': 'application/x-www-form-urlencoded',
    },
    data: encodedParams,
  };

  const response = await request(options);
  return response.data;
};

const vtUrlAnalysis = async (id) => {
  const options = {
    method: 'GET',
    url: `https://www.virustotal.com/api/v3/analyses/${id}`,
    headers: {
      accept: 'application/json',
      'x-apikey': virusTotalApiKey,
    },
  };

  const response = await request(options);
  return response.data;
};

export default {
  vtScanUrl,
  vtUrlAnalysis,
};
