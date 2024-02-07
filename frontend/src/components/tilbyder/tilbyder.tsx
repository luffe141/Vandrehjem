import style from "./tilbyder.module.css";

function tilbyder() {
  return (

    <div>
      <h1 className={style.tilbydH1}>Vi tilbyder </h1>
      <div className={style.tilbyd}>
      <div>
        {" "}
        <h2>Mad</h2> <img src="https://placehold.co/600x400" />
      </div>
      <div>
        {" "}
        <h2>Konferancerum</h2>
        <img src="https://placehold.co/600x400" />
      </div>
      <div>
        {" "}
        <h2>Bar</h2>
        <img src="https://placehold.co/600x400" />
      </div>
    </div>
    </div>
    
  );
}

export default tilbyder;
