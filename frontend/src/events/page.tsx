import React from "react";
import styles from "./page.module.css";

const Event = () => {
  return (
    <div>
      <div>
        <img
          className={styles.heroImg}
          src="https://placehold.jp/250x150.png"
          alt=""
        />
        <hr className={styles.heroHr1} />
        <h1 className={styles.herotext}>Events</h1>
        <hr className={styles.heroHr2} />
      </div>
      <div className={styles.container}>
        <div className={styles.leftImgContainer}>
          <img
            className={styles.leftImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
          <img
            className={styles.leftImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
        </div>
        <div className={styles.eventInfo}>
          <h2>Navnet på event</h2>
          <hr className={styles.eventHr} />
          <h3>Nuværende event</h3>
          <p className={styles.eventText}>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Qui vitae
            vel ratione ipsa iste. Mollitia qui repudiandae laboriosam obcaecati
            deleniti.
          </p>
          <button className={styles.eventCurrentBtn}>se mere</button>
        </div>
        <div className={styles.rightImgContainer}>
          <img
            className={styles.leftImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
          <img
            className={styles.leftImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
        </div>
      </div>
      <hr />
      <div>
        <h4>Kommende events</h4>
        <button>venstre pil</button>
        <button>højre pil</button>
      </div>
      <div>
        <div>
          <img src="" alt="" />
          <button>se mere</button>
        </div>
        <div>
          <img src="" alt="" />
          <button>se mere</button>
        </div>
        <div>
          <img src="" alt="" />
          <button>se mere</button>
        </div>
      </div>
    </div>
  );
};

export default Event;
