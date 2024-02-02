// Layout.tsx

import React from 'react';
import Navbar from "../navbar/navbar.tsx";
import './layout.module.css'
import Footer from "../footer/footer.tsx";

interface AppLayoutProps {
	children: React.ReactNode;
}

const Layout: React.FC<AppLayoutProps> = ({ children }) => {
	return (
		<div>
			<header>
				<Navbar></Navbar>
			</header>
			<main>{children}</main>
			<footer>
				<Footer></Footer>
			</footer>
		</div>
	);
};

export default Layout;