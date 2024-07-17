import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()]
})

// faroUploader({
//   appName: 'undefined',
//   endpoint: 'https://faro-api-prod-eu-north-0.grafana.net/faro/api/v1',
//   appId: 'undefined',
//   stackId: '984329',
//   // instructions on how to obtain your API key are in the documentation
//   // https://grafana.com/docs/grafana-cloud/monitor-applications/frontend-observability/sourcemap-upload-plugins/#obtain-an-api-key
//   apiKey: faroApiKey,
//   gzipContents: true,
// })
