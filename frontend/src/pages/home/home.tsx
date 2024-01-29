import './home.css'
import Tilbyder from '../../components/tilbyder/tilbyder'
import Hero from '../../components/homepage/hero/Hero.tsx'
import Omgivelser from '../../components/homepage/Omgivelser/Omgivelser.tsx'
import Rooms from '../../components/homepage/Rooms/Rooms.tsx'

function Home() {

  return (
    <>
    <Hero />
    <Omgivelser />
    <Rooms/>
    <Tilbyder/>
    </>
  )
}

export default Home
