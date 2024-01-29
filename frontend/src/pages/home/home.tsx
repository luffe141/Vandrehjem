import './home.css'
import Tilbyder from '../../components/tilbyder/tilbyder'
import Hero from '../../components/homepage/hero/Hero.tsx'
import Omgivelser from '../../components/homepage/Omgivelser/Omgivelser.tsx'

function Home() {

  return (
    <>
    <Hero />
    <Omgivelser />
    <Tilbyder/>
    </>
  )
}

export default Home
