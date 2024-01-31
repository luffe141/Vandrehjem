import React from "react";
import styles from "./Rooms.module.css";

const Rooms = () => {
  return (
    <div>
      <h1 className={styles.h1}>Rooms</h1>
      <div className={styles.container}>
        <div className={styles.roomDiv}>
          <div className={styles.roomImgDiv}>
            <div className={styles.topImgContainer}>
              <img src="https://placehold.jp/250x150.png" alt="" />
              <img src="https://placehold.jp/100x100.png" alt="" />
            </div>
            <div className={styles.bottomImgContainer}>
              <img src="https://placehold.jp/100x100.png" alt="" />
              <img src="https://placehold.jp/100x100.png" alt="" />
            </div>
          </div>
        </div>
        <div className={styles.roomInfo}>
          <h2>1-4 personers værelse</h2>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Omnis
            deserunt sit non quidem tenetur nihil nisi reiciendis corporis eius!
            Porro maxime adipisci perferendis cum debitis facere facilis
            quisquam repellat tempore.
          </p>
          <label htmlFor="">Pris,-</label>
          <button>Book nu</button>
        </div>
      </div>
    </div>
  );
};

export default Rooms;
