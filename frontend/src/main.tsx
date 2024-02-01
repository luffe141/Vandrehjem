import { createRoot } from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./global.css";
import Home from "./pages/home/home.jsx";
import Layout from "./components/layout/layout.tsx";
import Vaerelser from "./components/Vaerelser/Vaerelser.tsx";
import Detsker from "./components/Detsker/Detsker.tsx";

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
						path="/Detsker"
						element={
							<Layout>
								<Detsker />
							</Layout>
						}
					/>

					<Route
						path="/Vaerelser"
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
