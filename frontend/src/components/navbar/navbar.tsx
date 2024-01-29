import { useState } from "react";
import styles from "./navbar.module.css";
<link
  rel="stylesheet"
  href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
></link>;

function Navbar() {
  const [isLinksVisible, setIsLinksVisible] = useState(true);

  const toggleLinks = () => {
    setIsLinksVisible(!isLinksVisible);
  };

  return (
    <nav>
		
      <div className={styles.topnav}>
        <a href="#home" className={styles.active}>
          Logo
        </a>

        {/*<!-- Navigation links (hidden by default) -->*/}
        <div
          id="myLinks"
          style={{ display: isLinksVisible ? "block" : "none" }}>
          <a href="#">test 1</a>
          <a href="#">test 2</a>
        </div>

        {/*<!-- "Hamburger menu" / "Bar icon" to toggle the navigation links -->*/}
        <button onClick={toggleLinks} className={styles.icon}>
          Toggle Links
        </button>
        <i className="fa fa-bars"></i>
      </div>

    </nav>
  );
}

export default Navbar;
