import './home.css'
import Tilbyder from '../../components/tilbyder/tilbyder'
import Hero from '../../components/homepage/hero/Hero.tsx'
import Omgivelser from '../../components/homepage/Omgivelser/Omgivelser.tsx'
import Rooms from '../../components/homepage/Rooms/Rooms.tsx'
import Oplev from '../../components/oplevelse/oplev.tsx'
import Event from '../../components/events/event.tsx'
import Footer from '../../components/footer/footer.tsx'

function Home() {

  return (
    <>
    <Hero />
    <Omgivelser />
    <Rooms/>
    <Tilbyder/>
    <Oplev/>
    <Event/>
    <Footer/>
    </>
  )
}

export default Home
