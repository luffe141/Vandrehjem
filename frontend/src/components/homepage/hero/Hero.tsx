import Kalender from '../../calender/calender'
import styles from './Hero.module.css'




const Hero = () => {
  return (
    <div id={styles.heroDiv}>
      <img id={styles.heroImg} src="../../../../public/Img/Billeder/Billeder/Vandrerhjem/front2.jpg" alt="hero_image" />
   <Kalender/>
    </div>
  )
}

export default Hero