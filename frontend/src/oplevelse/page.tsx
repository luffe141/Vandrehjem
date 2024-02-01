import style from "./page.module.css";

function page() {
  return (
    <div >
      <div className={style.djursContainer}>
        <div>
          <img src="https://placehold.co/600x400" alt="" />
        </div>
        <div>
          <h2>Djurs sommerland</h2>
          <p>
            Danmarks største forlystelsespark, med vilde forlystelser og mange
            hyggelige spisesteder
          </p>
          <h2>20 KM</h2>
        </div>
      </div>

      <div className={style.katContainer}>
        <div>
          <h2>Kattagat Centret</h2>
          <p>
            Kattegatcentret er et imponerende akvarium og havcenter, der ligger
            ved Grenaa på Djursland, Danmark. Centret er kendt for sin
            dedikation til at formidle viden om det marine liv i
            Kattegat-området. Med mere end 250 forskellige arter og over 2,5
            millioner liter vand tilbyder Kattegatcentret en unik oplevelse for
            besøgende i alle aldre.
          </p>
          <h2>15 KM</h2>
        </div>
        <div>
          <img src="https://placehold.co/600x400" />
        </div>
      </div>

      <div className={style.nationalContainer}>
        <div>
          <img src="https://placehold.co/600x400" />
        </div>
        <div>
          <h2>Nationalpark Mols Bjerge</h2>
          <p>
            Nationalpark Mols Bjerge er en naturskøn perle beliggende på
            Djursland, Danmark. Parken blev etableret for at bevare og fremme
            det unikke landskab og den rige biodiversitet i området. Mols Bjerge
            byder på en imponerende variation af natur, herunder bakker, skove,
            søer og kyststrækninger.
          </p>
          <h2>40 KM</h2>
        </div>
      </div>

      <div className={style.skandiContainer}>
        <div>
          <h2>Skandinavisk Dyrepark</h2>
          <p>
            Skandinavisk Dyrepark er en spændende destination beliggende ved
            Kolind på Djursland, Danmark. Parken er kendt for at give besøgende
            en unik oplevelse af skandinavisk dyreliv og natur. Med mere end 50
            arter af dyr, herunder elge, bjørne, ulve og ræve, byder
            Skandinavisk Dyrepark på en autentisk og nærmest vildeoplevelse.
          </p>
          <h2>33 KM</h2>
        </div>
        <div>
          <img src="https://placehold.co/600x400" />
        </div>
      </div>

      <div className={style.reeContainer}>
        <div>
          <img src="https://placehold.co/600x400" />
        </div>
        <div>
          <h2>Ree Park Safari</h2>
          <p>
            Ree Park Safari er en spændende dyrepark beliggende i Ebeltoft,
            Danmark. Parken er kendt for at tilbyde besøgende en unik
            safarioplevelse med en imponerende samling af vilde dyr fra hele
            verden. Med mere end 100 forskellige dyrearter, herunder elefanter,
            løver, giraffer og næsehorn, giver Ree Park Safari gæsterne mulighed
            for at opleve dyrene i store, naturlige områder, der efterligner
            deres naturlige levesteder.
          </p>
          <h2>34 KM</h2>
        </div>
      </div>

      <div className={style.gjerrildContainer}>
        <div>
          <h2>Gjerrild Nordstrand</h2>
          <p>
            Gjerrild Nordstrand, beliggende på Djursland, er en betagende
            kyststrækning, der tiltrækker besøgende med sin rå og naturskønne
            skønhed. Den brede sandstrand og de dramatiske klinter skaber en
            imponerende kontrast, der gør dette sted unikt. Nordstranden er et
            ideelt sted for naturelskere, der ønsker at nyde rolige gåture langs
            vandkanten eller beundre solnedgangen over havet.
          </p>
          <h2>3,5km</h2>
        </div>
        <div>
          <img src="https://placehold.co/600x400" />
        </div>
      </div>
    </div>
  );
}

export default page;
