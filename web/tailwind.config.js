module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true
  },
  purge: false,
  theme: {
    container: {
      center: true,
      padding: {
        default: '1rem',
        sm: '2rem',
        lg: '4rem',
        xl: '5rem'
      }
    },
    extend: {
      fontFamily: {
        body:
          'Inter, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif'
      },

      colors: {
        primary: '#ff3c00',
        secondary: '#676778'
      },
      strokeWidth: {
        50: '50'
      },
      height: {
        box: '320px'
      }
    }
  },
  variants: {}
};
