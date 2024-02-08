import style from "./oplev.module.css";

function oplev() {
  return (
    <div>
      {" "}
      <h1 className={style.oplevH1}>Oplev Djursland</h1>
      <div className={style.container}>
        <div>
          <div>
            <img src="https://placehold.co/600x400" />
          </div>
        </div>
        <div className={style.bigRight}>
          <img src="https://placehold.co/1200x300" />
          <div className={style.smBot}>
            <div>
              <img src="https://placehold.co/600x400" alt="" />
            </div>
            <div>
              <img src="https://placehold.co/600x400" alt="" />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default oplev;
