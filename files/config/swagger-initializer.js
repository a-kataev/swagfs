window.onload = function() {
  window.ui = SwaggerUIBundle({
    url: "{{url}}",
    urls: {{urls}},
    dom_id: "#swagger-ui",
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "{{layout}}",
    validatorUrl: null
  });
};
