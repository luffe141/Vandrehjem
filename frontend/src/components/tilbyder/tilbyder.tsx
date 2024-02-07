import style from "./tilbyder.module.css";

function tilbyder() {
  const images: string[] = [
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse1.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse2.jpg",
    "http://localhost:5173/Img/Billeder/Billeder/Vandrerhjem/vaerelse2.jpg",
  ];
  return (
    <div>
      <h1 className={style.tilbydH1}>Vi tilbyder </h1>
      <div className={style.tilbyd}>
        <div>
          <h2>Mad</h2>
          <img src={images[0]} />
        </div>
        <div>
          <h2>Konferancerum</h2>
          <img src={images[1]} />
        </div>
        <div>
          <h2>Bar</h2>
          <img src={images[2]} />
        </div>
      </div>
    </div>
  );
}

export default tilbyder;
