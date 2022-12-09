/** @type {import('tailwindcss').Config} */

module.exports = {
  content: [
    "./app/templates/*/*.plush.html",
    "./app/render/helpers/**.go",
    "./app/assets/js/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
