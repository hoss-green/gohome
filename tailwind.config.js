/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*.html",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {
      fontFamily: {
        player: "'Press Start 2p'",
        sans: "'Red Hat Display'",
        silkscreen: "'Silkscreen'",
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        my_theme: {
          ...require("daisyui/src/theming/themes")["black"],
          "--rounded-box": "0.1rem", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0.1rem", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "0.1rem", // border radius rounded-badge utility class, used in badges and similar
          primary: "#FF3700",
          secondary: "#5941A9",
          accent: "#F9CB24",
          neutral: "#FEFFFE",
        },
      },
      // ,"dim"],
    ],
  },
};
