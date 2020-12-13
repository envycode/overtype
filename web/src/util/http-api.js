import axios from 'axios';
import Cookies from 'js-cookie';
import util from '@/util/index.js';

function dataAdapter(fn) {
  return (res) => {
    const response = res.response || res;

    fn && fn({
      status: response.status,
      statusText: response.status,
      body: response.data,
      headers: response.headers || (res.config && res.config.headers) || {},
      originalResponse: res,
    });
  };
}

function defaultHandlerConnectionError(err) {
  // check if status code is 401, redirect to login
  switch (err.status) {
    case 401:
      util.goUrl('/');
      break;
    default:
    // nothing
  }
}

export default {
  getDataViaApi(path, cb, errorHandler, httpStatusHandler, baseUrl) {
    const headers = {
      'Cache-Control': 'no-cache',
      // 'User-Agent': navigator.userAgent
    };
    if (Cookies.get('token')) {
      headers['X-Auth'] = `${Cookies.get('token')}`;
    }
    axios.get(path, {
      headers,
      baseURL: baseUrl || ENV.BASE_URL
    })
      .then(dataAdapter(cb))
      .catch((error) => {
        dataAdapter(httpStatusHandler || defaultHandlerConnectionError)(error);

        dataAdapter(errorHandler)(error);
      });
  },

  postDataViaApi(path, cb, data, errorHandler, httpStatusHandler, baseUrl) {
    const headers = {
      // 'User-Agent': navigator.userAgent
    };
    if (Cookies.get('token')) {
      headers['X-Auth'] = `${Cookies.get('token')}`;
    }
    axios.post(path, data, {
      headers,
      baseURL: baseUrl || ENV.BASE_URL
    })
      .then(dataAdapter(cb))
      .catch((error) => {
        dataAdapter(httpStatusHandler || defaultHandlerConnectionError)(error);

        dataAdapter(errorHandler)(error);
      });
  },

  deleteDataViaApi(path, cb, data, errorHandler, httpStatusHandler, baseUrl) {
    const headers = {
      // 'User-Agent': navigator.userAgent
    };
    if (Cookies.get('token')) {
      headers['X-Auth'] = `${Cookies.get('token')}`;
    }
    axios.delete(path, {
      headers,
      data,
      baseURL: baseUrl || ENV.BASE_URL
    })
      .then(dataAdapter(cb))
      .catch((error) => {
        dataAdapter(httpStatusHandler || defaultHandlerConnectionError)(error);

        dataAdapter(errorHandler)(error);
      });
  },

  patchDataViaApi(path, cb, data, errorHandler, httpStatusHandler, baseUrl) {
    const headers = {
      // 'User-Agent': navigator.userAgent
    };
    if (Cookies.get('token')) {
      headers['X-Auth'] = `${Cookies.get('token')}`;
    }
    axios.patch(path, data, {
      headers,
      baseURL: baseUrl || ENV.BASE_URL
    })
      .then(dataAdapter(cb))
      .catch((error) => {
        dataAdapter(httpStatusHandler || defaultHandlerConnectionError)(error);

        dataAdapter(errorHandler)(error);
      });
  },
  defaultHandlerConnectionError,
};
