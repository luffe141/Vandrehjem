import style from "./restaurant.module.css";

const Restaurant = () => {
  return (
    <div className={style.mainContainer}>
      <div className={style.highlightedContainer}>
        <div>
          <img src="https://placehold.co/600x400" />
          <p>Pandekager</p>
        </div>
        <div>
          <img src="https://placehold.co/600x400" />
          <p>Suppe</p>
        </div>
        <div>
          <img src="https://placehold.co/600x400" />
          <p>Kartoffelmos</p>
        </div>
      </div>
      <div className={style.menuContainer}>
        <div>
          <h2>Se vores menu</h2>
        </div>
        <div className={style.kategori}>
          <p>Kategori</p>
        </div>
        <div>
          <select name="meals" className={style.meals}>
            <option value="Morgenmad">Morgenmad</option>
            <option value="Frokost">Frokost</option>
            <option value="Aftensmad">Aftensmad</option>
          </select>
        </div>
      </div>
      <div className={style.morgenmad}>
        <div><img src="https://placehold.co/600x400" /></div>
        <div><img src="https://placehold.co/600x400" /></div>
        <div><img src="https://placehold.co/600x400" /></div>
        <div><img src="https://placehold.co/600x400" /></div>
        <div><img src="https://placehold.co/600x400" /></div>
        <div><img src="https://placehold.co/600x400" /></div>
      </div>
    </div>
  );
};

export default Restaurant;
