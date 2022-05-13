module.exports = {
  extends: [
    // add more generic rulesets here, such as:
    // 'plugin:vue/vue3-recommended',
    'plugin:vue/recommended', // Use this if you are using Vue.js 2.x.
  ],
  rules: {
    "vue/max-attributes-per-line": ["error", {
      "singleline": {
        "max": 3
      },
      "multiline": {
        "max": 3
      }
    }],
    "indent": ["error", 2]
  }
}