import { useState } from "react";
import styles from "./navbar.module.css";



function Navbar() {
  
  const [isNavVisible, setNavVisibility] = useState(false);

  const toggleNav = () => {
		setNavVisibility(!isNavVisible);
		if(isNavVisible) {
			document.getElementById("myLinks").style.display = "none"
		} else {
			document.getElementById("myLinks").style.display = "flex"
		}
	};

  return (


    <nav>

				<div className={styles.topnav}>
					<a href="/" className={styles.logo}><img src="logo.svg" alt="logo"/></a>

					{/*<!-- "Hamburger menu" / "Bar icon" to toggle the navigation links -->*/}
					<a className={styles.icon} onClick={toggleNav}>
						<img src="burger-icon.svg" alt="burger_menu_icon"/>
					</a>

					{/*<!-- Navigation links (hidden by default) -->*/}
					<div id="myLinks" className={styles.myLinks}>
						<a href="/">Forside</a>
						<a href="/Detsker">Det sker</a>
						<a href="/Vaerelser">Vaerelser</a>
						<a href="/Oplevelse">Oplevelser</a>
					</div>
				</div>

			</nav>


  );
}

export default Navbar;
