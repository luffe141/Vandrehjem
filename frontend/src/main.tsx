import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./global.css";
import Home from "./pages/home/home.jsx";
import Layout from "./components/layout/layout.tsx";

import Event from "../src/events/page.tsx";

import Oplevelse from "../src/oplevelse/page.tsx"


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
          {/* {

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
						path="/oplevelse"
						element={
							<Layout>
								<Oplevelse />
							</Layout>
						}
					/>

					} */}
          {/* Add more routes for additional pages */}
          {
            <Route
              path="/events"
              element={
                <Layout>
                  <Event />
                </Layout>
              }
            />
          }
        </Routes>
      </>
    </Router>
  );
}

				

