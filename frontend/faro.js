import {createRoutesFromChildren, matchRoutes, Routes, useLocation} from 'react-router-dom';
import {
    initializeFaro,
    createReactRouterV6DataOptions,
    ReactIntegration,
    getWebInstrumentations,
} from '@grafana/faro-react';
import { TracingInstrumentation } from "@grafana/faro-web-tracing";

export default initializeFaro({
    url: 'https://faro-collector-prod-eu-north-0.grafana.net/collect/70f21b4aa9f030446133f18bb884a891',
    app: {
        name: 'localhost',
        version: '1.0.0',
        environment: 'production'
    },

    instrumentations: [
        ...getWebInstrumentations(),

        new TracingInstrumentation(),

        new ReactIntegration({
            router: {
                version: "v6",
                dependencies: {
                    createRoutesFromChildren,
                    matchRoutes,
                    Routes,
                    useLocation
                }
            },
        }),
    ],
});