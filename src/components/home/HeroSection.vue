<script setup>
/* ===== DESKTOP assets (как было) ===== */
import heroShape from '@/assets/hero/hero-shape.png'
import heroWave from '@/assets/hero/hero-wave.svg'
import subtitleBrush from '@/assets/hero/subtitle-brush.svg'
import boy from '@/assets/hero/boy.png'
import brand from '@/assets/hero/hero-brand.png'

/* ===== MOBILE assets (твои новые) ===== */
import heroShapeMobile from '@/assets/hero/hero-shape-mobile.svg'
import brandMobile from '@/assets/hero/hero-brand-mobile.svg'
import subtitleBrushMobile from '@/assets/hero/subtitle-brush-mobile.svg'



</script>

<template>
  <!-- ================= DESKTOP HERO ================= -->
  <section class="hero hero--desktop">
    <div class="hero__inner">
      <img class="hero__wave" :src="heroWave" alt="" aria-hidden="true" />

      <div class="hero__card" :style="{ '--hero-shape': `url(${heroShape})` }">
        <img class="hero__shape" :src="heroShape" alt="" aria-hidden="true" />
        <img class="hero__brand" :src="brand" alt="" aria-hidden="true" />

        <h1 class="hero__title">
          АКАДЕМИЯ<br />
          БОЛЬШИХ<br />
          ВОЗМОЖНОСТЕЙ
        </h1>

        <img class="hero__subtitle-brush" :src="subtitleBrush" alt="" aria-hidden="true" />
        <div class="hero__subtitle-text">Частная семейная школа в Ханты-Мансийске</div>

        <button class="hero__btn" type="button">
          <span class="hero__btn-text">ЗАПИСАТЬСЯ НА ЛИЧНУЮ ВСТРЕЧУ</span>
        </button>

        <img class="hero__boy" :src="boy" alt="Ребёнок" />
      </div>
    </div>
  </section>

  <!-- ================= MOBILE HERO ================= -->
  <section class="hero-mobile hero--mobile">
    <div class="hero-mob">
      <!-- фон/форма -->
      <img class="hero-mob__shape" :src="heroShapeMobile" alt="" aria-hidden="true" />

      <!-- a->A -->
      <img class="hero-mob__brand" :src="brandMobile" alt="" aria-hidden="true" />

      <!-- заголовок -->
      <h1 class="hero-mob__title">
        АКАДЕМИЯ<br />
        БОЛЬШИХ<br />
        ВОЗМОЖНОСТЕЙ
      </h1>

      <!-- мазок + текст -->
      <div class="hero-mob__subtitle">
        <img class="hero-mob__brush" :src="subtitleBrushMobile" alt="" aria-hidden="true" />
        <div class="hero-mob__subtitle-text">Частная семейная школа<br /> в Ханты-Мансийске</div>
      </div>

      <!-- кнопка -->
      <button class="hero-mob__btn" type="button">
        ЗАПИСАТЬСЯ НА<br />
        ЛИЧНУЮ ВСТРЕЧУ
      </button>
    </div>
  </section>
</template>

<style scoped>
/* ===== переключение версий ===== */
.hero--mobile { display: none; }
@media (max-width: 800px) {
  .hero--desktop { display: none; }
  .hero--mobile { display: block; }
}

/* ===== MOBILE LAYOUT (360x420 base) ===== */
.hero-mobile{
  padding: 16px 0 24px;
}

.hero-mob{
  position: relative;
  width: min(360px, calc(100vw - 32px));
  aspect-ratio: 360 / 420;   /* твоя база */
  margin: 0 auto;
}

/* фон */
.hero-mob__shape{
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: fill;
  pointer-events: none;
}

/* бренд a->A */
.hero-mob__brand{
  position: absolute;
  left: 160px;
  top: 16px;
  width: 56px;
  height: auto;
  pointer-events: none;
}

/* заголовок */
.hero-mob__title{
  position: absolute;
  left: 18px;
  top: 56px;
  margin: 0;

  color: #fff;
  font-family: "Rubik Spray Paint", system-ui, sans-serif;
  text-align: center;
  font-weight: 400;
  font-size: 30px;
  line-height: 1.05;
  letter-spacing: 0.04em;
}

/* подзаголовок */
.hero-mob__subtitle{
  position: absolute;
  left: 90px;
  top: 175px;
  width: 240px;
  height: 22px;
  transform: rotate(-1deg);
  transform-origin: left center;
}

.hero-mob__brush{
  position: absolute;
  left: 0;
  top: -9px;
  width: 100%;
  height: auto;
  transform: rotate(2deg);
  transform-origin: left center;
  pointer-events: none;
}

.hero-mob__subtitle-text{
  position: absolute;
  left: 15px;
  top: 6px;

  color: #fff;
  font-family: "Montserrat", system-ui, sans-serif;
  text-align: center;
  font-weight: 600;
  font-size: 14px;
  line-height: 1;
  white-space: nowrap;
}

/* мальчик */
.hero-mob__boy{
  position: absolute;
  right: 4px;
  bottom: 0;
  width: 250px;
  height: auto;
  pointer-events: none;
}

/* кнопка */
.hero-mob__btn{
  position: absolute;
  left: 18px;
  bottom: 16px;
  top: 300px;

  width: 170px;
  height: 63px;
  border: none;
  border-radius: 10px;
  background: #F3E582;
  cursor: pointer;

  font-family: "Manrope", system-ui, sans-serif;
  text-align: center;
  font-weight: 600;
  font-size: 14px;
  color: #232323;

  padding: 0 10px;
}
</style>
