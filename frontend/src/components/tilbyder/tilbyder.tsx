import style from "./tilbyder.module.css";

function tilbyder() {
  return (
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
  );
}

export default tilbyder;
