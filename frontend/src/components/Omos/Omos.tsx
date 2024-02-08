import style from "./omos.module.css"

const Omos = () => {
  return (
    <div className={style.container}>
      <div className={style.txt}> <h2 className={style.head}>Lidt om os</h2><p><span className={style.intro}>Vi er Michael og Anna og vi glæder os til at møde Jer.</span>
<br /> <br />
Vi overtog vandrerhjemmet d. 1. september 2022 og har mange planer for stedet, som vi nu er i gang med at realisere. <br /> <br />

Michael har haft et eventbureau og spiller selv meget musik. Han elsker at bygge og opfinde og skal til at i gang med at renovere den togvogn, der står på vandrerhjemmet, så den kan bruges som bar. <br /> <br />

Anna har arbejdet i restaurationsbranchen i ind- og udland i 16 år og er meget glad for at lave mad og forkæle folk. <br /> <br />

Vi går meget op i det gode, personlige værtsskab og hos os vil du opleve hjerterum og masser af hygge. Her er højt til loftet og åbne døre og vi er altid med på en sludder og klar med en kop kaffe.</p></div>
      <div className={style.omOsImg}><img src="https://placehold.co/600x400"/></div>
    </div>
  )
}

export default Omos