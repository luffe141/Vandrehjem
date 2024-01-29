import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './global.css'
import Home from "./pages/home/home.jsx";
import Layout from "./components/layout/layout.tsx";

const rootElement:HTMLElement | null = document.getElementById('root');

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
					{/*
					<Route
						path="/about"
						element={
							<Layout>
								<About />
							</Layout>
						}
					/>
					*/}
					{/* Add more routes for additional pages */}
				</Routes>
			</>
		</Router>
	);
}