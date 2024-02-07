import styles from "./Galleri.module.css";
import React from "react";

const Galleri = () => {

  const imageUrls = [
    "https://placehold.co/200x200",
    "https://placehold.co/200x200",
    "https://placehold.co/200x200",
    "https://placehold.co/200x200"
  
  ];

  return (
    <div>
      <h2>Galleri</h2>
      <div className={styles.galleri}>
        {" "}
        {imageUrls.map((url, index) => (
          <img key={index} src={url} alt="" />
        ))}
      </div>
    </div>
  );
};
export default Galleri;
