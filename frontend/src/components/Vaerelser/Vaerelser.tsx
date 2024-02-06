import styles from "./Vaerelser.module.css";
import ImageSliderTwo from "./ImageSliderTwo";

const Vaerelser = () => {
  const images: string[] = [
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse1.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse2.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/danhostel-gjerrild.jpg", 
  ];      

  return (
    <div className={styles.vaerelserContainer}>
      <div className={styles.gallerySlider}>
        <ImageSliderTwo images={images} />
      </div>


      <div className={styles.bookingContainer}>
        <div className={styles.bookingGreyBG}>
          <p>1 - 4 personers værelse</p>
          <p>
            Dkk495,00<span>Per nat</span>
          </p>
        </div>
        <div>
          <form>
            <label htmlFor="dato">Tjek ind/ud</label>
            <input type="text" placeholder="Tjek ind -> Tjerk ud" id="dato" />
            <label htmlFor="dato">Gæster</label>
            <input type="text" placeholder="Gæster 1" id="dato" />
          </form>
        </div>
        <div>
          <p>TJEK TILGÆNGELIGHED</p>
        </div>
      </div>


      <div className={styles.roomInfo}>
        <p>
          Vi har plads til 88 overnattende gæster og værelsesgangen er placeret
          i god afstand fra husets fællesrum, så du har ro, når du
          <br /> vil sove. Der er dobbeltsenge på flere af værelserne, køjesenge
          og mulighed for ekstra senge, således at der kan sove alt fra 1 person
          til 5 personer på værelserne.
        </p>
        <p>Værelserne er med:</p>
        <br />
        <ul>
          <li>Eget bad og toilet</li>
          <li>Vinduer der kan åbnes ud til haven</li>
          <li>Et lille skrivebord med stol</li>
          <li>Mulighed for weekendseng og vådlagen.</li>
        </ul>
        <br />
        <p>
          Vi har handicapvenlige toilet- og badefaciliteter og nem adgang rundt
          i huset.
        </p>
      </div>


      <div className={styles.infoBoxes}>
        <div className={styles.paragraphBox}>
          <div className={styles.iconBox}>
            <img src="family.png" alt="icon" />
          </div>
          <div className={styles.iconTextBox}>
            <p>Maks. Gæster</p>
            <p><span>4 (+1) Gæster</span></p>
          </div>
        </div>
        <div className={styles.paragraphBox}>
          <div className={styles.iconBox}>
            <img src="calendar.png" alt="icon" />
          </div>
          <div className={styles.iconTextBox}>
            <p>Booking af overnatninger</p>
            <p><span>1 Min.</span></p>
          </div>
        </div>
        <div className={styles.paragraphBox}>
          <div className={styles.iconBox}>
            <img src="bed.png" alt="icon" />
          </div>
          <div className={styles.iconTextBox}>
            <p>Sengetype</p>
            <p><span>boxmadras + køjeseng</span></p>
          </div>
        </div>
        <div className={styles.paragraphBox}>
          <div className={styles.iconBox}>
            <img src="layout.png" alt="icon" />
          </div>
          <div className={styles.iconTextBox}>
            <p>Areal</p>
            <p><span>m²</span></p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Vaerelser;
