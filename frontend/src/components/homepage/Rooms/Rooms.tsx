import styles from "./Rooms.module.css";

const Rooms = () => {
  return (

    <div>
      <h1 id={styles.roomsTitle}>Rooms</h1>
      <div className={styles.container}>
          <div className={styles.roomImgDiv}>
            <div className={styles.topImgContainer}>
              <img src="https://placehold.jp/250x150.png" alt="img_1" />
              <img src="https://placehold.jp/100x100.png" alt="img_2" />
            </div>
            <div className={styles.bottomImgContainer}>
              <img src="https://placehold.jp/100x100.png" alt="img_3" />
              <img src="https://placehold.jp/100x100.png" alt="img_4" />
            </div>
          </div>
        <div className={styles.roomInfo}>
          <h2>1-4 personers værelse</h2>
          <p>
            Lorem ipsum deez nuts sit amet consectetur adipisicing elit. Omnis
            deserunt sit non quidem tenetur nihil nisi reiciendis corporis eius!
            Porro maxime adipisci perferendis cum debitis facere facilis
            quisquam repellat tempore.
          </p>
          <label htmlFor="">Pris,- 495,00</label>
          <button><a href="/Vaerelser">Book nu</a></button>
        </div>
      </div>
    </div>
  );
};


export default Rooms;
