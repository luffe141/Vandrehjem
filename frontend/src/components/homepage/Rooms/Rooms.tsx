import React from "react";
import styles from "./Rooms.module.css";

const Rooms = () => {
  const images: string[] = [
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse1.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse2.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/danhostel-gjerrild.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/danhostel-gjerrild.jpg",
  ];
  return (
    <div className={styles.bigContainer}>
      <h1 className={styles.roomH1}>Rooms</h1>
      <div className={styles.container}>
        <div className={styles.roomDiv}>
          <div className={styles.roomImgDiv}>
            <div className={styles.topImgContainer}>
              <div className={styles.bottomImgContainer}>
                <img
                  className={styles.bottomLeftImg}
                  src={images[2]}
                  alt="Room Image 3"
                />
                <img
                  className={styles.bottomRightImg}
                  src={images[3]}
                  alt="Room Image 4"
                />
              </div>
              <div className={styles.bottomImgContainer}>
                <img
                  className={styles.topLeftImg}
                  src={images[0]}
                  alt="Room Image 1"
                />
                <img
                  className={styles.topRightImg}
                  src={images[1]}
                  alt="Room Image 2"
                />
              </div>
            </div>
          </div>
          <div className={styles.roomInfo}>
            <h2>1-4 personers værelse</h2>
            <p className={styles.roomText}>
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Omnis
              deserunt sit non quidem tenetur nihil nisi reiciendis corporis
              eius! Porro maxime adipisci perferendis cum debitis facere facilis
              quisquam repellat tempore.
            </p>
            <label htmlFor="">Pris,-</label>
            <button>Book nu</button>
          </div>
        </div>
      </div>
    </div>
  );
};

<div className={styles.roomDiv}>
  <h1>Rooms</h1>
</div>;

export default Rooms;
