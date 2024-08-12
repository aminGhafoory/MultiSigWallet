import presetQuick from "franken-ui/shadcn-ui/preset-quick";

/** @type {import('tailwindcss').Config} */

module.exports = {

  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],


  presets: [presetQuick()],



}







