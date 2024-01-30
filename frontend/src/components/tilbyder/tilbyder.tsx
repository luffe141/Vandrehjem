import style from "./tilbyder.module.css"

function tilbyder() {
    return (
        <div className={style.tilbyd}>
            <div><img src="https://placehold.co/600x400" /><p>Mad</p></div>            
            <div><img src="https://placehold.co/600x400" /><p>Konferancerum</p></div>
            <div><img src="https://placehold.co/600x400" /><p>Bar</p></div>
        </div>
    );
}

export default tilbyder;