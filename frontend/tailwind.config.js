/** @type {import("tailwindcss").Config} */
export default {
    content: [
        "./src/**/*.blade.php",
        "./src/**/*.ts",
        "./src/**/*.vue",
    ],
    theme: {
        extend: {
            fontFamily: {
                'handwriting': ['"Architects Daughter"', 'handwriting'],
            }
        },
    },
    plugins: [],
};
