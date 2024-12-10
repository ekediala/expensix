/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templ/**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
