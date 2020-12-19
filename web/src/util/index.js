export const goUrl = url => {
  window.location.href = url;
};

export const scrollToTop = () => {
  const scrollStep = window.scrollY / (1000 / 30);
  const scrollInterval = setInterval(() => {
    if (window.scrollY > 0) {
      window.scrollTo(0, window.scrollY - scrollStep);
    } else {
      clearInterval(scrollInterval);
    }
  }, 15);
};

export const isEmpty = value => value === '' || value === null || typeof value === 'undefined';

export const decodeHtml = input => {
  const doc = new DOMParser().parseFromString(input, 'text/html');
  return doc.documentElement.textContent;
};

export default {
  goUrl,
  scrollToTop,
  isEmpty,
  decodeHtml
};
