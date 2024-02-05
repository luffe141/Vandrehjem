import style from "./Kontakt.module.css";

const Kontakt = () => {
  return (
    <div className={style.container}>
      <div>
        <h1>Kontakt</h1>
        <div className={style.kontaktContainer}>
          <p className={style.kontaktText}>
            Hvis du har et spørgsmål eller kan vi på nogen anden måde kan hjælpe
            dig, så tøv ikke med at tage fat i os, i formularen her under kan du
            indtaste dine oplysninger og så vil vi svare dig hurtigst muligt, du
            er også altid velkommen til at ringe på 40224199
          </p>
        </div>

        <div>
          <div>
            <input
              className={style.input}
              type="text"
              name=""
              id=""
              placeholder="Navn"
            />
            <input
              className={style.input}
              type="text"
              name=""
              id=""
              placeholder="Email"
            />
          </div>
          <div>
            <input
              className={style.input}
              type="text"
              name=""
              id=""
              placeholder="Tlf"
            />
            <input
              className={style.input}
              type="text"
              name=""
              id=""
              placeholder="Emne"
            />
          </div>

          <textarea
            className={style.textarea}
            name=""
            id=""
            cols={30}
            rows={10}
            placeholder="Besked"
          ></textarea>
          <button className={style.button}>Send</button>
        </div>
      </div>
      <div className={style.maps}>
        <div
          className={style.maps}
          dangerouslySetInnerHTML={{
            __html: `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2201.9080403638513!2d10.814114076929235!3d56.50379297337147!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x464e808fabeadcbb%3A0xc1b5b962ca5e779!2sDanhostel%20Gjerrild!5e0!3m2!1sen!2sdk!4v1707122928019!5m2!1sen!2sdk" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`,
          }}
        />
      </div>
    </div>
  );
};

export default Kontakt;
