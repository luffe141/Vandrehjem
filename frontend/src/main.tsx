import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./global.css";
import Home from "./pages/home/home.jsx";
import Layout from "./components/layout/layout.tsx";
import Vaerelser from "./components/Vaerelser/Vaerelser.tsx";
import Detsker from "./components/Detsker/Detsker.tsx";
import Event from "../src/events/page.tsx";
import Oplevelse from "../src/oplevelse/page.tsx";
import Restaurant from "./components/Restaurant/Restaurant.tsx";
import Galleri from "./components/Galleri/Galleri.tsx";
import Omos from "./components/Omos/Omos.tsx";
import Kontakt from "./components/Kontakt/Kontakt.tsx";
import Bestilonline from "./components/Bestilonline/Bestilonline.tsx";
import Togbar from "./components/Togbar/Togbar.tsx"

const rootElement: HTMLElement | null = document.getElementById("root");

if (rootElement) {
  createRoot(rootElement).render(
    <Router>
      <>
        <Routes>
          <Route
            path="/"
            element={
              <Layout>
                <Home />
              </Layout>
            }
          />

          <Route
            path="Detsker"
            element={
              <Layout>
                <Detsker />
              </Layout>
            }
          />

          <Route
            path="Vaerelser"
            element={
              <Layout>
                <Vaerelser />
              </Layout>
            }
          />

          <Route
            path="Oplevelse"
            element={
              <Layout>
                <Oplevelse />
              </Layout>
            }
          />

          <Route
            path="Event"
            element={
              <Layout>
                <Event />
              </Layout>
            }
          />

          <Route
            path="Restaurant"
            element={
              <Layout>
                <Restaurant />
              </Layout>
            }
          />

          <Route
            path="Galleri"
            element={
              <Layout>
                <Galleri />
              </Layout>
            }
          />

          <Route
            path="Togbar"
            element={
              <Layout>
                <Togbar />
              </Layout>
            }
          />

          <Route
            path="Omos"
            element={
              <Layout>
                <Omos />
              </Layout>
            }
          />

          <Route
            path="Kontakt"
            element={
              <Layout>
                <Kontakt />
              </Layout>
            }
          />

          <Route
            path="Bestilonline"
            element={
              <Layout>
                <Vaerelser />
              </Layout>
            }
          />
        </Routes>
      </>
    </Router>
  );
}
