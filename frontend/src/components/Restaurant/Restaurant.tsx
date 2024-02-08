import React, { useState } from "react";
import style from "./restaurant.module.css";

const ImageSwitcher: React.FC = () => {
  const [selectedImage, setSelectedImage] = useState<string>('');

  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedImage(event.target.value);
  };

  return (
    <div>
      <select value={selectedImage} onChange={handleChange}>
        <option value="morgenmad">Morgenmad</option>
        <option value="frokost">Frokost</option>
        <option value="aftensmad">Aftensmad</option>
      </select>

      <div className={style.madBilleder}>
        {selectedImage === "morgenmad" && (
          <div>
            <img src="https://placehold.co/600x300" alt="Billede 1" />
            <img src="https://placehold.co/600x300" alt="Billede 2" />
            <img src="https://placehold.co/600x300" alt="Billede 3" />
          </div>
        )}

        {selectedImage === "frokost" && (
          <div>
            <img src="https://placehold.co/600x400" alt="Billede 1" />
            <img src="https://placehold.co/600x400" alt="Billede 2" />
            <img src="https://placehold.co/600x400" alt="Billede 3" />
          </div>
        )}

        {selectedImage === "aftensmad" && (
          <div>
            <img src="https://placehold.co/600x200" alt="Billede 1" />
            <img src="https://placehold.co/600x200" alt="Billede 2" />
            <img src="https://placehold.co/600x200" alt="Billede 3" />
          </div>
        )}
      </div>
    </div>
  );
};



const Restaurant = () => {
  return (
    <div className={style.mainContainer}>
      <div className={style.highlightedContainer}>
        <div>
          <img src="https://placehold.co/600x400" alt="Pandekager" />
          <p>Pandekager</p>
        </div>
        <div>
          <img src="https://placehold.co/600x400" alt="Suppe" />
          <p>Suppe</p>
        </div>
        <div>
          <img src="https://placehold.co/600x400" alt="Kartoffelmos" />
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
        <ImageSwitcher />
      </div>
    </div>
  );
};

export default Restaurant;

