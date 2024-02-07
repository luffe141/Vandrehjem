import style from "./restaurant.module.css";

const imgChange = (event: any) => {
  const selectedValue = event.target.value;
  console.log('Selected value:', selectedValue);
};


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
          <select onChange={imgChange} name="meals"  className={style.meals}>
            <option value="1">Morgenmad</option>
            <option value="2">Frokost</option>
            <option value="3">Aftensmad</option>
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
