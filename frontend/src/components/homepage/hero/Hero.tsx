import styles from './Hero.module.css'

const Hero = () => {
  return (
    <div id={styles.heroDiv}>
      <img id={styles.heroImg} src="https://placehold.jp/1200x800.png" alt="hero_image" />
    </div>
  )
}

export default Hero