import httpApi from '@/util/http-api.js';
import { querySerializer } from '@/util/query.js';

const api = {
  getContent: (cb, params, errHandler) => {
    httpApi.getDataViaApi(`/api/content-translations${querySerializer(params)}`, cb, errHandler);
  },
  createRoom: (cb, params, errHandler) => {
    httpApi.postDataViaApi(`/api/create-room${querySerializer(params)}`, cb, errHandler);
  },
};

export default api;
