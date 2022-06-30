/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{js,vue}'],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    styled: true,
    themes: ['aqua', 'aqua', 'cmyk'],
    darkTheme: 'aqua',
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: '',
  },
}
