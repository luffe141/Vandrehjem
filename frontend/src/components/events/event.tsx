import style from "./event.module.css";

function event() {
  return (
    <div>
      <h2 className={style.event}>Events</h2>

      <div className={style.container}>
        <div>
          <p className={style.nuEvents}>Nuværende Events</p>
          <img src="https://placehold.co/600x400" />
        </div>
        <div className={style.kommende}>
          <div>
            <p className={style.kommendeTxt}>Kommende events</p>{" "}
            <img src="https://placehold.co/600x100" alt="" />
          </div>

          <div>
            <img src="https://placehold.co/600x100" alt="" />
          </div>

          <div>
            <img src="https://placehold.co/600x100" alt="" />
          </div>
        </div>
      </div>
    </div>
  );
}

export default event;
