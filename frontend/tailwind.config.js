/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{js,vue}'],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    styled: true,
    themes: ['emerald', 'night', 'aqua', 'cmyk'],
    darkTheme: 'night',
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: '',
  },
}
