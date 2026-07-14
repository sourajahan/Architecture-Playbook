// This is the infrastructure config that allows two apps to share code dynamically.
const { ModuleFederationPlugin } = require("webpack").container;

module.exports = {
  plugins: [
    new ModuleFederationPlugin({
      name: "checkout_app",
      filename: "remoteEntry.js",
      // Expose internal components for other teams to consume
      exposes: {
        "./CheckoutButton": "./src/components/CheckoutButton",
      },
      // Shared dependencies prevent multiple versions of React/Vue from loading
      shared: {
        react: { singleton: true, requiredVersion: "^18.0.0" },
        "react-dom": { singleton: true, requiredVersion: "^18.0.0" },
      },
    }),
  ],
};
