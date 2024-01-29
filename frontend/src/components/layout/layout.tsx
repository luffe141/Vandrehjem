// Layout.tsx

import React from 'react';
import Navbar from "../navbar/navbar.tsx";

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

			</footer>
		</div>
	);
};

export default Layout;