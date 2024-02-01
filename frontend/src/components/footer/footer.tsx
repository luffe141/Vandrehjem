import style from "./footer.module.css";

function footer() {
  return (
    <div className={style.container}>
      <div className={style.kontakt}>
        <h2>Kontakt os</h2>
        <p>
          Dyrehaven 9 <br />
          +45 40224199 <br />
          info@gjerrildvandrehjem.dk
        </p>
      </div>
      <div className={style.links}>
        <h2>Nyttige links</h2>
        <p>
          Galleri <br />
          Kontakt
        </p>
      </div>
      <div >
        <h2 className={style.sponsorH}>Sponsorer</h2>
        <div className={style.sponsors}>
          <div>
            <img src="https://placehold.co/600x400" />
          </div>

          <div className={style.smspons}>
            <img src="https://placehold.co/300x90" />
            <img src="https://placehold.co/300x90" />
          </div>
        </div>
      </div>
      <div className={style.socials}>
        <h2>Sociale Medier</h2>
      </div>
      <div className={style.copyright}>
        <p>© 2023 Gjerrild vandrerhjem. All Rights Reserved</p>
      </div>
    </div>
  );
}

export default footer;
