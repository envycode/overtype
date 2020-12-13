export function querySerializer(params) {
  return `?${Object.keys(params).map((i) => `${i}=${params[i]}`).join('&')}`;
}

export default {
  querySerializer,
};
