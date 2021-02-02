module.exports = {
  env: {
    node: true
  },
  extends: ["standard", "prettier"],
  parser: "@babel/eslint-parser",
  parserOptions: {
    ecmaVersion: 12,
    sourceType: "module"
  },
  rules: {}
}
