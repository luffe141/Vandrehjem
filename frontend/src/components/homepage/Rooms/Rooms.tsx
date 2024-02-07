import React from "react";
import styles from "./Rooms.module.css";

const Rooms = () => {
  return (

    <div className={styles.bigContainer}>
      <h1 className={styles.roomH1}>Rooms</h1>
      <div className={styles.container}>
        <div className={styles.roomDiv}>
          <div className={styles.roomImgDiv}>
            <div className={styles.topImgContainer}>
              <img className={styles.topLeftImg} src="https://placehold.jp/150x150.png" alt="" />
              <img className={styles.topRightImg} src="https://placehold.jp/150x150.png" alt="" />
            </div>
            <div className={styles.bottomImgContainer}>
              <img className={styles.bottomLeftImg} src="https://placehold.jp/150x150.png" alt="" />
              <img className={styles.bottomRightImg} src="https://placehold.jp/150x150.png" alt="" />
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

    <div className={styles.roomDiv}>
        <h1>Rooms</h1>
    </div>


export default Rooms;
