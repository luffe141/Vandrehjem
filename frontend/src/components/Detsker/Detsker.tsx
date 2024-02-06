import styles from "./Detsker.module.css";
const Detsker = () => {
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
            className={styles.rightImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
          <img
            className={styles.rightImg}
            src="https://placehold.jp/350x350.png"
            alt=""
          />
        </div>
      </div>
      <hr />
      <div>
        <div>
          <h4>Kommende events</h4>
        </div>
        <div className={styles.arrowContainer}>
          <button>venstre pil</button>
          <button>højre pil</button>
        </div>
      </div>
      <div className={styles.upcommingEventsContainer}>
        <div className={styles.card}>
          <img src="https://placehold.jp/350x350.png" alt="" />
          <button className={styles.upcommingEventsBtn}>se mere</button>
        </div>
        <div className={styles.card}>
          <img src="https://placehold.jp/350x350.png" alt="" />
          <button className={styles.upcommingEventsBtn}>se mere</button>
        </div>
        <div className={styles.card}>
          <img src="https://placehold.jp/350x350.png" alt="" />
          <button className={styles.upcommingEventsBtn}>se mere</button>
        </div>
      </div>
    </div>
  )
}

export default Detsker