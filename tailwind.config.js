/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*.html",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        my_theme: {
          ...require("daisyui/src/theming/themes")["synthwave"],
          "--rounded-box": "0.1rem", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0.1rem", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "0.1rem", // border radius rounded-badge utility class, used in badges and similar
        },
      },
      // ,"dim"],
    ],
  },
};
